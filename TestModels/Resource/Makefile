# Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
# SPDX-License-Identifier: Apache-2.0

CORES=2

ENABLE_EXTERN_PROCESSING=1
RUST_BENERATED=1

include ../SharedMakefile.mk

NAMESPACE=simple.resources

PROJECT_SERVICES := \
	SimpleResources

SERVICE_NAMESPACE_SimpleResources=simple.resources

SERVICE_DEPS_SimpleResources :=

SMITHY_DEPS=dafny-dependencies/Model/traits.smithy

# This project has no dependencies 
# DEPENDENT-MODELS:= 

# Python

PYTHON_MODULE_NAME=simple_resources

TRANSLATION_RECORD_PYTHON := \
    --translation-record ../dafny-dependencies/StandardLibrary/runtimes/python/src/standard_library/internaldafny/generated/dafny_src-py.dtr

format_net:
	pushd runtimes/net && dotnet format && popd

# Constants for languages that drop extern names (Python, Go)

TYPES_FILE_PATH=Model/SimpleResourcesTypes.dfy
TYPES_FILE_WITH_EXTERN_STRING="module {:extern \"simple.resources.internaldafny.types\" } SimpleResourcesTypes"
TYPES_FILE_WITHOUT_EXTERN_STRING="module SimpleResourcesTypes"

INDEX_FILE_PATH=src/Index.dfy
INDEX_FILE_WITH_EXTERN_STRING="module {:extern \"simple.resources.internaldafny\"} SimpleResources refines AbstractSimpleResourcesService"
INDEX_FILE_WITHOUT_EXTERN_STRING="module SimpleResources refines AbstractSimpleResourcesService"

WRAPPED_INDEX_FILE_PATH=src/WrappedIndex.dfy
WRAPPED_INDEX_FILE_WITH_EXTERN_STRING="module {:extern \"simple.resources.internaldafny.wrapped\"} WrappedSimpleResources refines WrappedAbstractSimpleResourcesService"
WRAPPED_INDEX_FILE_WITHOUT_EXTERN_STRING="module WrappedSimpleResources refines WrappedAbstractSimpleResourcesService"
