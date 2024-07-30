package software.amazon.polymorph.smithygo.codegen;

import java.util.Collection;
import java.util.HashSet;
import java.util.Set;
import java.util.TreeSet;
import java.util.stream.Collectors;

import software.amazon.smithy.codegen.core.Symbol;
import software.amazon.smithy.codegen.core.SymbolProvider;
import software.amazon.smithy.model.Model;
import software.amazon.smithy.model.shapes.MemberShape;
import software.amazon.smithy.model.shapes.SimpleShape;
import software.amazon.smithy.model.shapes.UnionShape;
import software.amazon.smithy.model.traits.StreamingTrait;
import software.amazon.smithy.model.shapes.Shape;
import software.amazon.smithy.model.traits.ErrorTrait;

public class UnionGenerator {
    public static final String UNKNOWN_MEMBER_NAME = "UnknownUnionMember";

    private final Model model;
    private final SymbolProvider symbolProvider;
    private final UnionShape shape;
    private final boolean isEventStream;

    public UnionGenerator(Model model, SymbolProvider symbolProvider, UnionShape shape) {
        this.model = model;
        this.symbolProvider = symbolProvider;
        this.shape = shape;
        this.isEventStream = StreamingTrait.isEventStream(shape);
    }

    /**
     * Generates the Go type definitions for the UnionShape.
     *
     * @param writer the writer
     */
    public void generateUnion(GoWriter writer) {
        Symbol symbol = symbolProvider.toSymbol(shape);
        Collection<MemberShape> memberShapes = shape.getAllMembers().values()
                .stream()
                .filter(memberShape -> !isEventStreamErrorMember(memberShape))
                .collect(Collectors.toCollection(TreeSet::new));

        memberShapes.stream().map(symbolProvider::toMemberName).forEach(name -> {
            writer.write("//  " + name);
        });
        writer.openBlock("type $L interface {", "}", symbol.getName(), () -> {
            writer.write("is$L()", symbol.getName());
        }).write("");

        // Create structs for each member that satisfy the interface.
        for (MemberShape member : memberShapes) {
            Symbol memberSymbol = symbolProvider.toSymbol(member);
            String exportedMemberName = symbolProvider.toMemberName(member);
            Shape target = model.expectShape(member.getTarget());

            writer.openBlock("type $L struct {", "}", exportedMemberName, () -> {
                // Union members can't have null values, so for simple shapes we don't
                // use pointers. We have to use pointers for complex shapes since,
                // for example, we could still have a map that's empty or which has
                // null values.
                if (target instanceof SimpleShape) {
                    writer.write("Value $T", memberSymbol);
                } else {
                    writer.write("Value $P", memberSymbol);
                }
                writer.write("");
            });

            writer.write("func (*$L) is$L() {}", exportedMemberName, symbol.getName());
        }
    }

    private boolean isEventStreamErrorMember(MemberShape memberShape) {
        return isEventStream && memberShape.getMemberTrait(model, ErrorTrait.class).isPresent();
    }
}
