# Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
# SPDX-License-Identifier: Apache-2.0

CORES=2

ENABLE_EXTERN_PROCESSING=1
RUST_BENERATED=1

include ../../SharedMakefile.mk

NAMESPACE=simple.types.smithyString

PROJECT_SERVICES := \
	SimpleString

SERVICE_NAMESPACE_SimpleString=simple.types.smithyString

SERVICE_DEPS_SimpleString :=

SMITHY_DEPS=dafny-dependencies/Model/traits.smithy

# This project has no dependencies 
# DEPENDENT-MODELS:=

# Python

PYTHON_MODULE_NAME=simple_types_smithystring

TRANSLATION_RECORD_PYTHON := \
    --translation-record ../../dafny-dependencies/StandardLibrary/runtimes/python/src/standard_library/internaldafny/generated/dafny_src-py.dtr

# Constants for languages that drop extern names (Python, Go)

TYPES_FILE_PATH=Model/SimpleTypesSmithyStringTypes.dfy
TYPES_FILE_WITH_EXTERN_STRING="module {:extern \"simple.types.smithystring.internaldafny.types\" } SimpleTypesSmithyStringTypes"
TYPES_FILE_WITHOUT_EXTERN_STRING="module SimpleTypesSmithyStringTypes"

INDEX_FILE_PATH=src/Index.dfy
INDEX_FILE_WITH_EXTERN_STRING="module {:extern \"simple.types.smithystring.internaldafny\" } SimpleString refines AbstractSimpleTypesSmithyStringService {"
INDEX_FILE_WITHOUT_EXTERN_STRING="module SimpleString refines AbstractSimpleTypesSmithyStringService {"

WRAPPED_INDEX_FILE_PATH=src/WrappedSimpleStringImpl.dfy
WRAPPED_INDEX_FILE_WITH_EXTERN_STRING="module {:extern \"simple.types.smithystring.internaldafny.wrapped\"} WrappedSimpleTypesStringService refines WrappedAbstractSimpleTypesSmithyStringService {"
WRAPPED_INDEX_FILE_WITHOUT_EXTERN_STRING="module WrappedSimpleTypesStringService refines WrappedAbstractSimpleTypesSmithyStringService {"
