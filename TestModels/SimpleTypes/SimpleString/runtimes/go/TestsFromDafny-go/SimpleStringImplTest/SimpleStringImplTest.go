// Package SimpleStringImplTest
// Dafny module SimpleStringImplTest compiled into Go

package SimpleStringImplTest

import (
	os "os"

	SimpleStringImpl "github.com/Smithy-dafny/TestModels/SimpleTypes/SimpleString/SimpleStringImpl"
	simpletypessmithystringinternaldafny "github.com/Smithy-dafny/TestModels/SimpleTypes/SimpleString/simpletypessmithystringinternaldafny"
	simpletypessmithystringinternaldafnytypes "github.com/Smithy-dafny/TestModels/SimpleTypes/SimpleString/simpletypessmithystringinternaldafnytypes"
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
var _ SimpleStringImpl.Dummy__

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
	return "SimpleStringImplTest.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) GetString() {
	var _5_client *simpletypessmithystringinternaldafny.SimpleStringClient
	_ = _5_client
	var _6_valueOrError0 Wrappers.Result = Wrappers.Result{}
	_ = _6_valueOrError0
	var _out3 Wrappers.Result
	_ = _out3
	_out3 = simpletypessmithystringinternaldafny.Companion_Default___.SimpleString(simpletypessmithystringinternaldafny.Companion_Default___.DefaultSimpleStringConfig())
	_6_valueOrError0 = _out3
	if !(!((_6_valueOrError0).IsFailure())) {
		panic("test/SimpleStringImplTest.dfy(10,22): " + (_6_valueOrError0).String())
	}
	_5_client = (_6_valueOrError0).Extract().(*simpletypessmithystringinternaldafny.SimpleStringClient)
	Companion_Default___.TestGetString(_5_client)
	Companion_Default___.TestGetStringKnownValue(_5_client)
	Companion_Default___.TestGetStringUTF8(_5_client)
}
func (_static *CompanionStruct_Default___) TestGetString(client simpletypessmithystringinternaldafnytypes.ISimpleTypesStringClient) {
	var _7_ret simpletypessmithystringinternaldafnytypes.GetStringOutput
	_ = _7_ret
	var _8_valueOrError0 Wrappers.Result = Wrappers.Companion_Result_.Default(simpletypessmithystringinternaldafnytypes.Companion_GetStringOutput_.Default())
	_ = _8_valueOrError0
	var _out4 Wrappers.Result
	_ = _out4
	_out4 = (client).GetString(simpletypessmithystringinternaldafnytypes.Companion_GetStringInput_.Create_GetStringInput_(Wrappers.Companion_Option_.Create_Some_(_dafny.SeqOfString("TEST_SIMPLE_STRING_VALUE"))))
	_8_valueOrError0 = _out4
	if !(!((_8_valueOrError0).IsFailure())) {
		panic("test/SimpleStringImplTest.dfy(21,19): " + (_8_valueOrError0).String())
	}
	_7_ret = (_8_valueOrError0).Extract().(simpletypessmithystringinternaldafnytypes.GetStringOutput)
	if !(_dafny.Companion_Sequence_.Equal(((_7_ret).Dtor_value()).UnwrapOr(_dafny.SeqOfString("")).(_dafny.Sequence), _dafny.SeqOfString("TEST_SIMPLE_STRING_VALUE"))) {
		panic("test/SimpleStringImplTest.dfy(22,8): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_dafny.Print(_7_ret)
}
func (_static *CompanionStruct_Default___) TestGetStringKnownValue(client simpletypessmithystringinternaldafnytypes.ISimpleTypesStringClient) {
	var _9_ret simpletypessmithystringinternaldafnytypes.GetStringOutput
	_ = _9_ret
	var _10_valueOrError0 Wrappers.Result = Wrappers.Companion_Result_.Default(simpletypessmithystringinternaldafnytypes.Companion_GetStringOutput_.Default())
	_ = _10_valueOrError0
	var _out5 Wrappers.Result
	_ = _out5
	_out5 = (client).GetStringKnownValue(simpletypessmithystringinternaldafnytypes.Companion_GetStringInput_.Create_GetStringInput_(Wrappers.Companion_Option_.Create_Some_(_dafny.SeqOfString("TEST_SIMPLE_STRING_KNOWN_VALUE"))))
	_10_valueOrError0 = _out5
	if !(!((_10_valueOrError0).IsFailure())) {
		panic("test/SimpleStringImplTest.dfy(30,19): " + (_10_valueOrError0).String())
	}
	_9_ret = (_10_valueOrError0).Extract().(simpletypessmithystringinternaldafnytypes.GetStringOutput)
	if !(_dafny.Companion_Sequence_.Equal(((_9_ret).Dtor_value()).UnwrapOr(_dafny.SeqOfString("")).(_dafny.Sequence), _dafny.SeqOfString("TEST_SIMPLE_STRING_KNOWN_VALUE"))) {
		panic("test/SimpleStringImplTest.dfy(31,8): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_dafny.Print(_9_ret)
}
func (_static *CompanionStruct_Default___) TestGetStringUTF8(client simpletypessmithystringinternaldafnytypes.ISimpleTypesStringClient) {
	var _11_utf8EncodedString _dafny.Sequence
	_ = _11_utf8EncodedString
	_11_utf8EncodedString = _dafny.SeqOfChars(2309, 2344, 2366, 2352)
	var _12_ret simpletypessmithystringinternaldafnytypes.GetStringOutput
	_ = _12_ret
	var _13_valueOrError0 Wrappers.Result = Wrappers.Companion_Result_.Default(simpletypessmithystringinternaldafnytypes.Companion_GetStringOutput_.Default())
	_ = _13_valueOrError0
	var _out6 Wrappers.Result
	_ = _out6
	_out6 = (client).GetStringUTF8(simpletypessmithystringinternaldafnytypes.Companion_GetStringInput_.Create_GetStringInput_(Wrappers.Companion_Option_.Create_Some_(_11_utf8EncodedString)))
	_13_valueOrError0 = _out6
	if !(!((_13_valueOrError0).IsFailure())) {
		panic("test/SimpleStringImplTest.dfy(41,19): " + (_13_valueOrError0).String())
	}
	_12_ret = (_13_valueOrError0).Extract().(simpletypessmithystringinternaldafnytypes.GetStringOutput)
	if !(_dafny.Companion_Sequence_.Equal(((_12_ret).Dtor_value()).UnwrapOr(_dafny.SeqOfString("")).(_dafny.Sequence), _11_utf8EncodedString)) {
		panic("test/SimpleStringImplTest.dfy(42,8): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_dafny.Print(_12_ret)
}

// End of class Default__
