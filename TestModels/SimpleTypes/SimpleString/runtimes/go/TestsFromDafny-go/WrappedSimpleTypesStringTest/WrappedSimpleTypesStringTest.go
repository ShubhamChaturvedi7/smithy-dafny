// Package WrappedSimpleTypesStringTest
// Dafny module WrappedSimpleTypesStringTest compiled into Go

package WrappedSimpleTypesStringTest

import (
	os "os"

	SimpleStringImpl "github.com/Smithy-dafny/TestModels/SimpleTypes/SimpleString/SimpleStringImpl"
	SimpleStringImplTest "github.com/Smithy-dafny/TestModels/SimpleTypes/SimpleString/SimpleStringImplTest"
	simpletypessmithystringinternaldafnytypes "github.com/Smithy-dafny/TestModels/SimpleTypes/SimpleString/simpletypessmithystringinternaldafnytypes"
	simpletypessmithystringinternaldafnywrapped "github.com/Smithy-dafny/TestModels/SimpleTypes/SimpleString/simpletypessmithystringinternaldafnywrapped"
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
var _ SimpleStringImplTest.Dummy__

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
	return "WrappedSimpleTypesStringTest.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) GetString() {
	var _14_client simpletypessmithystringinternaldafnytypes.ISimpleTypesStringClient
	_ = _14_client
	var _15_valueOrError0 Wrappers.Result = Wrappers.Result{}
	_ = _15_valueOrError0
	var _out7 Wrappers.Result
	_ = _out7
	_out7 = simpletypessmithystringinternaldafnywrapped.WrappedSimpleString(simpletypessmithystringinternaldafnywrapped.Companion_Default___.WrappedDefaultSimpleStringConfig())
	_15_valueOrError0 = _out7
	if !(!((_15_valueOrError0).IsFailure())) {
		panic("test/WrappedSimpleStringTest.dfy(11,22): " + (_15_valueOrError0).String())
	}
	_14_client = simpletypessmithystringinternaldafnytypes.Companion_ISimpleTypesStringClient_.CastTo_((_15_valueOrError0).Extract())
	SimpleStringImplTest.Companion_Default___.TestGetString(_14_client)
	SimpleStringImplTest.Companion_Default___.TestGetStringKnownValue(_14_client)
	SimpleStringImplTest.Companion_Default___.TestGetStringUTF8(_14_client)
}

// End of class Default__
