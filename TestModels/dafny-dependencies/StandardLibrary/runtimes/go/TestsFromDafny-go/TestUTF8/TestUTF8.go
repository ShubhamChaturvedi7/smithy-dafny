// Package TestUTF8
// Dafny module TestUTF8 compiled into Go

package TestUTF8

import (
	StandardLibrary "StandardLibrary"
	StandardLibraryInterop "StandardLibraryInterop"
	StandardLibrary_UInt "StandardLibrary_UInt"
	UTF8 "UTF8"
	Wrappers "Wrappers"
	os "os"

	_System "github.com/dafny-lang/DafnyRuntimeGo/System_"
	_dafny "github.com/dafny-lang/DafnyRuntimeGo/dafny"
)

var _ = os.Args
var _ _dafny.Dummy__
var _ _System.Dummy__
var _ Wrappers.Dummy__
var _ StandardLibrary_UInt.Dummy__
var _ StandardLibrary.Dummy__
var _ StandardLibraryInterop.Dummy__

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
	return "TestUTF8.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) TestEncodeHappyCase() {
	var _0_unicodeString _dafny.Sequence
	_ = _0_unicodeString
	_0_unicodeString = _dafny.SeqOfChars(97, 98, 99, 774, 509, 946)
	var _1_expectedBytes _dafny.Sequence
	_ = _1_expectedBytes
	_1_expectedBytes = _dafny.SeqOf(uint8(97), uint8(98), uint8(99), uint8(204), uint8(134), uint8(199), uint8(189), uint8(206), uint8(178))
	var _2_encoded _dafny.Sequence
	_ = _2_encoded
	var _3_valueOrError0 Wrappers.Result = Wrappers.Companion_Result_.Default(UTF8.Companion_ValidUTF8Bytes_.Witness())
	_ = _3_valueOrError0
	_3_valueOrError0 = UTF8.Encode(_0_unicodeString)
	if !(!((_3_valueOrError0).IsFailure())) {
		panic("test/UTF8.dfy(14,19): " + (_3_valueOrError0).String())
	}
	_2_encoded = (_3_valueOrError0).Extract().(_dafny.Sequence)
	if !(_dafny.Companion_Sequence_.Equal(_1_expectedBytes, _2_encoded)) {
		panic("test/UTF8.dfy(15,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
}
func (_static *CompanionStruct_Default___) TestEncodeInvalidUnicode() {
	var _4_invalidUnicode _dafny.Sequence
	_ = _4_invalidUnicode
	_4_invalidUnicode = _dafny.SeqOfChars(97, 98, 99, 55296)
	var _5_encoded Wrappers.Result
	_ = _5_encoded
	_5_encoded = UTF8.Encode(_4_invalidUnicode)
	if !((_5_encoded).Is_Failure()) {
		panic("test/UTF8.dfy(22,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
}
func (_static *CompanionStruct_Default___) TestDecodeHappyCase() {
	var _6_unicodeBytes _dafny.Sequence
	_ = _6_unicodeBytes
	_6_unicodeBytes = _dafny.SeqOf(uint8(97), uint8(98), uint8(99), uint8(204), uint8(134), uint8(199), uint8(189), uint8(206), uint8(178))
	var _7_expectedString _dafny.Sequence
	_ = _7_expectedString
	_7_expectedString = _dafny.SeqOfChars(97, 98, 99, 774, 509, 946)
	var _8_decoded _dafny.Sequence
	_ = _8_decoded
	var _9_valueOrError0 Wrappers.Result = Wrappers.Companion_Result_.Default(_dafny.EmptySeq.SetString())
	_ = _9_valueOrError0
	_9_valueOrError0 = UTF8.Decode(_6_unicodeBytes)
	if !(!((_9_valueOrError0).IsFailure())) {
		panic("test/UTF8.dfy(32,19): " + (_9_valueOrError0).String())
	}
	_8_decoded = (_9_valueOrError0).Extract().(_dafny.Sequence)
	if !(_dafny.Companion_Sequence_.Equal(_7_expectedString, _8_decoded)) {
		panic("test/UTF8.dfy(33,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
}
func (_static *CompanionStruct_Default___) TestDecodeInvalidUnicode() {
	var _10_invalidUnicode _dafny.Sequence
	_ = _10_invalidUnicode
	_10_invalidUnicode = _dafny.SeqOf(uint8(97), uint8(98), uint8(99), uint8(237), uint8(160), uint8(128))
	if !(!(UTF8.Companion_Default___.ValidUTF8Seq(_10_invalidUnicode))) {
		panic("test/UTF8.dfy(39,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	if !((UTF8.Decode(_10_invalidUnicode)).Is_Failure()) {
		panic("test/UTF8.dfy(40,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
}
func (_static *CompanionStruct_Default___) Test1Byte() {
	var _11_decoded _dafny.Sequence
	_ = _11_decoded
	_11_decoded = _dafny.SeqOfChars(0)
	var _12_encoded _dafny.Sequence
	_ = _12_encoded
	var _13_valueOrError0 Wrappers.Result = Wrappers.Companion_Result_.Default(UTF8.Companion_ValidUTF8Bytes_.Witness())
	_ = _13_valueOrError0
	_13_valueOrError0 = UTF8.Encode(_11_decoded)
	if !(!((_13_valueOrError0).IsFailure())) {
		panic("test/UTF8.dfy(46,19): " + (_13_valueOrError0).String())
	}
	_12_encoded = (_13_valueOrError0).Extract().(_dafny.Sequence)
	if !(_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(uint8(0)), _12_encoded)) {
		panic("test/UTF8.dfy(47,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	if !(UTF8.Companion_Default___.Uses1Byte(_12_encoded)) {
		panic("test/UTF8.dfy(48,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	var _14_redecoded _dafny.Sequence
	_ = _14_redecoded
	var _15_valueOrError1 Wrappers.Result = Wrappers.Companion_Result_.Default(_dafny.EmptySeq.SetString())
	_ = _15_valueOrError1
	_15_valueOrError1 = UTF8.Decode(_12_encoded)
	if !(!((_15_valueOrError1).IsFailure())) {
		panic("test/UTF8.dfy(49,21): " + (_15_valueOrError1).String())
	}
	_14_redecoded = (_15_valueOrError1).Extract().(_dafny.Sequence)
	if !(_dafny.Companion_Sequence_.Equal(_11_decoded, _14_redecoded)) {
		panic("test/UTF8.dfy(50,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_11_decoded = _dafny.SeqOfChars(32)
	var _16_valueOrError2 Wrappers.Result = Wrappers.Companion_Result_.Default(UTF8.Companion_ValidUTF8Bytes_.Witness())
	_ = _16_valueOrError2
	_16_valueOrError2 = UTF8.Encode(_11_decoded)
	if !(!((_16_valueOrError2).IsFailure())) {
		panic("test/UTF8.dfy(54,15): " + (_16_valueOrError2).String())
	}
	_12_encoded = (_16_valueOrError2).Extract().(_dafny.Sequence)
	if !(_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(uint8(32)), _12_encoded)) {
		panic("test/UTF8.dfy(55,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	if !(UTF8.Companion_Default___.Uses1Byte(_12_encoded)) {
		panic("test/UTF8.dfy(56,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	var _17_valueOrError3 Wrappers.Result = Wrappers.Companion_Result_.Default(_dafny.EmptySeq.SetString())
	_ = _17_valueOrError3
	_17_valueOrError3 = UTF8.Decode(_12_encoded)
	if !(!((_17_valueOrError3).IsFailure())) {
		panic("test/UTF8.dfy(57,17): " + (_17_valueOrError3).String())
	}
	_14_redecoded = (_17_valueOrError3).Extract().(_dafny.Sequence)
	if !(_dafny.Companion_Sequence_.Equal(_11_decoded, _14_redecoded)) {
		panic("test/UTF8.dfy(58,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_11_decoded = _dafny.SeqOfString("$")
	var _18_valueOrError4 Wrappers.Result = Wrappers.Companion_Result_.Default(UTF8.Companion_ValidUTF8Bytes_.Witness())
	_ = _18_valueOrError4
	_18_valueOrError4 = UTF8.Encode(_11_decoded)
	if !(!((_18_valueOrError4).IsFailure())) {
		panic("test/UTF8.dfy(61,15): " + (_18_valueOrError4).String())
	}
	_12_encoded = (_18_valueOrError4).Extract().(_dafny.Sequence)
	if !(_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(uint8(36)), _12_encoded)) {
		panic("test/UTF8.dfy(62,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	if !(UTF8.Companion_Default___.Uses1Byte(_12_encoded)) {
		panic("test/UTF8.dfy(63,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	var _19_valueOrError5 Wrappers.Result = Wrappers.Companion_Result_.Default(_dafny.EmptySeq.SetString())
	_ = _19_valueOrError5
	_19_valueOrError5 = UTF8.Decode(_12_encoded)
	if !(!((_19_valueOrError5).IsFailure())) {
		panic("test/UTF8.dfy(64,17): " + (_19_valueOrError5).String())
	}
	_14_redecoded = (_19_valueOrError5).Extract().(_dafny.Sequence)
	if !(_dafny.Companion_Sequence_.Equal(_11_decoded, _14_redecoded)) {
		panic("test/UTF8.dfy(65,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_11_decoded = _dafny.SeqOfString("0")
	var _20_valueOrError6 Wrappers.Result = Wrappers.Companion_Result_.Default(UTF8.Companion_ValidUTF8Bytes_.Witness())
	_ = _20_valueOrError6
	_20_valueOrError6 = UTF8.Encode(_11_decoded)
	if !(!((_20_valueOrError6).IsFailure())) {
		panic("test/UTF8.dfy(68,15): " + (_20_valueOrError6).String())
	}
	_12_encoded = (_20_valueOrError6).Extract().(_dafny.Sequence)
	if !(_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(uint8(48)), _12_encoded)) {
		panic("test/UTF8.dfy(69,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	if !(UTF8.Companion_Default___.Uses1Byte(_12_encoded)) {
		panic("test/UTF8.dfy(70,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	var _21_valueOrError7 Wrappers.Result = Wrappers.Companion_Result_.Default(_dafny.EmptySeq.SetString())
	_ = _21_valueOrError7
	_21_valueOrError7 = UTF8.Decode(_12_encoded)
	if !(!((_21_valueOrError7).IsFailure())) {
		panic("test/UTF8.dfy(71,17): " + (_21_valueOrError7).String())
	}
	_14_redecoded = (_21_valueOrError7).Extract().(_dafny.Sequence)
	if !(_dafny.Companion_Sequence_.Equal(_11_decoded, _14_redecoded)) {
		panic("test/UTF8.dfy(72,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_11_decoded = _dafny.SeqOfString("A")
	var _22_valueOrError8 Wrappers.Result = Wrappers.Companion_Result_.Default(UTF8.Companion_ValidUTF8Bytes_.Witness())
	_ = _22_valueOrError8
	_22_valueOrError8 = UTF8.Encode(_11_decoded)
	if !(!((_22_valueOrError8).IsFailure())) {
		panic("test/UTF8.dfy(75,15): " + (_22_valueOrError8).String())
	}
	_12_encoded = (_22_valueOrError8).Extract().(_dafny.Sequence)
	if !(_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(uint8(65)), _12_encoded)) {
		panic("test/UTF8.dfy(76,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	if !(UTF8.Companion_Default___.Uses1Byte(_12_encoded)) {
		panic("test/UTF8.dfy(77,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	var _23_valueOrError9 Wrappers.Result = Wrappers.Companion_Result_.Default(_dafny.EmptySeq.SetString())
	_ = _23_valueOrError9
	_23_valueOrError9 = UTF8.Decode(_12_encoded)
	if !(!((_23_valueOrError9).IsFailure())) {
		panic("test/UTF8.dfy(78,17): " + (_23_valueOrError9).String())
	}
	_14_redecoded = (_23_valueOrError9).Extract().(_dafny.Sequence)
	if !(_dafny.Companion_Sequence_.Equal(_11_decoded, _14_redecoded)) {
		panic("test/UTF8.dfy(79,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_11_decoded = _dafny.SeqOfString("a")
	var _24_valueOrError10 Wrappers.Result = Wrappers.Companion_Result_.Default(UTF8.Companion_ValidUTF8Bytes_.Witness())
	_ = _24_valueOrError10
	_24_valueOrError10 = UTF8.Encode(_11_decoded)
	if !(!((_24_valueOrError10).IsFailure())) {
		panic("test/UTF8.dfy(82,15): " + (_24_valueOrError10).String())
	}
	_12_encoded = (_24_valueOrError10).Extract().(_dafny.Sequence)
	if !(_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(uint8(97)), _12_encoded)) {
		panic("test/UTF8.dfy(83,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	if !(UTF8.Companion_Default___.Uses1Byte(_12_encoded)) {
		panic("test/UTF8.dfy(84,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	var _25_valueOrError11 Wrappers.Result = Wrappers.Companion_Result_.Default(_dafny.EmptySeq.SetString())
	_ = _25_valueOrError11
	_25_valueOrError11 = UTF8.Decode(_12_encoded)
	if !(!((_25_valueOrError11).IsFailure())) {
		panic("test/UTF8.dfy(85,17): " + (_25_valueOrError11).String())
	}
	_14_redecoded = (_25_valueOrError11).Extract().(_dafny.Sequence)
	if !(_dafny.Companion_Sequence_.Equal(_11_decoded, _14_redecoded)) {
		panic("test/UTF8.dfy(86,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
}
func (_static *CompanionStruct_Default___) Test2Bytes() {
	var _26_decoded _dafny.Sequence
	_ = _26_decoded
	_26_decoded = _dafny.SeqOfChars(163)
	var _27_encoded _dafny.Sequence
	_ = _27_encoded
	var _28_valueOrError0 Wrappers.Result = Wrappers.Companion_Result_.Default(UTF8.Companion_ValidUTF8Bytes_.Witness())
	_ = _28_valueOrError0
	_28_valueOrError0 = UTF8.Encode(_26_decoded)
	if !(!((_28_valueOrError0).IsFailure())) {
		panic("test/UTF8.dfy(92,19): " + (_28_valueOrError0).String())
	}
	_27_encoded = (_28_valueOrError0).Extract().(_dafny.Sequence)
	if !(_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(uint8(194), uint8(163)), _27_encoded)) {
		panic("test/UTF8.dfy(93,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	if !(UTF8.Companion_Default___.Uses2Bytes(_27_encoded)) {
		panic("test/UTF8.dfy(94,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	var _29_redecoded _dafny.Sequence
	_ = _29_redecoded
	var _30_valueOrError1 Wrappers.Result = Wrappers.Companion_Result_.Default(_dafny.EmptySeq.SetString())
	_ = _30_valueOrError1
	_30_valueOrError1 = UTF8.Decode(_27_encoded)
	if !(!((_30_valueOrError1).IsFailure())) {
		panic("test/UTF8.dfy(95,21): " + (_30_valueOrError1).String())
	}
	_29_redecoded = (_30_valueOrError1).Extract().(_dafny.Sequence)
	if !(_dafny.Companion_Sequence_.Equal(_26_decoded, _29_redecoded)) {
		panic("test/UTF8.dfy(96,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_26_decoded = _dafny.SeqOfChars(169)
	var _31_valueOrError2 Wrappers.Result = Wrappers.Companion_Result_.Default(UTF8.Companion_ValidUTF8Bytes_.Witness())
	_ = _31_valueOrError2
	_31_valueOrError2 = UTF8.Encode(_26_decoded)
	if !(!((_31_valueOrError2).IsFailure())) {
		panic("test/UTF8.dfy(100,15): " + (_31_valueOrError2).String())
	}
	_27_encoded = (_31_valueOrError2).Extract().(_dafny.Sequence)
	if !(_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(uint8(194), uint8(169)), _27_encoded)) {
		panic("test/UTF8.dfy(101,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	if !(UTF8.Companion_Default___.Uses2Bytes(_27_encoded)) {
		panic("test/UTF8.dfy(102,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	var _32_valueOrError3 Wrappers.Result = Wrappers.Companion_Result_.Default(_dafny.EmptySeq.SetString())
	_ = _32_valueOrError3
	_32_valueOrError3 = UTF8.Decode(_27_encoded)
	if !(!((_32_valueOrError3).IsFailure())) {
		panic("test/UTF8.dfy(103,17): " + (_32_valueOrError3).String())
	}
	_29_redecoded = (_32_valueOrError3).Extract().(_dafny.Sequence)
	if !(_dafny.Companion_Sequence_.Equal(_26_decoded, _29_redecoded)) {
		panic("test/UTF8.dfy(104,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_26_decoded = _dafny.SeqOfChars(174)
	var _33_valueOrError4 Wrappers.Result = Wrappers.Companion_Result_.Default(UTF8.Companion_ValidUTF8Bytes_.Witness())
	_ = _33_valueOrError4
	_33_valueOrError4 = UTF8.Encode(_26_decoded)
	if !(!((_33_valueOrError4).IsFailure())) {
		panic("test/UTF8.dfy(108,15): " + (_33_valueOrError4).String())
	}
	_27_encoded = (_33_valueOrError4).Extract().(_dafny.Sequence)
	if !(_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(uint8(194), uint8(174)), _27_encoded)) {
		panic("test/UTF8.dfy(109,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	if !(UTF8.Companion_Default___.Uses2Bytes(_27_encoded)) {
		panic("test/UTF8.dfy(110,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	var _34_valueOrError5 Wrappers.Result = Wrappers.Companion_Result_.Default(_dafny.EmptySeq.SetString())
	_ = _34_valueOrError5
	_34_valueOrError5 = UTF8.Decode(_27_encoded)
	if !(!((_34_valueOrError5).IsFailure())) {
		panic("test/UTF8.dfy(111,17): " + (_34_valueOrError5).String())
	}
	_29_redecoded = (_34_valueOrError5).Extract().(_dafny.Sequence)
	if !(_dafny.Companion_Sequence_.Equal(_26_decoded, _29_redecoded)) {
		panic("test/UTF8.dfy(112,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_26_decoded = _dafny.SeqOfChars(960)
	var _35_valueOrError6 Wrappers.Result = Wrappers.Companion_Result_.Default(UTF8.Companion_ValidUTF8Bytes_.Witness())
	_ = _35_valueOrError6
	_35_valueOrError6 = UTF8.Encode(_26_decoded)
	if !(!((_35_valueOrError6).IsFailure())) {
		panic("test/UTF8.dfy(116,15): " + (_35_valueOrError6).String())
	}
	_27_encoded = (_35_valueOrError6).Extract().(_dafny.Sequence)
	if !(_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(uint8(207), uint8(128)), _27_encoded)) {
		panic("test/UTF8.dfy(117,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	if !(UTF8.Companion_Default___.Uses2Bytes(_27_encoded)) {
		panic("test/UTF8.dfy(118,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	var _36_valueOrError7 Wrappers.Result = Wrappers.Companion_Result_.Default(_dafny.EmptySeq.SetString())
	_ = _36_valueOrError7
	_36_valueOrError7 = UTF8.Decode(_27_encoded)
	if !(!((_36_valueOrError7).IsFailure())) {
		panic("test/UTF8.dfy(119,17): " + (_36_valueOrError7).String())
	}
	_29_redecoded = (_36_valueOrError7).Extract().(_dafny.Sequence)
	if !(_dafny.Companion_Sequence_.Equal(_26_decoded, _29_redecoded)) {
		panic("test/UTF8.dfy(120,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
}
func (_static *CompanionStruct_Default___) Test3Bytes() {
	var _37_decoded _dafny.Sequence
	_ = _37_decoded
	_37_decoded = _dafny.SeqOfChars(9094)
	var _38_encoded _dafny.Sequence
	_ = _38_encoded
	var _39_valueOrError0 Wrappers.Result = Wrappers.Companion_Result_.Default(UTF8.Companion_ValidUTF8Bytes_.Witness())
	_ = _39_valueOrError0
	_39_valueOrError0 = UTF8.Encode(_37_decoded)
	if !(!((_39_valueOrError0).IsFailure())) {
		panic("test/UTF8.dfy(126,19): " + (_39_valueOrError0).String())
	}
	_38_encoded = (_39_valueOrError0).Extract().(_dafny.Sequence)
	if !(_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(uint8(226), uint8(142), uint8(134)), _38_encoded)) {
		panic("test/UTF8.dfy(127,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	if !(UTF8.Companion_Default___.Uses3Bytes(_38_encoded)) {
		panic("test/UTF8.dfy(128,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	var _40_redecoded _dafny.Sequence
	_ = _40_redecoded
	var _41_valueOrError1 Wrappers.Result = Wrappers.Companion_Result_.Default(_dafny.EmptySeq.SetString())
	_ = _41_valueOrError1
	_41_valueOrError1 = UTF8.Decode(_38_encoded)
	if !(!((_41_valueOrError1).IsFailure())) {
		panic("test/UTF8.dfy(129,21): " + (_41_valueOrError1).String())
	}
	_40_redecoded = (_41_valueOrError1).Extract().(_dafny.Sequence)
	if !(_dafny.Companion_Sequence_.Equal(_37_decoded, _40_redecoded)) {
		panic("test/UTF8.dfy(130,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_37_decoded = _dafny.SeqOfChars(9095)
	var _42_valueOrError2 Wrappers.Result = Wrappers.Companion_Result_.Default(UTF8.Companion_ValidUTF8Bytes_.Witness())
	_ = _42_valueOrError2
	_42_valueOrError2 = UTF8.Encode(_37_decoded)
	if !(!((_42_valueOrError2).IsFailure())) {
		panic("test/UTF8.dfy(134,15): " + (_42_valueOrError2).String())
	}
	_38_encoded = (_42_valueOrError2).Extract().(_dafny.Sequence)
	if !(_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(uint8(226), uint8(142), uint8(135)), _38_encoded)) {
		panic("test/UTF8.dfy(135,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	if !(UTF8.Companion_Default___.Uses3Bytes(_38_encoded)) {
		panic("test/UTF8.dfy(136,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	var _43_valueOrError3 Wrappers.Result = Wrappers.Companion_Result_.Default(_dafny.EmptySeq.SetString())
	_ = _43_valueOrError3
	_43_valueOrError3 = UTF8.Decode(_38_encoded)
	if !(!((_43_valueOrError3).IsFailure())) {
		panic("test/UTF8.dfy(137,17): " + (_43_valueOrError3).String())
	}
	_40_redecoded = (_43_valueOrError3).Extract().(_dafny.Sequence)
	if !(_dafny.Companion_Sequence_.Equal(_37_decoded, _40_redecoded)) {
		panic("test/UTF8.dfy(138,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_37_decoded = _dafny.SeqOfChars(8987)
	var _44_valueOrError4 Wrappers.Result = Wrappers.Companion_Result_.Default(UTF8.Companion_ValidUTF8Bytes_.Witness())
	_ = _44_valueOrError4
	_44_valueOrError4 = UTF8.Encode(_37_decoded)
	if !(!((_44_valueOrError4).IsFailure())) {
		panic("test/UTF8.dfy(142,15): " + (_44_valueOrError4).String())
	}
	_38_encoded = (_44_valueOrError4).Extract().(_dafny.Sequence)
	if !(_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(uint8(226), uint8(140), uint8(155)), _38_encoded)) {
		panic("test/UTF8.dfy(143,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	if !(UTF8.Companion_Default___.Uses3Bytes(_38_encoded)) {
		panic("test/UTF8.dfy(144,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	var _45_valueOrError5 Wrappers.Result = Wrappers.Companion_Result_.Default(_dafny.EmptySeq.SetString())
	_ = _45_valueOrError5
	_45_valueOrError5 = UTF8.Decode(_38_encoded)
	if !(!((_45_valueOrError5).IsFailure())) {
		panic("test/UTF8.dfy(145,17): " + (_45_valueOrError5).String())
	}
	_40_redecoded = (_45_valueOrError5).Extract().(_dafny.Sequence)
	if !(_dafny.Companion_Sequence_.Equal(_37_decoded, _40_redecoded)) {
		panic("test/UTF8.dfy(146,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_37_decoded = _dafny.SeqOfChars(7544)
	var _46_valueOrError6 Wrappers.Result = Wrappers.Companion_Result_.Default(UTF8.Companion_ValidUTF8Bytes_.Witness())
	_ = _46_valueOrError6
	_46_valueOrError6 = UTF8.Encode(_37_decoded)
	if !(!((_46_valueOrError6).IsFailure())) {
		panic("test/UTF8.dfy(150,15): " + (_46_valueOrError6).String())
	}
	_38_encoded = (_46_valueOrError6).Extract().(_dafny.Sequence)
	if !(_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(uint8(225), uint8(181), uint8(184)), _38_encoded)) {
		panic("test/UTF8.dfy(151,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	if !(UTF8.Companion_Default___.Uses3Bytes(_38_encoded)) {
		panic("test/UTF8.dfy(152,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	var _47_valueOrError7 Wrappers.Result = Wrappers.Companion_Result_.Default(_dafny.EmptySeq.SetString())
	_ = _47_valueOrError7
	_47_valueOrError7 = UTF8.Decode(_38_encoded)
	if !(!((_47_valueOrError7).IsFailure())) {
		panic("test/UTF8.dfy(153,17): " + (_47_valueOrError7).String())
	}
	_40_redecoded = (_47_valueOrError7).Extract().(_dafny.Sequence)
	if !(_dafny.Companion_Sequence_.Equal(_37_decoded, _40_redecoded)) {
		panic("test/UTF8.dfy(154,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_37_decoded = _dafny.SeqOfChars(29483)
	var _48_valueOrError8 Wrappers.Result = Wrappers.Companion_Result_.Default(UTF8.Companion_ValidUTF8Bytes_.Witness())
	_ = _48_valueOrError8
	_48_valueOrError8 = UTF8.Encode(_37_decoded)
	if !(!((_48_valueOrError8).IsFailure())) {
		panic("test/UTF8.dfy(158,15): " + (_48_valueOrError8).String())
	}
	_38_encoded = (_48_valueOrError8).Extract().(_dafny.Sequence)
	if !(_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(uint8(231), uint8(140), uint8(171)), _38_encoded)) {
		panic("test/UTF8.dfy(159,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	if !(UTF8.Companion_Default___.Uses3Bytes(_38_encoded)) {
		panic("test/UTF8.dfy(160,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	var _49_valueOrError9 Wrappers.Result = Wrappers.Companion_Result_.Default(_dafny.EmptySeq.SetString())
	_ = _49_valueOrError9
	_49_valueOrError9 = UTF8.Decode(_38_encoded)
	if !(!((_49_valueOrError9).IsFailure())) {
		panic("test/UTF8.dfy(161,17): " + (_49_valueOrError9).String())
	}
	_40_redecoded = (_49_valueOrError9).Extract().(_dafny.Sequence)
	if !(_dafny.Companion_Sequence_.Equal(_37_decoded, _40_redecoded)) {
		panic("test/UTF8.dfy(162,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
}
func (_static *CompanionStruct_Default___) Test4Bytes() {
	var _50_decoded _dafny.Sequence
	_ = _50_decoded
	_50_decoded = _dafny.SeqOfChars(55304, 56320)
	var _51_encoded _dafny.Sequence
	_ = _51_encoded
	var _52_valueOrError0 Wrappers.Result = Wrappers.Companion_Result_.Default(UTF8.Companion_ValidUTF8Bytes_.Witness())
	_ = _52_valueOrError0
	_52_valueOrError0 = UTF8.Encode(_50_decoded)
	if !(!((_52_valueOrError0).IsFailure())) {
		panic("test/UTF8.dfy(168,19): " + (_52_valueOrError0).String())
	}
	_51_encoded = (_52_valueOrError0).Extract().(_dafny.Sequence)
	if !(_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(uint8(240), uint8(146), uint8(128), uint8(128)), _51_encoded)) {
		panic("test/UTF8.dfy(169,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	if !(UTF8.Companion_Default___.Uses4Bytes(_51_encoded)) {
		panic("test/UTF8.dfy(170,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	var _53_redecoded _dafny.Sequence
	_ = _53_redecoded
	var _54_valueOrError1 Wrappers.Result = Wrappers.Companion_Result_.Default(_dafny.EmptySeq.SetString())
	_ = _54_valueOrError1
	_54_valueOrError1 = UTF8.Decode(_51_encoded)
	if !(!((_54_valueOrError1).IsFailure())) {
		panic("test/UTF8.dfy(171,21): " + (_54_valueOrError1).String())
	}
	_53_redecoded = (_54_valueOrError1).Extract().(_dafny.Sequence)
	if !(_dafny.Companion_Sequence_.Equal(_50_decoded, _53_redecoded)) {
		panic("test/UTF8.dfy(172,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_50_decoded = _dafny.SeqOfChars(55349, 57281)
	var _55_valueOrError2 Wrappers.Result = Wrappers.Companion_Result_.Default(UTF8.Companion_ValidUTF8Bytes_.Witness())
	_ = _55_valueOrError2
	_55_valueOrError2 = UTF8.Encode(_50_decoded)
	if !(!((_55_valueOrError2).IsFailure())) {
		panic("test/UTF8.dfy(176,15): " + (_55_valueOrError2).String())
	}
	_51_encoded = (_55_valueOrError2).Extract().(_dafny.Sequence)
	if !(_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(uint8(240), uint8(157), uint8(159), uint8(129)), _51_encoded)) {
		panic("test/UTF8.dfy(177,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	if !(UTF8.Companion_Default___.Uses4Bytes(_51_encoded)) {
		panic("test/UTF8.dfy(178,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	var _56_valueOrError3 Wrappers.Result = Wrappers.Companion_Result_.Default(_dafny.EmptySeq.SetString())
	_ = _56_valueOrError3
	_56_valueOrError3 = UTF8.Decode(_51_encoded)
	if !(!((_56_valueOrError3).IsFailure())) {
		panic("test/UTF8.dfy(179,17): " + (_56_valueOrError3).String())
	}
	_53_redecoded = (_56_valueOrError3).Extract().(_dafny.Sequence)
	if !(_dafny.Companion_Sequence_.Equal(_50_decoded, _53_redecoded)) {
		panic("test/UTF8.dfy(180,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
}

// End of class Default__
