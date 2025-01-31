# Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
# SPDX-License-Identifier: Apache-2.0

CORES=2

ENABLE_EXTERN_PROCESSING=1

include ../../SharedMakefile.mk

PROJECT_SERVICES := \
	ComAmazonawsDynamodb \

SERVICE_NAMESPACE_ComAmazonawsDynamodb=com.amazonaws.dynamodb

SERVICE_DEPS_ComAmazonawsDynamodb :=

SMITHY_DEPS=dafny-dependencies/Model/traits.smithy

AWS_SDK_CMD=--aws-sdk

# This project has no dependencies 
# DEPENDENT-MODELS:= 


# There is no wrapped target for aws-sdk types
_polymorph_wrapped: ;
_polymorph_wrapped_dafny: ;
_polymorph_wrapped_net: ;
_polymorph_wrapped_java: ;
_polymorph_wrapped_python: ;

format_net:
	pushd runtimes/net && dotnet format && popd

clean: _clean
	rm -rf $(LIBRARY_ROOT)/runtimes/java/src/main/dafny-generated
	rm -rf $(LIBRARY_ROOT)/runtimes/java/src/main/smithy-generated
	rm -rf $(LIBRARY_ROOT)/runtimes/java/src/test/dafny-generated

# Python

PYTHON_MODULE_NAME=com_amazonaws_dynamodb

TRANSLATION_RECORD_PYTHON := \
    --translation-record ../../dafny-dependencies/StandardLibrary/runtimes/python/src/standard_library/internaldafny/generated/dafny_src-py.dtr

# Constants for languages that drop extern names (Python, Go)

TYPES_FILE_PATH=Model/ComAmazonawsDynamodbTypes.dfy
TYPES_FILE_WITH_EXTERN_STRING="module {:extern \"software.amazon.cryptography.services.dynamodb.internaldafny.types\" } ComAmazonawsDynamodbTypes"
TYPES_FILE_WITHOUT_EXTERN_STRING="module ComAmazonawsDynamodbTypes"

INDEX_FILE_PATH=src/Index.dfy
INDEX_FILE_WITH_EXTERN_STRING="module {:extern \"software.amazon.cryptography.services.dynamodb.internaldafny\"} Com.Amazonaws.Dynamodb refines AbstractComAmazonawsDynamodbService {"
INDEX_FILE_WITHOUT_EXTERN_STRING="module Com.Amazonaws.Dynamodb refines AbstractComAmazonawsDynamodbService {"
