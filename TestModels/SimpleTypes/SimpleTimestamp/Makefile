# Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
# SPDX-License-Identifier: Apache-2.0

CORES=2

ENABLE_EXTERN_PROCESSING=1
RUST_BENERATED=1

include ../../SharedMakefile.mk

NAMESPACE=simple.types.timestamp

PROJECT_SERVICES := \
	SimpleTypesTimestamp

SERVICE_NAMESPACE_SimpleTypesTimestamp=simple.types.timestamp

SERVICE_DEPS_SimpleTypesTimestamp :=

SMITHY_DEPS=dafny-dependencies/Model/traits.smithy

# This project has no dependencies 
# DEPENDENT-MODELS:= 


POLYMORPH_OPTIONS=--generate client-constructors,project-files


# Constants for languages that drop extern names (Python, Go)

TYPES_FILE_PATH=Model/SimpleTypesTimestampTypes.dfy
TYPES_FILE_WITH_EXTERN_STRING="module {:extern \"simple.types.timestamp.internaldafny.types\" } SimpleTypesTimestampTypes"
TYPES_FILE_WITHOUT_EXTERN_STRING="module SimpleTypesTimestampTypes"

INDEX_FILE_PATH=src/Index.dfy
INDEX_FILE_WITH_EXTERN_STRING="module {:extern \"simple.types.timestamp.internaldafny\" } SimpleTimestamp refines AbstractSimpleTypesTimestampService {"
INDEX_FILE_WITHOUT_EXTERN_STRING="module SimpleTimestamp refines AbstractSimpleTypesTimestampService {"

WRAPPED_INDEX_FILE_PATH=src/WrappedSimpleTypesTimestampImpl.dfy
WRAPPED_INDEX_FILE_WITH_EXTERN_STRING="module {:options \"--function-syntax:4\"} {:extern \"simple.types.timestamp.internaldafny.wrapped\"} WrappedSimpleTypesTimestampService refines WrappedAbstractSimpleTypesTimestampService {"
WRAPPED_INDEX_FILE_WITHOUT_EXTERN_STRING="module {:options \"--function-syntax:4\"} WrappedSimpleTypesTimestampService refines WrappedAbstractSimpleTypesTimestampService {"
