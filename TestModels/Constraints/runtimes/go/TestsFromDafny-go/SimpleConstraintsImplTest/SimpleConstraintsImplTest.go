// Package SimpleConstraintsImplTest
// Dafny module SimpleConstraintsImplTest compiled into Go

package SimpleConstraintsImplTest

import (
	os "os"

	Helpers "github.com/Smithy-dafny/TestModels/Constraints/Helpers"
	SimpleConstraintsImpl "github.com/Smithy-dafny/TestModels/Constraints/SimpleConstraintsImpl"
	simpleconstraintsinternaldafny "github.com/Smithy-dafny/TestModels/Constraints/simpleconstraintsinternaldafny"
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
var _ Helpers.Dummy__

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
	return "SimpleConstraintsImplTest.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) TestConstraints() {
	var _44_client *simpleconstraintsinternaldafny.SimpleConstraintsClient
	_ = _44_client
	var _45_valueOrError0 Wrappers.Result = Wrappers.Result{}
	_ = _45_valueOrError0
	var _out1 Wrappers.Result
	_ = _out1
	_out1 = simpleconstraintsinternaldafny.Companion_Default___.SimpleConstraints(simpleconstraintsinternaldafny.Companion_Default___.DefaultSimpleConstraintsConfig())
	_45_valueOrError0 = _out1
	if !(!((_45_valueOrError0).IsFailure())) {
		panic("test/SimpleConstraintsImplTest.dfy(16,22): " + (_45_valueOrError0).String())
	}
	_44_client = (_45_valueOrError0).Extract().(*simpleconstraintsinternaldafny.SimpleConstraintsClient)
	Companion_Default___.TestGetConstraintWithValidInputs(_44_client)
	Companion_Default___.TestGetConstraintWithInvalidMyString(_44_client)
}
func (_static *CompanionStruct_Default___) TestGetConstraintWithValidInputs(client simpleconstraintsinternaldafnytypes.ISimpleConstraintsClient) {
	var _46_input simpleconstraintsinternaldafnytypes.GetConstraintsInput
	_ = _46_input
	_46_input = Helpers.Companion_Default___.GetValidInput()
	var _47_ret Wrappers.Result
	_ = _47_ret
	var _out2 Wrappers.Result
	_ = _out2
	_out2 = (client).GetConstraints(_46_input)
	_47_ret = _out2
	if !((_47_ret).Is_Success()) {
		panic("test/SimpleConstraintsImplTest.dfy(28,6): " + (_dafny.SeqOfString("expectation violation")).String())
	}
}
func (_static *CompanionStruct_Default___) TestGetConstraintWithInvalidMyString(client simpleconstraintsinternaldafnytypes.ISimpleConstraintsClient) {
	var _48_input simpleconstraintsinternaldafnytypes.GetConstraintsInput
	_ = _48_input
	_48_input = Helpers.Companion_Default___.GetValidInput()
	_48_input = func(_pat_let0_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_49_dt__update__tmp_h0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let1_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_50_dt__update_hMyString_h0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_(_50_dt__update_hMyString_h0, (_49_dt__update__tmp_h0).Dtor_NonEmptyString(), (_49_dt__update__tmp_h0).Dtor_StringLessThanOrEqualToTen(), (_49_dt__update__tmp_h0).Dtor_MyBlob(), (_49_dt__update__tmp_h0).Dtor_NonEmptyBlob(), (_49_dt__update__tmp_h0).Dtor_BlobLessThanOrEqualToTen(), (_49_dt__update__tmp_h0).Dtor_MyList(), (_49_dt__update__tmp_h0).Dtor_NonEmptyList(), (_49_dt__update__tmp_h0).Dtor_ListLessThanOrEqualToTen(), (_49_dt__update__tmp_h0).Dtor_MyMap(), (_49_dt__update__tmp_h0).Dtor_NonEmptyMap(), (_49_dt__update__tmp_h0).Dtor_MapLessThanOrEqualToTen(), (_49_dt__update__tmp_h0).Dtor_OneToTen(), (_49_dt__update__tmp_h0).Dtor_myTenToTen(), (_49_dt__update__tmp_h0).Dtor_GreaterThanOne(), (_49_dt__update__tmp_h0).Dtor_LessThanTen(), (_49_dt__update__tmp_h0).Dtor_MyUtf8Bytes(), (_49_dt__update__tmp_h0).Dtor_MyListOfUtf8Bytes())
				}(_pat_let1_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceMyString(_dafny.SeqOfString("this string is too long"))))
		}(_pat_let0_0)
	}(_48_input)
	var _51_ret Wrappers.Result
	_ = _51_ret
	var _out3 Wrappers.Result
	_ = _out3
	_out3 = (client).GetConstraints(_48_input)
	_51_ret = _out3
	if !((_51_ret).Is_Success()) {
		panic("test/SimpleConstraintsImplTest.dfy(41,6): " + (_dafny.SeqOfString("expectation violation")).String())
	}
}

// End of class Default__
