// Package SimpleUnionImplTest
// Dafny module SimpleUnionImplTest compiled into Go

package SimpleUnionImplTest

import (
	os "os"

	SimpleUnionImpl "github.com/Smithy-dafny/TestModels/Union/SimpleUnionImpl"
	simpleunioninternaldafny "github.com/Smithy-dafny/TestModels/Union/simpleunioninternaldafny"
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
	return "SimpleUnionImplTest.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) TestUnion() {
	var _4_client *simpleunioninternaldafny.SimpleUnionClient
	_ = _4_client
	var _5_valueOrError0 Wrappers.Result = Wrappers.Result{}
	_ = _5_valueOrError0
	var _out2 Wrappers.Result
	_ = _out2
	_out2 = simpleunioninternaldafny.Companion_Default___.SimpleUnion(simpleunioninternaldafny.Companion_Default___.DefaultSimpleUnionConfig())
	_5_valueOrError0 = _out2
	if !(!((_5_valueOrError0).IsFailure())) {
		panic("test/SimpleUnionImplTest.dfy(10,22): " + (_5_valueOrError0).String())
	}
	_4_client = (_5_valueOrError0).Extract().(*simpleunioninternaldafny.SimpleUnionClient)
	Companion_Default___.TestMyUnionInteger(_4_client)
	Companion_Default___.TestMyUnionString(_4_client)
	Companion_Default___.TestKnownValueUnionString(_4_client)
}
func (_static *CompanionStruct_Default___) TestMyUnionInteger(client simpleunioninternaldafnytypes.ISimpleUnionClient) {
	var _6_ret simpleunioninternaldafnytypes.GetUnionOutput
	_ = _6_ret
	var _7_valueOrError0 Wrappers.Result = Wrappers.Companion_Result_.Default(simpleunioninternaldafnytypes.Companion_GetUnionOutput_.Default())
	_ = _7_valueOrError0
	var _out3 Wrappers.Result
	_ = _out3
	_out3 = (client).GetUnion(simpleunioninternaldafnytypes.Companion_GetUnionInput_.Create_GetUnionInput_(Wrappers.Companion_Option_.Create_Some_(simpleunioninternaldafnytypes.Companion_MyUnion_.Create_IntegerValue_(int32(100)))))
	_7_valueOrError0 = _out3
	if !(!((_7_valueOrError0).IsFailure())) {
		panic("test/SimpleUnionImplTest.dfy(21,19): " + (_7_valueOrError0).String())
	}
	_6_ret = (_7_valueOrError0).Extract().(simpleunioninternaldafnytypes.GetUnionOutput)
	if !(((_6_ret).Dtor_union()).Is_Some()) {
		panic("test/SimpleUnionImplTest.dfy(23,8): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	if !((((_6_ret).Dtor_union()).Dtor_value().(simpleunioninternaldafnytypes.MyUnion)).Is_IntegerValue()) {
		panic("test/SimpleUnionImplTest.dfy(24,8): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	if !(((((_6_ret).Dtor_union()).Dtor_value().(simpleunioninternaldafnytypes.MyUnion)).Dtor_IntegerValue()) == (int32(100))) {
		panic("test/SimpleUnionImplTest.dfy(25,8): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	if !(((((_6_ret).Dtor_union()).Dtor_value().(simpleunioninternaldafnytypes.MyUnion)).Is_StringValue()) == (false)) {
		panic("test/SimpleUnionImplTest.dfy(26,8): " + (_dafny.SeqOfString("expectation violation")).String())
	}
}
func (_static *CompanionStruct_Default___) TestMyUnionString(client simpleunioninternaldafnytypes.ISimpleUnionClient) {
	var _8_ret simpleunioninternaldafnytypes.GetUnionOutput
	_ = _8_ret
	var _9_valueOrError0 Wrappers.Result = Wrappers.Companion_Result_.Default(simpleunioninternaldafnytypes.Companion_GetUnionOutput_.Default())
	_ = _9_valueOrError0
	var _out4 Wrappers.Result
	_ = _out4
	_out4 = (client).GetUnion(simpleunioninternaldafnytypes.Companion_GetUnionInput_.Create_GetUnionInput_(Wrappers.Companion_Option_.Create_Some_(simpleunioninternaldafnytypes.Companion_MyUnion_.Create_StringValue_(_dafny.SeqOfString("TestString")))))
	_9_valueOrError0 = _out4
	if !(!((_9_valueOrError0).IsFailure())) {
		panic("test/SimpleUnionImplTest.dfy(34,19): " + (_9_valueOrError0).String())
	}
	_8_ret = (_9_valueOrError0).Extract().(simpleunioninternaldafnytypes.GetUnionOutput)
	if !(((_8_ret).Dtor_union()).Is_Some()) {
		panic("test/SimpleUnionImplTest.dfy(36,8): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	if !((((_8_ret).Dtor_union()).Dtor_value().(simpleunioninternaldafnytypes.MyUnion)).Is_StringValue()) {
		panic("test/SimpleUnionImplTest.dfy(37,8): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	if !(_dafny.Companion_Sequence_.Equal((((_8_ret).Dtor_union()).Dtor_value().(simpleunioninternaldafnytypes.MyUnion)).Dtor_StringValue(), _dafny.SeqOfString("TestString"))) {
		panic("test/SimpleUnionImplTest.dfy(38,8): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	if !(((((_8_ret).Dtor_union()).Dtor_value().(simpleunioninternaldafnytypes.MyUnion)).Is_IntegerValue()) == (false)) {
		panic("test/SimpleUnionImplTest.dfy(39,8): " + (_dafny.SeqOfString("expectation violation")).String())
	}
}
func (_static *CompanionStruct_Default___) TestKnownValueUnionString(client simpleunioninternaldafnytypes.ISimpleUnionClient) {
	var _10_ret simpleunioninternaldafnytypes.GetKnownValueUnionOutput
	_ = _10_ret
	var _11_valueOrError0 Wrappers.Result = Wrappers.Companion_Result_.Default(simpleunioninternaldafnytypes.Companion_GetKnownValueUnionOutput_.Default())
	_ = _11_valueOrError0
	var _out5 Wrappers.Result
	_ = _out5
	_out5 = (client).GetKnownValueUnion(simpleunioninternaldafnytypes.Companion_GetKnownValueUnionInput_.Create_GetKnownValueUnionInput_(Wrappers.Companion_Option_.Create_Some_(simpleunioninternaldafnytypes.Companion_KnownValueUnion_.Create_Value_(int32(10)))))
	_11_valueOrError0 = _out5
	if !(!((_11_valueOrError0).IsFailure())) {
		panic("test/SimpleUnionImplTest.dfy(47,19): " + (_11_valueOrError0).String())
	}
	_10_ret = (_11_valueOrError0).Extract().(simpleunioninternaldafnytypes.GetKnownValueUnionOutput)
	if !(((_10_ret).Dtor_union()).Is_Some()) {
		panic("test/SimpleUnionImplTest.dfy(49,8): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	if !(true) {
		panic("test/SimpleUnionImplTest.dfy(50,8): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	if !(((((_10_ret).Dtor_union()).Dtor_value().(simpleunioninternaldafnytypes.KnownValueUnion)).Dtor_Value()) == (int32(10))) {
		panic("test/SimpleUnionImplTest.dfy(51,8): " + (_dafny.SeqOfString("expectation violation")).String())
	}
}

// End of class Default__
