// Package WrappedSimpleUnionTest
// Dafny module WrappedSimpleUnionTest compiled into Go

package WrappedSimpleUnionTest

import (
	os "os"

	SimpleUnionImpl "github.com/Smithy-dafny/TestModels/Union/SimpleUnionImpl"
	SimpleUnionImplTest "github.com/Smithy-dafny/TestModels/Union/SimpleUnionImplTest"
	simpleunioninternaldafnytypes "github.com/Smithy-dafny/TestModels/Union/simpleunioninternaldafnytypes"
	simpleunioninternaldafnywrapped "github.com/Smithy-dafny/TestModels/Union/simpleunioninternaldafnywrapped"
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
var _ SimpleUnionImplTest.Dummy__

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
	return "WrappedSimpleUnionTest.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) TestUnion() {
	var _12_client simpleunioninternaldafnytypes.ISimpleUnionClient
	_ = _12_client
	var _13_valueOrError0 Wrappers.Result = Wrappers.Result{}
	_ = _13_valueOrError0
	var _out6 Wrappers.Result
	_ = _out6
	_out6 = simpleunioninternaldafnywrapped.WrappedSimpleUnion(simpleunioninternaldafnywrapped.Companion_Default___.WrappedDefaultSimpleUnionConfig())
	_13_valueOrError0 = _out6
	if !(!((_13_valueOrError0).IsFailure())) {
		panic("test/WrappedSimpleUnionImplTest.dfy(11,22): " + (_13_valueOrError0).String())
	}
	_12_client = simpleunioninternaldafnytypes.Companion_ISimpleUnionClient_.CastTo_((_13_valueOrError0).Extract())
	SimpleUnionImplTest.Companion_Default___.TestMyUnionInteger(_12_client)
	SimpleUnionImplTest.Companion_Default___.TestMyUnionString(_12_client)
	SimpleUnionImplTest.Companion_Default___.TestKnownValueUnionString(_12_client)
}

// End of class Default__
