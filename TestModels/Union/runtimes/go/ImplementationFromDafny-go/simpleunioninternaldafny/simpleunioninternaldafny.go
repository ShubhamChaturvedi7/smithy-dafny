// Package simpleunioninternaldafny
// Dafny module simpleunioninternaldafny compiled into Go

package simpleunioninternaldafny

import (
	os "os"

	SimpleUnionImpl "github.com/Smithy-dafny/TestModels/Union/SimpleUnionImpl"
	simpleunioninternaldafnytypes "github.com/Smithy-dafny/TestModels/Union/simpleunioninternaldafnytypes"
	_System "github.com/dafny-lang/DafnyRuntimeGo/System_"
	_dafny "github.com/dafny-lang/DafnyRuntimeGo/dafny"
	StandardLibrary "github.com/dafny-lang/DafnyStandardLibGo/StandardLibrary"
	StandardLibraryInterop "github.com/dafny-lang/DafnyStandardLibGo/StandardLibraryInterop"
	StandardLibrary_UInt "github.com/dafny-lang/DafnyStandardLibGo/StandardLibrary_UInt"
	Wrappers "github.com/dafny-lang/DafnyStandardLibGo/Wrappers"
)

var _ = os.Args
var _ _dafny.Dummy__
var _ _System.Dummy__
var _ Wrappers.Dummy__
var _ StandardLibrary_UInt.Dummy__
var _ StandardLibrary.Dummy__
var _ StandardLibraryInterop.Dummy__
var _ SimpleUnionImpl.Dummy__

type Dummy__ struct{}

// Definition of class Default__
type Default__ struct {
	dummy byte
}

func New_Default___() *Default__ {
	_this := Default__{}

	return &_this
}

type CompanionStruct_Default___ struct {
}

var Companion_Default___ = CompanionStruct_Default___{}

func (_this *Default__) Equals(other *Default__) bool {
	return _this == other
}

func (_this *Default__) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*Default__)
	return ok && _this.Equals(other)
}

func (*Default__) String() string {
	return "simpleunioninternaldafny.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) DefaultSimpleUnionConfig() simpleunioninternaldafnytypes.SimpleUnionConfig {
	return simpleunioninternaldafnytypes.Companion_SimpleUnionConfig_.Create_SimpleUnionConfig_()
}
func (_static *CompanionStruct_Default___) SimpleUnion(config simpleunioninternaldafnytypes.SimpleUnionConfig) Wrappers.Result {
	var res Wrappers.Result = Wrappers.Result{}
	_ = res
	var _3_client *SimpleUnionClient
	_ = _3_client
	var _nw0 *SimpleUnionClient = New_SimpleUnionClient_()
	_ = _nw0
	_nw0.Ctor__(SimpleUnionImpl.Companion_Config_.Create_Config_())
	_3_client = _nw0
	res = Wrappers.Companion_Result_.Create_Success_(_3_client)
	return res
	return res
}
func (_static *CompanionStruct_Default___) CreateSuccessOfClient(client simpleunioninternaldafnytypes.ISimpleUnionClient) Wrappers.Result {
	return Wrappers.Companion_Result_.Create_Success_(client)
}
func (_static *CompanionStruct_Default___) CreateFailureOfError(error_ simpleunioninternaldafnytypes.Error) Wrappers.Result {
	return Wrappers.Companion_Result_.Create_Failure_(error_)
}

// End of class Default__

// Definition of class SimpleUnionClient
type SimpleUnionClient struct {
	_config SimpleUnionImpl.Config
}

func New_SimpleUnionClient_() *SimpleUnionClient {
	_this := SimpleUnionClient{}

	_this._config = SimpleUnionImpl.Companion_Config_.Default()
	return &_this
}

type CompanionStruct_SimpleUnionClient_ struct {
}

var Companion_SimpleUnionClient_ = CompanionStruct_SimpleUnionClient_{}

func (_this *SimpleUnionClient) Equals(other *SimpleUnionClient) bool {
	return _this == other
}

func (_this *SimpleUnionClient) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*SimpleUnionClient)
	return ok && _this.Equals(other)
}

func (*SimpleUnionClient) String() string {
	return "simpleunioninternaldafny.SimpleUnionClient"
}

func Type_SimpleUnionClient_() _dafny.TypeDescriptor {
	return type_SimpleUnionClient_{}
}

type type_SimpleUnionClient_ struct {
}

func (_this type_SimpleUnionClient_) Default() interface{} {
	return (*SimpleUnionClient)(nil)
}

func (_this type_SimpleUnionClient_) String() string {
	return "simpleunioninternaldafny.SimpleUnionClient"
}
func (_this *SimpleUnionClient) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){simpleunioninternaldafnytypes.Companion_ISimpleUnionClient_.TraitID_}
}

var _ simpleunioninternaldafnytypes.ISimpleUnionClient = &SimpleUnionClient{}
var _ _dafny.TraitOffspring = &SimpleUnionClient{}

func (_this *SimpleUnionClient) Ctor__(config SimpleUnionImpl.Config) {
	{
		(_this)._config = config
	}
}
func (_this *SimpleUnionClient) GetUnion(input simpleunioninternaldafnytypes.GetUnionInput) Wrappers.Result {
	{
		var output Wrappers.Result = Wrappers.Companion_Result_.Default(simpleunioninternaldafnytypes.Companion_GetUnionOutput_.Default())
		_ = output
		var _out0 Wrappers.Result
		_ = _out0
		_out0 = SimpleUnionImpl.Companion_Default___.GetUnion((_this).Config(), input)
		output = _out0
		return output
	}
}
func (_this *SimpleUnionClient) GetKnownValueUnion(input simpleunioninternaldafnytypes.GetKnownValueUnionInput) Wrappers.Result {
	{
		var output Wrappers.Result = Wrappers.Companion_Result_.Default(simpleunioninternaldafnytypes.Companion_GetKnownValueUnionOutput_.Default())
		_ = output
		var _out1 Wrappers.Result
		_ = _out1
		_out1 = SimpleUnionImpl.Companion_Default___.GetKnownValueUnion((_this).Config(), input)
		output = _out1
		return output
	}
}
func (_this *SimpleUnionClient) Config() SimpleUnionImpl.Config {
	{
		return _this._config
	}
}

// End of class SimpleUnionClient
