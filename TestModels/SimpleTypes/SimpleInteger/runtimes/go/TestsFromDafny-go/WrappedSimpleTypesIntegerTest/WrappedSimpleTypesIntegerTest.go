// Package WrappedSimpleTypesIntegerTest
// Dafny module WrappedSimpleTypesIntegerTest compiled into Go

package WrappedSimpleTypesIntegerTest

import (
	os "os"

	SimpleIntegerImpl "github.com/Smithy-dafny/TestModels/SimpleTypes/SimpleInteger/SimpleIntegerImpl"
	SimpleIntegerImplTest "github.com/Smithy-dafny/TestModels/SimpleTypes/SimpleInteger/SimpleIntegerImplTest"
	simpletypesintegerinternaldafnytypes "github.com/Smithy-dafny/TestModels/SimpleTypes/SimpleInteger/simpletypesintegerinternaldafnytypes"
	simpletypesintegerinternaldafnywrapped "github.com/Smithy-dafny/TestModels/SimpleTypes/SimpleInteger/simpletypesintegerinternaldafnywrapped"
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
var _ SimpleIntegerImplTest.Dummy__

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
	return "WrappedSimpleTypesIntegerTest.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) GetInteger() {
	var _15_client simpletypesintegerinternaldafnytypes.ISimpleTypesIntegerClient
	_ = _15_client
	var _16_valueOrError0 Wrappers.Result = Wrappers.Result{}
	_ = _16_valueOrError0
	var _out6 Wrappers.Result
	_ = _out6
	_out6 = simpletypesintegerinternaldafnywrapped.WrappedSimpleInteger(simpletypesintegerinternaldafnywrapped.Companion_Default___.WrappedDefaultSimpleIntegerConfig())
	_16_valueOrError0 = _out6
	if !(!((_16_valueOrError0).IsFailure())) {
		panic("test/WrappedSimpleIntegerTest.dfy(11,22): " + (_16_valueOrError0).String())
	}
	_15_client = simpletypesintegerinternaldafnytypes.Companion_ISimpleTypesIntegerClient_.CastTo_((_16_valueOrError0).Extract())
	SimpleIntegerImplTest.Companion_Default___.TestGetInteger(_15_client)
	SimpleIntegerImplTest.Companion_Default___.TestGetIntegerKnownValue(_15_client)
	SimpleIntegerImplTest.Companion_Default___.TestGetIntegerEdgeCases(_15_client)
}

// End of class Default__
