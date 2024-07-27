package software.amazon.polymorph.smithygo.codegen;

import java.math.BigDecimal;
import java.util.Optional;

import software.amazon.polymorph.traits.DafnyUtf8BytesTrait;
import software.amazon.polymorph.traits.ReferenceTrait;
import software.amazon.smithy.codegen.core.Symbol;
import software.amazon.smithy.codegen.core.SymbolProvider;
import software.amazon.smithy.model.Model;
import software.amazon.smithy.model.shapes.Shape;
import software.amazon.smithy.model.traits.LengthTrait;
import software.amazon.smithy.model.traits.RangeTrait;
import software.amazon.smithy.model.traits.RequiredTrait;
import software.amazon.smithy.model.traits.StreamingTrait;

import static software.amazon.polymorph.smithygo.codegen.SymbolUtils.POINTABLE;

// Renders constraint validation
public class ValidationGenerator {
    private final Model model;
    private final Shape shape;
    private final Symbol symbol;
    private final SymbolProvider symbolProvider;
    private final GoWriter writer;
    
    public ValidationGenerator(
        Model model,
        Shape shape,
        SymbolProvider symbolProvider,
        GoWriter writer
    ) {
        this.model = model;
        this.shape = shape;
        this.symbol = symbolProvider.toSymbol(shape);
        this.symbolProvider = symbolProvider;
        this.writer = writer;
    }

    public void renderValidator (boolean isInputStructure) {
        CodegenUtils.SortedMembers sortedMembers = new CodegenUtils.SortedMembers(symbolProvider);
        writer.openBlock("func (input $L) Validate() (error) {", symbol.getName());
        renderValidatorHelper( shape, sortedMembers, isInputStructure, "input");
        writer.write("return nil");
        writer.closeBlock("}").write("");
    }

    private void renderValidatorHelper (Shape containerShape, CodegenUtils.SortedMembers sortedMembers, boolean isInputStructure, String dataSource) {
        containerShape.getAllMembers().values().stream()
                .filter(memberShape -> !StreamingTrait.isEventStream(model, memberShape))
                .sorted(sortedMembers)
                .forEach((member) -> {
                    String memberName = dataSource + "." + symbolProvider.toMemberName(member);
                    if (containerShape.isListShape() || containerShape.isMapShape())
                        memberName = dataSource;
                    Symbol memberSymbol = symbolProvider.toSymbol(member);
                    if (isInputStructure) {
                        memberSymbol = memberSymbol.getProperty(SymbolUtils.INPUT_VARIANT, Symbol.class)
                                .orElse(memberSymbol);
                    }
                    if (model.expectShape(member.getTarget()).hasTrait(ReferenceTrait.class)) {
                        memberSymbol = memberSymbol.getProperty("Referred", Symbol.class).get();
                    }
                    Shape currentShape = model.expectShape(member.getTarget());   
                    if (currentShape.hasTrait(RangeTrait.class)) {
                        addRangeCheck(memberSymbol, currentShape, memberName);
                    }
                    if (currentShape.hasTrait(LengthTrait.class)) {
                        addLengthCheck(memberSymbol, currentShape, memberName);
                    }
                    if (member.hasTrait(RequiredTrait.class)) {
                        addRequiredCheck(memberSymbol, currentShape, memberName);
                    }
                    if (currentShape.hasTrait(DafnyUtf8BytesTrait.class)) {
                        addUTFCheck(memberSymbol, currentShape, memberName);
                    }
                    // Broke list and map into two different if else because for _, item := range %s looked good for list
                    // And for key, value := range %s looked good for map
                    if (currentShape.isListShape()) {
                        writer.write("""
                                for _, item := range %s {
                                    // To avoid declared and not used error for shapes which does not need validation check
                                    _ = item
                                """.formatted(memberName));
                        renderValidatorHelper(currentShape,sortedMembers,isInputStructure, "item");
                        writer.write("""
                                }
                                """);
                    }
                    else if (currentShape.isMapShape()) {
                        writer.write("""
                                for key, value := range %s {
                                    // To avoid declared and not used error for shapes which does not need validation check
                                    _ = key
                                    _ = value
                                """.formatted(memberName));
                        renderValidatorHelper(currentShape,sortedMembers,isInputStructure, "key");
                        renderValidatorHelper(currentShape,sortedMembers,isInputStructure, "value");
                        writer.write("""
                            }
                        """);
                    }
                    // Add else to avoid calling renderValidatorHelper twice when currentShape is ListShape and MapShape
                    else {
                        // This call will help when structure is inside structure. 
                        // But what about unions?
                        renderValidatorHelper(currentShape,sortedMembers,isInputStructure, dataSource);
                    }
                });
    }

    private void addRangeCheck(Symbol memberSymbol, Shape currentShape, String memberName) {
        String pointableString = "";
        String rangeCheck = "";
        RangeTrait rangeTraitShape = currentShape.expectTrait(RangeTrait.class);
        Optional<BigDecimal> min = rangeTraitShape.getMin();
        Optional<BigDecimal> max = rangeTraitShape.getMax();

        if ((boolean) memberSymbol.getProperty(POINTABLE).orElse(false) == true){
            pointableString = "*";
        }
        
        if (pointableString.equals("*")){
            rangeCheck += """
                    if (%s != nil) {
                """.formatted(memberName);
        }

        if (min.isPresent()) {
            rangeCheck += """
                    if (%s%s < %s) {
                        return fmt.Errorf(\"%s has a minimum of %s but has the value of %%d.\", %s%s)
                    }
                    """.formatted(
                        pointableString,
                        memberName,
                        min.get().toString(),
                        currentShape.getId().getName(),
                        min.get().toString(),
                        pointableString,
                        memberName);
        }
        if (max.isPresent()) {
            rangeCheck += """
                    if (%s%s > %s) {
                        return fmt.Errorf(\"%s has a maximum of %s but has the value of %%d.\", %s%s)
                    }
                    """.formatted(
                        pointableString,
                        memberName,
                        max.get().toString(),
                        currentShape.getId().getName(),
                        max.get().toString(),
                        pointableString,
                        memberName);
        }
        if (pointableString.equals("*")){
            rangeCheck += """
                }
                """;
        }
        writer.write(rangeCheck);
        
    }

    private void addLengthCheck(Symbol memberSymbol, Shape currentShape, String memberName) {
        String pointableString = "";
        String lengthCheck = "";
        LengthTrait lengthTraitShape = currentShape.expectTrait(LengthTrait.class);
        Optional<Long> min = lengthTraitShape.getMin();
        Optional<Long> max = lengthTraitShape.getMax();
        if ((boolean) memberSymbol.getProperty(POINTABLE).orElse(false) == true){
            pointableString = "*";
        }
        if (pointableString.equals("*")){
            lengthCheck += """
                    if (%s != nil) {
                """.formatted(memberName);
        }
        if (min.isPresent()) {
            if (currentShape.hasTrait(DafnyUtf8BytesTrait.class)) {
                lengthCheck += """
                    if (utf8.RuneCountInString(%s%s) < %s) {
                        return fmt.Errorf(\"%s has a minimum length of %s but has the length of %%d.\", utf8.RuneCountInString(%s%s))
                    }
                    """.formatted(
                        pointableString,
                        memberName,
                        min.get().toString(),
                        currentShape.getId().getName(),
                        min.get().toString(),
                        pointableString,
                        memberName);           
            }
            else {
                lengthCheck += """
                        if (len(%s%s) < %s) {
                            return fmt.Errorf(\"%s has a minimum length of %s but has the length of %%d.\", len(%s%s))
                        }
                        """.formatted(
                            pointableString,
                            memberName,
                            min.get().toString(),
                            currentShape.getId().getName(),
                            min.get().toString(),
                            pointableString,
                            memberName);
            }
        }
        if (max.isPresent()) {
            if (currentShape.hasTrait(DafnyUtf8BytesTrait.class)) {
                lengthCheck += """
                    if (utf8.RuneCountInString(%s%s) > %s) {
                        return fmt.Errorf(\"%s has a maximum length of %s but has the length of %%d.\", utf8.RuneCountInString(%s%s))
                    }
                    """.formatted(
                        pointableString,
                        memberName,
                        max.get().toString(),
                        currentShape.getId().getName(),
                        max.get().toString(),
                        pointableString,
                        memberName);
            }
            else {
                lengthCheck += """
                        if (len(%s%s) > %s) {
                            return fmt.Errorf(\"%s has a maximum length of %s but has the length of %%d.\", len(%s%s))
                        }
                        """.formatted(
                            pointableString,
                            memberName,
                            max.get().toString(),
                            currentShape.getId().getName(),
                            max.get().toString(),
                            pointableString,
                            memberName);
            }
        }
        if (pointableString.equals("*")){
            lengthCheck += """
                }
                """;
        }
        writer.write(lengthCheck);
    }

    private void addRequiredCheck(Symbol memberSymbol, Shape currentShape, String memberName) {
        String RequiredCheck = "";
        if( memberSymbol.getProperty(POINTABLE).isPresent() && (boolean) memberSymbol.getProperty(POINTABLE).get()) 
            RequiredCheck += """
                    if ( %s == nil ) {
                        return fmt.Errorf(\"%s is required but has a nil value.\")
                    }
                    """.formatted(
                    memberName,
                    memberName);
        writer.write(RequiredCheck);
    }

    private void addUTFCheck(Symbol memberSymbol, Shape currentShape, String memberName) {
        String pointableString = "";
        String UTFCheck = "";
        if ((boolean) memberSymbol.getProperty(POINTABLE).orElse(false) == true){
            pointableString = "*";
        }
        if (pointableString.equals("*")){
            UTFCheck += """
                    if ( %s != nil ) {
                """.formatted(memberName);
        }
        UTFCheck += """
                    if (!utf8.ValidString(%s%s)) {
                        return fmt.Errorf(\"Invalid UTF bytes %%s \", %s%s)
                    }
                    """.formatted(
                        pointableString,
                        memberName,
                        pointableString,
                        memberName);
        if (pointableString.equals("*")){
            UTFCheck += """
                }
                """;
        }
        writer.write(UTFCheck);
    }
}
