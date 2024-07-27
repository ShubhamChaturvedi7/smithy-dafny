package software.amazon.polymorph.smithygo.localservice.nameresolver;

import software.amazon.smithy.codegen.core.Symbol;
import software.amazon.smithy.model.shapes.ServiceShape;
import software.amazon.smithy.model.shapes.Shape;

import java.util.Map;

import static software.amazon.polymorph.smithygo.localservice.nameresolver.Constants.BLANK;
import static software.amazon.polymorph.smithygo.localservice.nameresolver.Constants.DOT;

public class SmithyNameResolver {

    private static Map<String, String> smithyNamespaceToGoModuleNameMap;

    public static void setSmithyNamespaceToGoModuleNameMap(
            Map<String, String> smithyNamespaceToGoModuleNameMap) {
        SmithyNameResolver.smithyNamespaceToGoModuleNameMap = smithyNamespaceToGoModuleNameMap;
    }

    public static String getGoModuleNameForSmithyNamespace(final String smithyNamespace) {
        if (smithyNamespace.contains("smithy.")) return "";
        if (!smithyNamespaceToGoModuleNameMap.containsKey(smithyNamespace)) {
            throw new IllegalArgumentException("Go module name not found for Smithy namespace: " + smithyNamespace);
        }
        return smithyNamespaceToGoModuleNameMap.get(smithyNamespace);
    }

    public static String shapeNamespace(final Shape shape) {
        return shape.toShapeId().getNamespace().replace(DOT, BLANK).toLowerCase();
    }

    public static String smithyTypesNamespace(final Shape shape) {
        return shape.toShapeId().getNamespace().replace(DOT, BLANK).toLowerCase().concat("types");
    }

    public static String getGoModuleNameForSdkNamespace(final String smithyNamespace) {
        return getGoModuleNameForSmithyNamespace("sdk.".concat(smithyNamespace));
    }

    public static String smithyTypesNamespaceAws(final Shape shape, boolean isAwsSubType) {
        if (isAwsSubType) {
            return "types";
        }
        return "kms";
    }

    public static String getSmithyType(final Shape shape, final Symbol symbol) {
        if(symbol.getNamespace().contains("smithy")) {
            return symbol.getName();
        }
        return SmithyNameResolver.smithyTypesNamespace(shape).concat(DOT).concat(symbol.getName());
    }

    public static String getSmithyTypeAws(final Shape shape, final Symbol symbol, boolean subtype) {
        if(symbol.getNamespace().contains("smithy.")) {
            return symbol.getName();
        }
        return SmithyNameResolver.smithyTypesNamespaceAws(shape, subtype).concat(DOT).concat(symbol.getName());
    }

    public static String getSmithyType(final Shape shape) {
        return SmithyNameResolver.smithyTypesNamespace(shape).concat(DOT).concat(shape.toShapeId().getName());
    }

    public static String getToDafnyMethodName(final ServiceShape serviceShape, final Shape shape, final String disambiguator) {
        final var methodName = serviceShape.getContextualName(shape).concat("_ToDafny");
        if (serviceShape.toShapeId().getNamespace().equals(shape.toShapeId().getNamespace())) {
            return methodName;
        } else {
            return SmithyNameResolver.shapeNamespace(shape).concat(DOT).concat(methodName);
        }
    }

    public static String getToDafnyMethodName(final Shape shape, final String disambiguator) {
        final var methodName = shape.getId().getName().concat("_ToDafny");
            return SmithyNameResolver.shapeNamespace(shape).concat(DOT).concat(methodName);
    }

    public static String getFromDafnyMethodName(final ServiceShape serviceShape, final Shape shape, final String disambiguator) {
        final var methodName = serviceShape.getContextualName(shape).concat("_FromDafny");
        if (serviceShape.toShapeId().getNamespace().equals(shape.toShapeId().getNamespace())) {
            return methodName;
        } else {
            return SmithyNameResolver.shapeNamespace(shape).concat(DOT).concat(methodName);
        }
    }

    public static String getFromDafnyMethodName(final Shape shape, final String disambiguator) {
        final var methodName = shape.getId().getName().concat("_FromDafny");
        return SmithyNameResolver.shapeNamespace(shape).concat(DOT).concat(methodName);
    }
}
