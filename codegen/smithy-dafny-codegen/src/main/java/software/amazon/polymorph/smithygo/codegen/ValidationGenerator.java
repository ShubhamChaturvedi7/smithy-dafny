package software.amazon.polymorph.smithygo.codegen;

import java.math.BigDecimal;
import java.util.Optional;

import software.amazon.polymorph.traits.DafnyUtf8BytesTrait;
import software.amazon.polymorph.traits.ReferenceTrait;
import software.amazon.smithy.codegen.core.Symbol;
import software.amazon.smithy.codegen.core.SymbolProvider;
import software.amazon.smithy.model.Model;
import software.amazon.smithy.model.shapes.MemberShape;
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

    // This strings are used for list and map
    private static final String LIST_ITEM = "item";
    private static final String MAP_KEY = "key";
    private static final String MAP_VALUE = "value";

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
                    String memberName;
                    if (containerShape.isListShape() || containerShape.isMapShape())
                        memberName = dataSource;
                    else 
                        memberName = dataSource + "." + symbolProvider.toMemberName(member);
                    
                    renderValidatorForEachShape(model.expectShape(member.getTarget()), sortedMembers, isInputStructure, memberName);
                });
    }

    private void renderValidatorForEachShape (Shape currentShape, CodegenUtils.SortedMembers sortedMembers, boolean isInputStructure, String dataSource) {
                    
                    Symbol memberSymbol = symbolProvider.toSymbol(currentShape);
                    if (isInputStructure) {
                        memberSymbol = memberSymbol.getProperty(SymbolUtils.INPUT_VARIANT, Symbol.class)
                                .orElse(memberSymbol);
                    }
                    if (currentShape.hasTrait(ReferenceTrait.class)) {
                        memberSymbol = memberSymbol.getProperty("Referred", Symbol.class).get();
                    } 
                    if (currentShape.hasTrait(RangeTrait.class)) {
                        addRangeCheck(memberSymbol, currentShape, dataSource);
                    }
                    if (currentShape.hasTrait(LengthTrait.class)) {
                        addLengthCheck(memberSymbol, currentShape, dataSource);
                    }
                    if (currentShape.hasTrait(RequiredTrait.class)) {
                        addRequiredCheck(memberSymbol, currentShape, dataSource);
                    }
                    if (currentShape.hasTrait(DafnyUtf8BytesTrait.class)) {
                        addUTFCheck(memberSymbol, currentShape, dataSource);
                    }
                    // Broke list and map into two different if else because for _, item := range %s looked good for list
                    // And for key, value := range %s looked good for map
                    if (currentShape.isListShape()) {
                        writer.write("""
                                for _, %s := range %s {
                                    // To avoid declared and not used error for shapes which does not need validation check
                                    _ = item
                                """.formatted(LIST_ITEM, dataSource));
                        renderValidatorHelper(currentShape,sortedMembers,isInputStructure, LIST_ITEM);
                        writer.write("""
                                }
                                """);
                    }
                    else if (currentShape.isMapShape()) {
                        writer.write("""
                                for %s, %s := range %s {
                                    // To avoid declared and not used error for shapes which does not need validation check
                                    _ = key
                                    _ = value
                                """.formatted(MAP_KEY, MAP_VALUE, dataSource));
                        renderValidatorHelper(currentShape,sortedMembers,isInputStructure, MAP_KEY);
                        renderValidatorHelper(currentShape,sortedMembers,isInputStructure, MAP_VALUE);
                        writer.write("""
                            }
                        """);
                    }
                    else if (currentShape.isUnionShape()) {
                        writer.write("""
                            switch unionType := %s.(type) {
                                """.formatted(dataSource));
                        for (var memberInUnion : currentShape.getAllMembers().values()) {
                            writer.write("""
                                    case *%s:
                                    """.formatted(
                                        symbolProvider.toMemberName(memberInUnion)
                                        ));
                                    
                            renderValidatorForEachShape(model.expectShape(memberInUnion.getTarget()), sortedMembers, isInputStructure, "unionType.Value");
                        }
                        writer.write("""
                                // Default case should not be reached. 
                                default:
                                    // To avoid used and not used error when nothing to validate
                                    _ = unionType
				                    panic("Unhandled union type")
                                }
                                    """);
                    }
                    else {
                        // This call will help when structure is inside structure. 
                        // But what about unions?
                        renderValidatorHelper(currentShape,sortedMembers,isInputStructure, dataSource);
                    }
    }

    private void addRangeCheck(Symbol memberSymbol, Shape currentShape, String dataSource) {
        String pointableString = "";
        String rangeCheck = "";
        RangeTrait rangeTraitShape = currentShape.expectTrait(RangeTrait.class);
        Optional<BigDecimal> min = rangeTraitShape.getMin();
        Optional<BigDecimal> max = rangeTraitShape.getMax();
        if (!(dataSource.equals(LIST_ITEM) || dataSource.equals(MAP_KEY) || dataSource.equals(MAP_VALUE))) {
            if ((boolean) memberSymbol.getProperty(POINTABLE, Boolean.class).orElse(false)){
                pointableString = "*";
            }
        }
        if (pointableString.equals("*")){
            rangeCheck += """
                    if (%s != nil) {
                """.formatted(dataSource);
        }

        if (min.isPresent()) {
            rangeCheck += """
                    if (%s%s < %s) {
                        return fmt.Errorf(\"%s has a minimum of %s but has the value of %%d.\", %s%s)
                    }
                    """.formatted(
                        pointableString,
                        dataSource,
                        min.get().toString(),
                        currentShape.getId().getName(),
                        min.get().toString(),
                        pointableString,
                        dataSource);
        }
        if (max.isPresent()) {
            rangeCheck += """
                    if (%s%s > %s) {
                        return fmt.Errorf(\"%s has a maximum of %s but has the value of %%d.\", %s%s)
                    }
                    """.formatted(
                        pointableString,
                        dataSource,
                        max.get().toString(),
                        currentShape.getId().getName(),
                        max.get().toString(),
                        pointableString,
                        dataSource);
        }
        if (pointableString.equals("*")){
            rangeCheck += """
                }
                """;
        }
        writer.write(rangeCheck);
        
    }

    private void addLengthCheck(Symbol memberSymbol, Shape currentShape, String dataSource) {
        String pointableString = "";
        String lengthCheck = "";
        LengthTrait lengthTraitShape = currentShape.expectTrait(LengthTrait.class);
        Optional<Long> min = lengthTraitShape.getMin();
        Optional<Long> max = lengthTraitShape.getMax();
        System.out.println(memberSymbol.getProperties());
        if (!(dataSource.equals(LIST_ITEM) || dataSource.equals(MAP_KEY) || dataSource.equals(MAP_VALUE))) {
            if ((boolean) memberSymbol.getProperty(POINTABLE, Boolean.class).orElse(false)){
                pointableString = "*";
            }
        }
        if (pointableString.equals("*")){
            lengthCheck += """
                    if (%s != nil) {
                """.formatted(dataSource);
        }
        if (min.isPresent()) {
            if (currentShape.hasTrait(DafnyUtf8BytesTrait.class)) {
                lengthCheck += """
                    if (utf8.RuneCountInString(%s%s) < %s) {
                        return fmt.Errorf(\"%s has a minimum length of %s but has the length of %%d.\", utf8.RuneCountInString(%s%s))
                    }
                    """.formatted(
                        pointableString,
                        dataSource,
                        min.get().toString(),
                        currentShape.getId().getName(),
                        min.get().toString(),
                        pointableString,
                        dataSource);           
            }
            else {
                lengthCheck += """
                        if (len(%s%s) < %s) {
                            return fmt.Errorf(\"%s has a minimum length of %s but has the length of %%d.\", len(%s%s))
                        }
                        """.formatted(
                            pointableString,
                            dataSource,
                            min.get().toString(),
                            currentShape.getId().getName(),
                            min.get().toString(),
                            pointableString,
                            dataSource);
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
                        dataSource,
                        max.get().toString(),
                        currentShape.getId().getName(),
                        max.get().toString(),
                        pointableString,
                        dataSource);
            }
            else {
                lengthCheck += """
                        if (len(%s%s) > %s) {
                            return fmt.Errorf(\"%s has a maximum length of %s but has the length of %%d.\", len(%s%s))
                        }
                        """.formatted(
                            pointableString,
                            dataSource,
                            max.get().toString(),
                            currentShape.getId().getName(),
                            max.get().toString(),
                            pointableString,
                            dataSource);
            }
        }
        if (pointableString.equals("*")){
            lengthCheck += """
                }
                """;
        }
        writer.write(lengthCheck);
    }

    private void addRequiredCheck(Symbol memberSymbol, Shape currentShape, String dataSource) {
        String RequiredCheck = "";
        if( memberSymbol.getProperty(POINTABLE).isPresent() && (boolean) memberSymbol.getProperty(POINTABLE).get()) 
            RequiredCheck += """
                    if ( %s == nil ) {
                        return fmt.Errorf(\"%s is required but has a nil value.\")
                    }
                    """.formatted(
                    dataSource,
                    dataSource);
        writer.write(RequiredCheck);
    }

    private void addUTFCheck(Symbol memberSymbol, Shape currentShape, String dataSource) {
        String pointableString = "";
        String UTFCheck = "";
        if (!(dataSource.equals(LIST_ITEM) || dataSource.equals(MAP_KEY) || dataSource.equals(MAP_VALUE))) {
            if ((boolean) memberSymbol.getProperty(POINTABLE, Boolean.class).orElse(false)){
                pointableString = "*";
            }
        }
        if (pointableString.equals("*")){
            UTFCheck += """
                    if ( %s != nil ) {
                """.formatted(dataSource);
        }
        UTFCheck += """
                    if (!utf8.ValidString(%s%s)) {
                        return fmt.Errorf(\"Invalid UTF bytes %%s \", %s%s)
                    }
                    """.formatted(
                        pointableString,
                        dataSource,
                        pointableString,
                        dataSource);
        if (pointableString.equals("*")){
            UTFCheck += """
                }
                """;
        }
        writer.write(UTFCheck);
    }
}
