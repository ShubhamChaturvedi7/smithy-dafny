// Package SimpleIntegerImplTest
// Dafny module SimpleIntegerImplTest compiled into Go

package SimpleIntegerImplTest

import (
	os "os"

	SimpleIntegerImpl "github.com/Smithy-dafny/TestModels/SimpleTypes/SimpleInteger/SimpleIntegerImpl"
	simpletypesintegerinternaldafny "github.com/Smithy-dafny/TestModels/SimpleTypes/SimpleInteger/simpletypesintegerinternaldafny"
	simpletypesintegerinternaldafnytypes "github.com/Smithy-dafny/TestModels/SimpleTypes/SimpleInteger/simpletypesintegerinternaldafnytypes"
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
var _ SimpleIntegerImpl.Dummy__

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
	return "SimpleIntegerImplTest.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) GetInteger() {
	var _4_client *simpletypesintegerinternaldafny.SimpleIntegerClient
	_ = _4_client
	var _5_valueOrError0 Wrappers.Result = Wrappers.Result{}
	_ = _5_valueOrError0
	var _out2 Wrappers.Result
	_ = _out2
	_out2 = simpletypesintegerinternaldafny.Companion_Default___.SimpleInteger(simpletypesintegerinternaldafny.Companion_Default___.DefaultSimpleIntegerConfig())
	_5_valueOrError0 = _out2
	if !(!((_5_valueOrError0).IsFailure())) {
		panic("test/SimpleIntegerImplTest.dfy(11,22): " + (_5_valueOrError0).String())
	}
	_4_client = (_5_valueOrError0).Extract().(*simpletypesintegerinternaldafny.SimpleIntegerClient)
	Companion_Default___.TestGetInteger(_4_client)
	Companion_Default___.TestGetIntegerKnownValue(_4_client)
	Companion_Default___.TestGetIntegerEdgeCases(_4_client)
}
func (_static *CompanionStruct_Default___) TestGetInteger(client simpletypesintegerinternaldafnytypes.ISimpleTypesIntegerClient) {
	var _6_ret simpletypesintegerinternaldafnytypes.GetIntegerOutput
	_ = _6_ret
	var _7_valueOrError0 Wrappers.Result = Wrappers.Companion_Result_.Default(simpletypesintegerinternaldafnytypes.Companion_GetIntegerOutput_.Default())
	_ = _7_valueOrError0
	var _out3 Wrappers.Result
	_ = _out3
	_out3 = (client).GetInteger(simpletypesintegerinternaldafnytypes.Companion_GetIntegerInput_.Create_GetIntegerInput_(Wrappers.Companion_Option_.Create_Some_(int32(1))))
	_7_valueOrError0 = _out3
	if !(!((_7_valueOrError0).IsFailure())) {
		panic("test/SimpleIntegerImplTest.dfy(22,19): " + (_7_valueOrError0).String())
	}
	_6_ret = (_7_valueOrError0).Extract().(simpletypesintegerinternaldafnytypes.GetIntegerOutput)
	if !((((_6_ret).Dtor_value()).UnwrapOr(int32(0)).(int32)) == (int32(1))) {
		panic("test/SimpleIntegerImplTest.dfy(23,8): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_dafny.Print(_6_ret)
}
func (_static *CompanionStruct_Default___) TestGetIntegerKnownValue(client simpletypesintegerinternaldafnytypes.ISimpleTypesIntegerClient) {
	var _8_ret simpletypesintegerinternaldafnytypes.GetIntegerOutput
	_ = _8_ret
	var _9_valueOrError0 Wrappers.Result = Wrappers.Companion_Result_.Default(simpletypesintegerinternaldafnytypes.Companion_GetIntegerOutput_.Default())
	_ = _9_valueOrError0
	var _out4 Wrappers.Result
	_ = _out4
	_out4 = (client).GetIntegerKnownValueTest(simpletypesintegerinternaldafnytypes.Companion_GetIntegerInput_.Create_GetIntegerInput_(Wrappers.Companion_Option_.Create_Some_(int32(20))))
	_9_valueOrError0 = _out4
	if !(!((_9_valueOrError0).IsFailure())) {
		panic("test/SimpleIntegerImplTest.dfy(32,19): " + (_9_valueOrError0).String())
	}
	_8_ret = (_9_valueOrError0).Extract().(simpletypesintegerinternaldafnytypes.GetIntegerOutput)
	if !((((_8_ret).Dtor_value()).UnwrapOr(int32(0)).(int32)) == (int32(20))) {
		panic("test/SimpleIntegerImplTest.dfy(33,8): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_dafny.Print(_8_ret)
}
func (_static *CompanionStruct_Default___) TestGetIntegerEdgeCases(client simpletypesintegerinternaldafnytypes.ISimpleTypesIntegerClient) {
	var _10_inputInteger _dafny.Sequence
	_ = _10_inputInteger
	_10_inputInteger = _dafny.SeqOf(int32(-1), int32(0), int32(1), ((StandardLibrary_UInt.Companion_Default___.INT32__MAX__LIMIT()).Minus(_dafny.One)).Int32(), (int32(0))-(((StandardLibrary_UInt.Companion_Default___.INT32__MAX__LIMIT()).Minus(_dafny.One)).Int32()))
	var _hi0 _dafny.Int = _dafny.IntOfUint32((_10_inputInteger).Cardinality())
	_ = _hi0
	for _11_i := _dafny.Zero; _11_i.Cmp(_hi0) < 0; _11_i = _11_i.Plus(_dafny.One) {
		var _12_integerValue int32
		_ = _12_integerValue
		_12_integerValue = (_10_inputInteger).Select((_11_i).Uint32()).(int32)
		_dafny.Print(_12_integerValue)
		var _13_ret simpletypesintegerinternaldafnytypes.GetIntegerOutput
		_ = _13_ret
		var _14_valueOrError0 Wrappers.Result = Wrappers.Companion_Result_.Default(simpletypesintegerinternaldafnytypes.Companion_GetIntegerOutput_.Default())
		_ = _14_valueOrError0
		var _out5 Wrappers.Result
		_ = _out5
		_out5 = (client).GetInteger(simpletypesintegerinternaldafnytypes.Companion_GetIntegerInput_.Create_GetIntegerInput_(Wrappers.Companion_Option_.Create_Some_(_12_integerValue)))
		_14_valueOrError0 = _out5
		if !(!((_14_valueOrError0).IsFailure())) {
			panic("test/SimpleIntegerImplTest.dfy(53,23): " + (_14_valueOrError0).String())
		}
		_13_ret = (_14_valueOrError0).Extract().(simpletypesintegerinternaldafnytypes.GetIntegerOutput)
		if !((((_13_ret).Dtor_value()).UnwrapOr(int32(0)).(int32)) == (_12_integerValue)) {
			panic("test/SimpleIntegerImplTest.dfy(54,12): " + (_dafny.SeqOfString("expectation violation")).String())
		}
		_dafny.Print(_13_ret)
	}
}

// End of class Default__
