# Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
# SPDX-License-Identifier: Apache-2.0

CORES=2

ENABLE_EXTERN_PROCESSING=1
RUST_BENERATED=1

include ../../SharedMakefile.mk

NAMESPACE=simple.types.smithyDouble

PROJECT_SERVICES := \
	SimpleDouble

SERVICE_NAMESPACE_SimpleDouble=simple.types.smithyDouble

SERVICE_DEPS_SimpleDouble :=

SMITHY_DEPS=dafny-dependencies/Model/traits.smithy

# This project has no dependencies 
# DEPENDENT-MODELS:= 

transpile_net_dependencies:

format_net:
	pushd runtimes/net && dotnet format && popd

# Python

PYTHON_MODULE_NAME=simple_types_smithydouble

TRANSLATION_RECORD_PYTHON := \
    --translation-record ../../dafny-dependencies/StandardLibrary/runtimes/python/src/standard_library/internaldafny/generated/dafny_src-py.dtr

# Constants for languages that drop extern names (Python, Go)

TYPES_FILE_PATH=Model/SimpleTypesSmithyDoubleTypes.dfy
TYPES_FILE_WITH_EXTERN_STRING="module {:extern \"simple.types.smithydouble.internaldafny.types\" } SimpleTypesSmithyDoubleTypes"
TYPES_FILE_WITHOUT_EXTERN_STRING="module SimpleTypesSmithyDoubleTypes"

INDEX_FILE_PATH=src/Index.dfy
INDEX_FILE_WITH_EXTERN_STRING="module {:extern \"simple.types.smithydouble.internaldafny\" } SimpleDouble refines AbstractSimpleTypesSmithyDoubleService {"
INDEX_FILE_WITHOUT_EXTERN_STRING="module SimpleDouble refines AbstractSimpleTypesSmithyDoubleService {"

WRAPPED_INDEX_FILE_PATH=src/WrappedIndex.dfy
WRAPPED_INDEX_FILE_WITH_EXTERN_STRING="module {:extern \"simple.types.smithydouble.internaldafny.wrapped\"} WrappedSimpleTypesDouble refines WrappedAbstractSimpleTypesSmithyDoubleService"
WRAPPED_INDEX_FILE_WITHOUT_EXTERN_STRING="module WrappedSimpleTypesDouble refines WrappedAbstractSimpleTypesSmithyDoubleService"
