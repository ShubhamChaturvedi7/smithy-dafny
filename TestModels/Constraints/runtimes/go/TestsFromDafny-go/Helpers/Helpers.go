// Package Helpers
// Dafny module Helpers compiled into Go

package Helpers

import (
	os "os"

	SimpleConstraintsImpl "github.com/Smithy-dafny/TestModels/Constraints/SimpleConstraintsImpl"
	simpleconstraintsinternaldafnytypes "github.com/Smithy-dafny/TestModels/Constraints/simpleconstraintsinternaldafnytypes"
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
var _ SimpleConstraintsImpl.Dummy__

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
	return "Helpers.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) GetValidInput() simpleconstraintsinternaldafnytypes.GetConstraintsInput {
	return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_(Wrappers.Companion_Option_.Create_Some_(_dafny.SeqOfString("bw1and10")), Wrappers.Companion_Option_.Create_Some_(_dafny.SeqOfString("atleast1")), Wrappers.Companion_Option_.Create_Some_(_dafny.SeqOfString("leq10")), Wrappers.Companion_Option_.Create_Some_(_dafny.SeqOf(uint8(0), uint8(1), uint8(0), uint8(1))), Wrappers.Companion_Option_.Create_Some_(_dafny.SeqOf(uint8(0), uint8(1), uint8(0), uint8(1))), Wrappers.Companion_Option_.Create_Some_(_dafny.SeqOf(uint8(0), uint8(1), uint8(0), uint8(1))), Wrappers.Companion_Option_.Create_Some_(_dafny.SeqOf(_dafny.SeqOfString("00"), _dafny.SeqOfString("11"))), Wrappers.Companion_Option_.Create_Some_(_dafny.SeqOf(_dafny.SeqOfString("00"), _dafny.SeqOfString("11"))), Wrappers.Companion_Option_.Create_Some_(_dafny.SeqOf(_dafny.SeqOfString("00"), _dafny.SeqOfString("11"))), Wrappers.Companion_Option_.Create_Some_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOfString("0"), _dafny.SeqOfString("1")).UpdateUnsafe(_dafny.SeqOfString("2"), _dafny.SeqOfString("3"))), Wrappers.Companion_Option_.Create_Some_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOfString("0"), _dafny.SeqOfString("1")).UpdateUnsafe(_dafny.SeqOfString("2"), _dafny.SeqOfString("3"))), Wrappers.Companion_Option_.Create_Some_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOfString("0"), _dafny.SeqOfString("1")).UpdateUnsafe(_dafny.SeqOfString("2"), _dafny.SeqOfString("3"))), Wrappers.Companion_Option_.Create_Some_(int32(3)), Wrappers.Companion_Option_.Create_Some_(int64(3)), Wrappers.Companion_Option_.Create_Some_(int32(3)), Wrappers.Companion_Option_.Create_Some_(int32(3)), Wrappers.Companion_Option_.Create_Some_(Companion_Default___.PROVIDER__ID()), Wrappers.Companion_Option_.Create_Some_(_dafny.SeqOf(Companion_Default___.PROVIDER__ID(), Companion_Default___.PROVIDER__ID())))
}
func (_static *CompanionStruct_Default___) ValidInt32(x _dafny.Int) bool {
	return ((_dafny.IntOfInt64(-2147483648)).Cmp(x) <= 0) && ((x).Cmp(_dafny.IntOfInt64(2147483648)) < 0)
}
func (_static *CompanionStruct_Default___) ValidInt64(x _dafny.Int) bool {
	return ((_dafny.IntOfInt64(-9223372036854775808)).Cmp(x) <= 0) && ((x).Cmp(_dafny.IntOfString("9223372036854775808")) < 0)
}
func (_static *CompanionStruct_Default___) ForceUtf8Bytes(value _dafny.Sequence) _dafny.Sequence {
	var _21_myUtf8Bytes _dafny.Sequence = value
	_ = _21_myUtf8Bytes
	return _21_myUtf8Bytes
}
func (_static *CompanionStruct_Default___) ForceListOfUtf8Bytes(value _dafny.Sequence) _dafny.Sequence {
	var _22_myListOfUtf8Bytes _dafny.Sequence = value
	_ = _22_myListOfUtf8Bytes
	return _22_myListOfUtf8Bytes
}
func (_static *CompanionStruct_Default___) ForceLessThanTen(value _dafny.Int) int32 {
	var _23_v32 int32 = (value).Int32()
	_ = _23_v32
	var _24_myLessThanTen int32 = _23_v32
	_ = _24_myLessThanTen
	return _24_myLessThanTen
}
func (_static *CompanionStruct_Default___) ForceOneToTen(value _dafny.Int) int32 {
	var _25_v32 int32 = (value).Int32()
	_ = _25_v32
	var _26_myOneToTen int32 = _25_v32
	_ = _26_myOneToTen
	return _26_myOneToTen
}
func (_static *CompanionStruct_Default___) ForceTenToTen(value _dafny.Int) int64 {
	var _27_v64 int64 = (value).Int64()
	_ = _27_v64
	var _28_myTenToTen int64 = _27_v64
	_ = _28_myTenToTen
	return _28_myTenToTen
}
func (_static *CompanionStruct_Default___) ForceMyString(value _dafny.Sequence) _dafny.Sequence {
	var _29_myMyString _dafny.Sequence = value
	_ = _29_myMyString
	return _29_myMyString
}
func (_static *CompanionStruct_Default___) ForceNonEmptyString(value _dafny.Sequence) _dafny.Sequence {
	var _30_myNonEmptyString _dafny.Sequence = value
	_ = _30_myNonEmptyString
	return _30_myNonEmptyString
}
func (_static *CompanionStruct_Default___) ForceStringLessThanOrEqualToTen(value _dafny.Sequence) _dafny.Sequence {
	var _31_myStringLessThanOrEqualToTen _dafny.Sequence = value
	_ = _31_myStringLessThanOrEqualToTen
	return _31_myStringLessThanOrEqualToTen
}
func (_static *CompanionStruct_Default___) ForceMyBlob(value _dafny.Sequence) _dafny.Sequence {
	var _32_myMyBlob _dafny.Sequence = value
	_ = _32_myMyBlob
	return _32_myMyBlob
}
func (_static *CompanionStruct_Default___) ForceNonEmptyBlob(value _dafny.Sequence) _dafny.Sequence {
	var _33_myNonEmptyBlob _dafny.Sequence = value
	_ = _33_myNonEmptyBlob
	return _33_myNonEmptyBlob
}
func (_static *CompanionStruct_Default___) ForceBlobLessThanOrEqualToTen(value _dafny.Sequence) _dafny.Sequence {
	var _34_myBlobLessThanOrEqualToTen _dafny.Sequence = value
	_ = _34_myBlobLessThanOrEqualToTen
	return _34_myBlobLessThanOrEqualToTen
}
func (_static *CompanionStruct_Default___) ForceMyList(value _dafny.Sequence) _dafny.Sequence {
	var _35_myMyList _dafny.Sequence = value
	_ = _35_myMyList
	return _35_myMyList
}
func (_static *CompanionStruct_Default___) ForceNonEmptyList(value _dafny.Sequence) _dafny.Sequence {
	var _36_myNonEmptyList _dafny.Sequence = value
	_ = _36_myNonEmptyList
	return _36_myNonEmptyList
}
func (_static *CompanionStruct_Default___) ForceListLessThanOrEqualToTen(value _dafny.Sequence) _dafny.Sequence {
	var _37_myListLessThanOrEqualToTen _dafny.Sequence = value
	_ = _37_myListLessThanOrEqualToTen
	return _37_myListLessThanOrEqualToTen
}
func (_static *CompanionStruct_Default___) ForceMyMap(value _dafny.Map) _dafny.Map {
	var _38_myMyMap _dafny.Map = value
	_ = _38_myMyMap
	return _38_myMyMap
}
func (_static *CompanionStruct_Default___) ForceNonEmptyMap(value _dafny.Map) _dafny.Map {
	var _39_myNonEmptyMap _dafny.Map = value
	_ = _39_myNonEmptyMap
	return _39_myNonEmptyMap
}
func (_static *CompanionStruct_Default___) ForceMapLessThanOrEqualToTen(value _dafny.Map) _dafny.Map {
	var _40_myMapLessThanOrEqualToTen _dafny.Map = value
	_ = _40_myMapLessThanOrEqualToTen
	return _40_myMapLessThanOrEqualToTen
}
func (_static *CompanionStruct_Default___) ForceGreaterThanOne(value _dafny.Int) int32 {
	var _41_v32 int32 = (value).Int32()
	_ = _41_v32
	var _42_myGreaterThanOne int32 = _41_v32
	_ = _42_myGreaterThanOne
	return _42_myGreaterThanOne
}
func (_static *CompanionStruct_Default___) PROVIDER__ID() _dafny.Sequence {
	var _43_s _dafny.Sequence = _dafny.SeqOf(uint8(97), uint8(119), uint8(115), uint8(45), uint8(107), uint8(109), uint8(115))
	_ = _43_s
	return _43_s
}

// End of class Default__
