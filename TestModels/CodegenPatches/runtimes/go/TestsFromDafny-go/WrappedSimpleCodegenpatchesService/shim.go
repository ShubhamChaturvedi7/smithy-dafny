// Code generated by smithy-go-codegen DO NOT EDIT.

package WrappedSimpleCodegenPatchesService

import (
	"context"

	"github.com/dafny-lang/DafnyStandardLibGo/Wrappers"
	"github.com/smithy-lang/smithy-dafny/TestModels/CodegenPatches/SimpleCodegenpatchesTypes"
	"github.com/smithy-lang/smithy-dafny/TestModels/CodegenPatches/simplecodegenpatchessmithygenerated"
)

type Shim struct {
	SimpleCodegenpatchesTypes.ICodegenPatchesClient
	client *simplecodegenpatchessmithygenerated.Client
}

func (_static *CompanionStruct_Default___) WrappedCodegenPatches(inputConfig SimpleCodegenpatchesTypes.CodegenPatchesConfig) Wrappers.Result {
	var nativeConfig = simplecodegenpatchessmithygenerated.CodegenPatchesConfig_FromDafny(inputConfig)
	var nativeClient, nativeError = simplecodegenpatchessmithygenerated.NewClient(nativeConfig)
	if nativeError != nil {
		return Wrappers.Companion_Result_.Create_Failure_(SimpleCodegenpatchesTypes.Companion_Error_.Create_Opaque_(nativeError))
	}
	return Wrappers.Companion_Result_.Create_Success_(&Shim{client: nativeClient})
}

func (shim *Shim) GetString(input SimpleCodegenpatchesTypes.GetStringInput) Wrappers.Result {
	var native_request = simplecodegenpatchessmithygenerated.GetStringInput_FromDafny(input)
	var native_response, native_error = shim.client.GetString(context.Background(), native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(simplecodegenpatchessmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(simplecodegenpatchessmithygenerated.GetStringOutput_ToDafny(*native_response))
}
