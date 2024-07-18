// Code generated by smithy-go-codegen DO NOT EDIT.

package simpleconstraintsinternaldafnywrapped

import (
	"github.com/dafny-lang/DafnyStandardLibGo/Wrappers"
	"context"
	"github.com/Smithy-dafny/TestModels/Constraints/simpleconstraints"
	"github.com/Smithy-dafny/TestModels/Constraints/simpleconstraintsinternaldafnytypes"
	"github.com/Smithy-dafny/TestModels/Constraints/simpleconstraintstypes"
)

type Shim struct {
    simpleconstraintsinternaldafnytypes.ISimpleConstraintsClient
    client *simpleconstraints.Client
}

func WrappedSimpleConstraints(inputConfig simpleconstraintsinternaldafnytypes.SimpleConstraintsConfig) Wrappers.Result {
    var nativeConfig = simpleconstraints.SimpleConstraintsConfig_FromDafny(inputConfig)
    var nativeClient, nativeError = simpleconstraints.NewClient(nativeConfig)
    if nativeError != nil {
       return Wrappers.Companion_Result_.Create_Failure_(simpleconstraintsinternaldafnytypes.Companion_Error_.Create_Opaque_(nativeError))
    }
    return Wrappers.Companion_Result_.Create_Success_(&Shim{client: nativeClient})
}

  func (shim *Shim) GetConstraints(input simpleconstraintsinternaldafnytypes.GetConstraintsInput) Wrappers.Result {
      var native_request = simpleconstraints.GetConstraintsInput_FromDafny(input)
      var native_response, native_error = shim.client.GetConstraints(context.Background() , native_request)
      if native_error != nil {
          return Wrappers.Companion_Result_.Create_Failure_(simpleconstraints.Error_ToDafny(native_error))
      }
      return Wrappers.Companion_Result_.Create_Success_(simpleconstraints.GetConstraintsOutput_ToDafny(*native_response))
  }
