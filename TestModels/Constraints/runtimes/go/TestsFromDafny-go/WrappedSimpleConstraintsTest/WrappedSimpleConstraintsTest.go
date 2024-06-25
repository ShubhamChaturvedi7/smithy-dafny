// Package WrappedSimpleConstraintsTest
// Dafny module WrappedSimpleConstraintsTest compiled into Go

package WrappedSimpleConstraintsTest

import (
	os "os"

	Helpers "github.com/Smithy-dafny/TestModels/Constraints/Helpers"
	SimpleConstraintsImpl "github.com/Smithy-dafny/TestModels/Constraints/SimpleConstraintsImpl"
	SimpleConstraintsImplTest "github.com/Smithy-dafny/TestModels/Constraints/SimpleConstraintsImplTest"
	simpleconstraintsinternaldafnytypes "github.com/Smithy-dafny/TestModels/Constraints/simpleconstraintsinternaldafnytypes"
	simpleconstraintsinternaldafnywrapped "github.com/Smithy-dafny/TestModels/Constraints/simpleconstraintsinternaldafnywrapped"
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
var _ SimpleConstraintsImplTest.Dummy__

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
	return "WrappedSimpleConstraintsTest.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) TestConstraints() {
	var _52_client simpleconstraintsinternaldafnytypes.ISimpleConstraintsClient
	_ = _52_client
	var _53_valueOrError0 Wrappers.Result = Wrappers.Result{}
	_ = _53_valueOrError0
	var _out4 Wrappers.Result
	_ = _out4
	_out4 = simpleconstraintsinternaldafnywrapped.WrappedSimpleConstraints(simpleconstraintsinternaldafnywrapped.Companion_Default___.WrappedDefaultSimpleConstraintsConfig())
	_53_valueOrError0 = _out4
	if !(!((_53_valueOrError0).IsFailure())) {
		panic("test/WrappedSimpleConstraintsTest.dfy(14,18): " + (_53_valueOrError0).String())
	}
	_52_client = simpleconstraintsinternaldafnytypes.Companion_ISimpleConstraintsClient_.CastTo_((_53_valueOrError0).Extract())
	Companion_Default___.TestGetConstraintWithValidInputs(_52_client)
	Companion_Default___.TestGetConstraintWithMyString(_52_client)
	Companion_Default___.TestGetConstraintWithOneToTen(_52_client)
	Companion_Default___.TestGetConstraintWithTenToTen(_52_client)
	Companion_Default___.TestGetConstraintWithLessThanTen(_52_client)
	Companion_Default___.TestGetConstraintWithNonEmptyString(_52_client)
	Companion_Default___.TestGetConstraintWithStringLessThanOrEqualToTen(_52_client)
	Companion_Default___.TestGetConstraintWithMyBlob(_52_client)
	Companion_Default___.TestGetConstraintWithNonEmptyBlob(_52_client)
	Companion_Default___.TestGetConstraintWithBlobLessThanOrEqualToTen(_52_client)
	Companion_Default___.TestGetConstraintWithMyList(_52_client)
	Companion_Default___.TestGetConstraintWithNonEmptyList(_52_client)
	Companion_Default___.TestGetConstraintWithListLessThanOrEqualToTen(_52_client)
	Companion_Default___.TestGetConstraintWithMyMap(_52_client)
	Companion_Default___.TestGetConstraintWithNonEmptyMap(_52_client)
	Companion_Default___.TestGetConstraintWithMapLessThanOrEqualToTen(_52_client)
	Companion_Default___.TestGetConstraintWithGreaterThanOne(_52_client)
	Companion_Default___.TestGetConstraintWithUtf8Bytes(_52_client)
	Companion_Default___.TestGetConstraintWithListOfUtf8Bytes(_52_client)
}
func (_static *CompanionStruct_Default___) TestGetConstraintWithValidInputs(client simpleconstraintsinternaldafnytypes.ISimpleConstraintsClient) {
	var _54_input simpleconstraintsinternaldafnytypes.GetConstraintsInput
	_ = _54_input
	_54_input = Helpers.Companion_Default___.GetValidInput()
	var _55_ret Wrappers.Result
	_ = _55_ret
	var _out5 Wrappers.Result
	_ = _out5
	_out5 = (client).GetConstraints(_54_input)
	_55_ret = _out5
	if !((_55_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(43,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
}
func (_static *CompanionStruct_Default___) TestGetConstraintWithMyString(client simpleconstraintsinternaldafnytypes.ISimpleConstraintsClient) {
	var _56_input simpleconstraintsinternaldafnytypes.GetConstraintsInput
	_ = _56_input
	_56_input = Helpers.Companion_Default___.GetValidInput()
	_56_input = func(_pat_let2_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_57_dt__update__tmp_h0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let3_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_58_dt__update_hMyString_h0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_(_58_dt__update_hMyString_h0, (_57_dt__update__tmp_h0).Dtor_NonEmptyString(), (_57_dt__update__tmp_h0).Dtor_StringLessThanOrEqualToTen(), (_57_dt__update__tmp_h0).Dtor_MyBlob(), (_57_dt__update__tmp_h0).Dtor_NonEmptyBlob(), (_57_dt__update__tmp_h0).Dtor_BlobLessThanOrEqualToTen(), (_57_dt__update__tmp_h0).Dtor_MyList(), (_57_dt__update__tmp_h0).Dtor_NonEmptyList(), (_57_dt__update__tmp_h0).Dtor_ListLessThanOrEqualToTen(), (_57_dt__update__tmp_h0).Dtor_MyMap(), (_57_dt__update__tmp_h0).Dtor_NonEmptyMap(), (_57_dt__update__tmp_h0).Dtor_MapLessThanOrEqualToTen(), (_57_dt__update__tmp_h0).Dtor_OneToTen(), (_57_dt__update__tmp_h0).Dtor_myTenToTen(), (_57_dt__update__tmp_h0).Dtor_GreaterThanOne(), (_57_dt__update__tmp_h0).Dtor_LessThanTen(), (_57_dt__update__tmp_h0).Dtor_MyUtf8Bytes(), (_57_dt__update__tmp_h0).Dtor_MyListOfUtf8Bytes())
				}(_pat_let3_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceMyString(_dafny.SeqOfString("this string is really way too long"))))
		}(_pat_let2_0)
	}(_56_input)
	var _59_ret Wrappers.Result
	_ = _59_ret
	var _out6 Wrappers.Result
	_ = _out6
	_out6 = (client).GetConstraints(_56_input)
	_59_ret = _out6
	if !((_59_ret).Is_Failure()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(55,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_56_input = func(_pat_let4_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_60_dt__update__tmp_h1 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let5_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_61_dt__update_hMyString_h1 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_(_61_dt__update_hMyString_h1, (_60_dt__update__tmp_h1).Dtor_NonEmptyString(), (_60_dt__update__tmp_h1).Dtor_StringLessThanOrEqualToTen(), (_60_dt__update__tmp_h1).Dtor_MyBlob(), (_60_dt__update__tmp_h1).Dtor_NonEmptyBlob(), (_60_dt__update__tmp_h1).Dtor_BlobLessThanOrEqualToTen(), (_60_dt__update__tmp_h1).Dtor_MyList(), (_60_dt__update__tmp_h1).Dtor_NonEmptyList(), (_60_dt__update__tmp_h1).Dtor_ListLessThanOrEqualToTen(), (_60_dt__update__tmp_h1).Dtor_MyMap(), (_60_dt__update__tmp_h1).Dtor_NonEmptyMap(), (_60_dt__update__tmp_h1).Dtor_MapLessThanOrEqualToTen(), (_60_dt__update__tmp_h1).Dtor_OneToTen(), (_60_dt__update__tmp_h1).Dtor_myTenToTen(), (_60_dt__update__tmp_h1).Dtor_GreaterThanOne(), (_60_dt__update__tmp_h1).Dtor_LessThanTen(), (_60_dt__update__tmp_h1).Dtor_MyUtf8Bytes(), (_60_dt__update__tmp_h1).Dtor_MyListOfUtf8Bytes())
				}(_pat_let5_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceMyString(_dafny.SeqOfString("12345678901"))))
		}(_pat_let4_0)
	}(_56_input)
	var _out7 Wrappers.Result
	_ = _out7
	_out7 = (client).GetConstraints(_56_input)
	_59_ret = _out7
	if !((_59_ret).Is_Failure()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(59,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_56_input = func(_pat_let6_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_62_dt__update__tmp_h2 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let7_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_63_dt__update_hMyString_h2 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_(_63_dt__update_hMyString_h2, (_62_dt__update__tmp_h2).Dtor_NonEmptyString(), (_62_dt__update__tmp_h2).Dtor_StringLessThanOrEqualToTen(), (_62_dt__update__tmp_h2).Dtor_MyBlob(), (_62_dt__update__tmp_h2).Dtor_NonEmptyBlob(), (_62_dt__update__tmp_h2).Dtor_BlobLessThanOrEqualToTen(), (_62_dt__update__tmp_h2).Dtor_MyList(), (_62_dt__update__tmp_h2).Dtor_NonEmptyList(), (_62_dt__update__tmp_h2).Dtor_ListLessThanOrEqualToTen(), (_62_dt__update__tmp_h2).Dtor_MyMap(), (_62_dt__update__tmp_h2).Dtor_NonEmptyMap(), (_62_dt__update__tmp_h2).Dtor_MapLessThanOrEqualToTen(), (_62_dt__update__tmp_h2).Dtor_OneToTen(), (_62_dt__update__tmp_h2).Dtor_myTenToTen(), (_62_dt__update__tmp_h2).Dtor_GreaterThanOne(), (_62_dt__update__tmp_h2).Dtor_LessThanTen(), (_62_dt__update__tmp_h2).Dtor_MyUtf8Bytes(), (_62_dt__update__tmp_h2).Dtor_MyListOfUtf8Bytes())
				}(_pat_let7_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceMyString(_dafny.SeqOfString("1234567890"))))
		}(_pat_let6_0)
	}(_56_input)
	var _out8 Wrappers.Result
	_ = _out8
	_out8 = (client).GetConstraints(_56_input)
	_59_ret = _out8
	if !((_59_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(63,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_56_input = func(_pat_let8_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_64_dt__update__tmp_h3 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let9_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_65_dt__update_hMyString_h3 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_(_65_dt__update_hMyString_h3, (_64_dt__update__tmp_h3).Dtor_NonEmptyString(), (_64_dt__update__tmp_h3).Dtor_StringLessThanOrEqualToTen(), (_64_dt__update__tmp_h3).Dtor_MyBlob(), (_64_dt__update__tmp_h3).Dtor_NonEmptyBlob(), (_64_dt__update__tmp_h3).Dtor_BlobLessThanOrEqualToTen(), (_64_dt__update__tmp_h3).Dtor_MyList(), (_64_dt__update__tmp_h3).Dtor_NonEmptyList(), (_64_dt__update__tmp_h3).Dtor_ListLessThanOrEqualToTen(), (_64_dt__update__tmp_h3).Dtor_MyMap(), (_64_dt__update__tmp_h3).Dtor_NonEmptyMap(), (_64_dt__update__tmp_h3).Dtor_MapLessThanOrEqualToTen(), (_64_dt__update__tmp_h3).Dtor_OneToTen(), (_64_dt__update__tmp_h3).Dtor_myTenToTen(), (_64_dt__update__tmp_h3).Dtor_GreaterThanOne(), (_64_dt__update__tmp_h3).Dtor_LessThanTen(), (_64_dt__update__tmp_h3).Dtor_MyUtf8Bytes(), (_64_dt__update__tmp_h3).Dtor_MyListOfUtf8Bytes())
				}(_pat_let9_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceMyString(_dafny.SeqOfString("1"))))
		}(_pat_let8_0)
	}(_56_input)
	var _out9 Wrappers.Result
	_ = _out9
	_out9 = (client).GetConstraints(_56_input)
	_59_ret = _out9
	if !((_59_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(67,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_56_input = func(_pat_let10_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_66_dt__update__tmp_h4 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let11_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_67_dt__update_hMyString_h4 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_(_67_dt__update_hMyString_h4, (_66_dt__update__tmp_h4).Dtor_NonEmptyString(), (_66_dt__update__tmp_h4).Dtor_StringLessThanOrEqualToTen(), (_66_dt__update__tmp_h4).Dtor_MyBlob(), (_66_dt__update__tmp_h4).Dtor_NonEmptyBlob(), (_66_dt__update__tmp_h4).Dtor_BlobLessThanOrEqualToTen(), (_66_dt__update__tmp_h4).Dtor_MyList(), (_66_dt__update__tmp_h4).Dtor_NonEmptyList(), (_66_dt__update__tmp_h4).Dtor_ListLessThanOrEqualToTen(), (_66_dt__update__tmp_h4).Dtor_MyMap(), (_66_dt__update__tmp_h4).Dtor_NonEmptyMap(), (_66_dt__update__tmp_h4).Dtor_MapLessThanOrEqualToTen(), (_66_dt__update__tmp_h4).Dtor_OneToTen(), (_66_dt__update__tmp_h4).Dtor_myTenToTen(), (_66_dt__update__tmp_h4).Dtor_GreaterThanOne(), (_66_dt__update__tmp_h4).Dtor_LessThanTen(), (_66_dt__update__tmp_h4).Dtor_MyUtf8Bytes(), (_66_dt__update__tmp_h4).Dtor_MyListOfUtf8Bytes())
				}(_pat_let11_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceMyString(_dafny.SeqOfString(""))))
		}(_pat_let10_0)
	}(_56_input)
	var _out10 Wrappers.Result
	_ = _out10
	_out10 = (client).GetConstraints(_56_input)
	_59_ret = _out10
	if !((_59_ret).Is_Failure()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(71,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
}
func (_static *CompanionStruct_Default___) TestGetConstraintWithOneToTen(client simpleconstraintsinternaldafnytypes.ISimpleConstraintsClient) {
	var _68_input simpleconstraintsinternaldafnytypes.GetConstraintsInput
	_ = _68_input
	_68_input = Helpers.Companion_Default___.GetValidInput()
	_68_input = func(_pat_let12_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_69_dt__update__tmp_h0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let13_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_70_dt__update_hOneToTen_h0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_69_dt__update__tmp_h0).Dtor_MyString(), (_69_dt__update__tmp_h0).Dtor_NonEmptyString(), (_69_dt__update__tmp_h0).Dtor_StringLessThanOrEqualToTen(), (_69_dt__update__tmp_h0).Dtor_MyBlob(), (_69_dt__update__tmp_h0).Dtor_NonEmptyBlob(), (_69_dt__update__tmp_h0).Dtor_BlobLessThanOrEqualToTen(), (_69_dt__update__tmp_h0).Dtor_MyList(), (_69_dt__update__tmp_h0).Dtor_NonEmptyList(), (_69_dt__update__tmp_h0).Dtor_ListLessThanOrEqualToTen(), (_69_dt__update__tmp_h0).Dtor_MyMap(), (_69_dt__update__tmp_h0).Dtor_NonEmptyMap(), (_69_dt__update__tmp_h0).Dtor_MapLessThanOrEqualToTen(), _70_dt__update_hOneToTen_h0, (_69_dt__update__tmp_h0).Dtor_myTenToTen(), (_69_dt__update__tmp_h0).Dtor_GreaterThanOne(), (_69_dt__update__tmp_h0).Dtor_LessThanTen(), (_69_dt__update__tmp_h0).Dtor_MyUtf8Bytes(), (_69_dt__update__tmp_h0).Dtor_MyListOfUtf8Bytes())
				}(_pat_let13_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceOneToTen(_dafny.IntOfInt64(1000))))
		}(_pat_let12_0)
	}(_68_input)
	var _71_ret Wrappers.Result
	_ = _71_ret
	var _out11 Wrappers.Result
	_ = _out11
	_out11 = (client).GetConstraints(_68_input)
	_71_ret = _out11
	if !((_71_ret).Is_Failure()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(83,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_68_input = func(_pat_let14_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_72_dt__update__tmp_h1 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let15_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_73_dt__update_hOneToTen_h1 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_72_dt__update__tmp_h1).Dtor_MyString(), (_72_dt__update__tmp_h1).Dtor_NonEmptyString(), (_72_dt__update__tmp_h1).Dtor_StringLessThanOrEqualToTen(), (_72_dt__update__tmp_h1).Dtor_MyBlob(), (_72_dt__update__tmp_h1).Dtor_NonEmptyBlob(), (_72_dt__update__tmp_h1).Dtor_BlobLessThanOrEqualToTen(), (_72_dt__update__tmp_h1).Dtor_MyList(), (_72_dt__update__tmp_h1).Dtor_NonEmptyList(), (_72_dt__update__tmp_h1).Dtor_ListLessThanOrEqualToTen(), (_72_dt__update__tmp_h1).Dtor_MyMap(), (_72_dt__update__tmp_h1).Dtor_NonEmptyMap(), (_72_dt__update__tmp_h1).Dtor_MapLessThanOrEqualToTen(), _73_dt__update_hOneToTen_h1, (_72_dt__update__tmp_h1).Dtor_myTenToTen(), (_72_dt__update__tmp_h1).Dtor_GreaterThanOne(), (_72_dt__update__tmp_h1).Dtor_LessThanTen(), (_72_dt__update__tmp_h1).Dtor_MyUtf8Bytes(), (_72_dt__update__tmp_h1).Dtor_MyListOfUtf8Bytes())
				}(_pat_let15_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceOneToTen(_dafny.IntOfInt64(-1000))))
		}(_pat_let14_0)
	}(_68_input)
	var _out12 Wrappers.Result
	_ = _out12
	_out12 = (client).GetConstraints(_68_input)
	_71_ret = _out12
	if !((_71_ret).Is_Failure()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(87,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_68_input = func(_pat_let16_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_74_dt__update__tmp_h2 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let17_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_75_dt__update_hOneToTen_h2 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_74_dt__update__tmp_h2).Dtor_MyString(), (_74_dt__update__tmp_h2).Dtor_NonEmptyString(), (_74_dt__update__tmp_h2).Dtor_StringLessThanOrEqualToTen(), (_74_dt__update__tmp_h2).Dtor_MyBlob(), (_74_dt__update__tmp_h2).Dtor_NonEmptyBlob(), (_74_dt__update__tmp_h2).Dtor_BlobLessThanOrEqualToTen(), (_74_dt__update__tmp_h2).Dtor_MyList(), (_74_dt__update__tmp_h2).Dtor_NonEmptyList(), (_74_dt__update__tmp_h2).Dtor_ListLessThanOrEqualToTen(), (_74_dt__update__tmp_h2).Dtor_MyMap(), (_74_dt__update__tmp_h2).Dtor_NonEmptyMap(), (_74_dt__update__tmp_h2).Dtor_MapLessThanOrEqualToTen(), _75_dt__update_hOneToTen_h2, (_74_dt__update__tmp_h2).Dtor_myTenToTen(), (_74_dt__update__tmp_h2).Dtor_GreaterThanOne(), (_74_dt__update__tmp_h2).Dtor_LessThanTen(), (_74_dt__update__tmp_h2).Dtor_MyUtf8Bytes(), (_74_dt__update__tmp_h2).Dtor_MyListOfUtf8Bytes())
				}(_pat_let17_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceOneToTen(_dafny.Zero)))
		}(_pat_let16_0)
	}(_68_input)
	var _out13 Wrappers.Result
	_ = _out13
	_out13 = (client).GetConstraints(_68_input)
	_71_ret = _out13
	if !((_71_ret).Is_Failure()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(91,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_68_input = func(_pat_let18_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_76_dt__update__tmp_h3 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let19_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_77_dt__update_hOneToTen_h3 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_76_dt__update__tmp_h3).Dtor_MyString(), (_76_dt__update__tmp_h3).Dtor_NonEmptyString(), (_76_dt__update__tmp_h3).Dtor_StringLessThanOrEqualToTen(), (_76_dt__update__tmp_h3).Dtor_MyBlob(), (_76_dt__update__tmp_h3).Dtor_NonEmptyBlob(), (_76_dt__update__tmp_h3).Dtor_BlobLessThanOrEqualToTen(), (_76_dt__update__tmp_h3).Dtor_MyList(), (_76_dt__update__tmp_h3).Dtor_NonEmptyList(), (_76_dt__update__tmp_h3).Dtor_ListLessThanOrEqualToTen(), (_76_dt__update__tmp_h3).Dtor_MyMap(), (_76_dt__update__tmp_h3).Dtor_NonEmptyMap(), (_76_dt__update__tmp_h3).Dtor_MapLessThanOrEqualToTen(), _77_dt__update_hOneToTen_h3, (_76_dt__update__tmp_h3).Dtor_myTenToTen(), (_76_dt__update__tmp_h3).Dtor_GreaterThanOne(), (_76_dt__update__tmp_h3).Dtor_LessThanTen(), (_76_dt__update__tmp_h3).Dtor_MyUtf8Bytes(), (_76_dt__update__tmp_h3).Dtor_MyListOfUtf8Bytes())
				}(_pat_let19_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceOneToTen(_dafny.One)))
		}(_pat_let18_0)
	}(_68_input)
	var _out14 Wrappers.Result
	_ = _out14
	_out14 = (client).GetConstraints(_68_input)
	_71_ret = _out14
	if !((_71_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(95,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_68_input = func(_pat_let20_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_78_dt__update__tmp_h4 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let21_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_79_dt__update_hOneToTen_h4 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_78_dt__update__tmp_h4).Dtor_MyString(), (_78_dt__update__tmp_h4).Dtor_NonEmptyString(), (_78_dt__update__tmp_h4).Dtor_StringLessThanOrEqualToTen(), (_78_dt__update__tmp_h4).Dtor_MyBlob(), (_78_dt__update__tmp_h4).Dtor_NonEmptyBlob(), (_78_dt__update__tmp_h4).Dtor_BlobLessThanOrEqualToTen(), (_78_dt__update__tmp_h4).Dtor_MyList(), (_78_dt__update__tmp_h4).Dtor_NonEmptyList(), (_78_dt__update__tmp_h4).Dtor_ListLessThanOrEqualToTen(), (_78_dt__update__tmp_h4).Dtor_MyMap(), (_78_dt__update__tmp_h4).Dtor_NonEmptyMap(), (_78_dt__update__tmp_h4).Dtor_MapLessThanOrEqualToTen(), _79_dt__update_hOneToTen_h4, (_78_dt__update__tmp_h4).Dtor_myTenToTen(), (_78_dt__update__tmp_h4).Dtor_GreaterThanOne(), (_78_dt__update__tmp_h4).Dtor_LessThanTen(), (_78_dt__update__tmp_h4).Dtor_MyUtf8Bytes(), (_78_dt__update__tmp_h4).Dtor_MyListOfUtf8Bytes())
				}(_pat_let21_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceOneToTen(_dafny.IntOfInt64(10))))
		}(_pat_let20_0)
	}(_68_input)
	var _out15 Wrappers.Result
	_ = _out15
	_out15 = (client).GetConstraints(_68_input)
	_71_ret = _out15
	if !((_71_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(99,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_68_input = func(_pat_let22_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_80_dt__update__tmp_h5 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let23_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_81_dt__update_hOneToTen_h5 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_80_dt__update__tmp_h5).Dtor_MyString(), (_80_dt__update__tmp_h5).Dtor_NonEmptyString(), (_80_dt__update__tmp_h5).Dtor_StringLessThanOrEqualToTen(), (_80_dt__update__tmp_h5).Dtor_MyBlob(), (_80_dt__update__tmp_h5).Dtor_NonEmptyBlob(), (_80_dt__update__tmp_h5).Dtor_BlobLessThanOrEqualToTen(), (_80_dt__update__tmp_h5).Dtor_MyList(), (_80_dt__update__tmp_h5).Dtor_NonEmptyList(), (_80_dt__update__tmp_h5).Dtor_ListLessThanOrEqualToTen(), (_80_dt__update__tmp_h5).Dtor_MyMap(), (_80_dt__update__tmp_h5).Dtor_NonEmptyMap(), (_80_dt__update__tmp_h5).Dtor_MapLessThanOrEqualToTen(), _81_dt__update_hOneToTen_h5, (_80_dt__update__tmp_h5).Dtor_myTenToTen(), (_80_dt__update__tmp_h5).Dtor_GreaterThanOne(), (_80_dt__update__tmp_h5).Dtor_LessThanTen(), (_80_dt__update__tmp_h5).Dtor_MyUtf8Bytes(), (_80_dt__update__tmp_h5).Dtor_MyListOfUtf8Bytes())
				}(_pat_let23_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceOneToTen(_dafny.IntOfInt64(11))))
		}(_pat_let22_0)
	}(_68_input)
	var _out16 Wrappers.Result
	_ = _out16
	_out16 = (client).GetConstraints(_68_input)
	_71_ret = _out16
	if !((_71_ret).Is_Failure()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(103,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
}
func (_static *CompanionStruct_Default___) TestGetConstraintWithTenToTen(client simpleconstraintsinternaldafnytypes.ISimpleConstraintsClient) {
	var _82_input simpleconstraintsinternaldafnytypes.GetConstraintsInput
	_ = _82_input
	_82_input = Helpers.Companion_Default___.GetValidInput()
	_82_input = func(_pat_let24_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_83_dt__update__tmp_h0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let25_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_84_dt__update_hmyTenToTen_h0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_83_dt__update__tmp_h0).Dtor_MyString(), (_83_dt__update__tmp_h0).Dtor_NonEmptyString(), (_83_dt__update__tmp_h0).Dtor_StringLessThanOrEqualToTen(), (_83_dt__update__tmp_h0).Dtor_MyBlob(), (_83_dt__update__tmp_h0).Dtor_NonEmptyBlob(), (_83_dt__update__tmp_h0).Dtor_BlobLessThanOrEqualToTen(), (_83_dt__update__tmp_h0).Dtor_MyList(), (_83_dt__update__tmp_h0).Dtor_NonEmptyList(), (_83_dt__update__tmp_h0).Dtor_ListLessThanOrEqualToTen(), (_83_dt__update__tmp_h0).Dtor_MyMap(), (_83_dt__update__tmp_h0).Dtor_NonEmptyMap(), (_83_dt__update__tmp_h0).Dtor_MapLessThanOrEqualToTen(), (_83_dt__update__tmp_h0).Dtor_OneToTen(), _84_dt__update_hmyTenToTen_h0, (_83_dt__update__tmp_h0).Dtor_GreaterThanOne(), (_83_dt__update__tmp_h0).Dtor_LessThanTen(), (_83_dt__update__tmp_h0).Dtor_MyUtf8Bytes(), (_83_dt__update__tmp_h0).Dtor_MyListOfUtf8Bytes())
				}(_pat_let25_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceTenToTen(_dafny.IntOfInt64(1000))))
		}(_pat_let24_0)
	}(_82_input)
	var _85_ret Wrappers.Result
	_ = _85_ret
	var _out17 Wrappers.Result
	_ = _out17
	_out17 = (client).GetConstraints(_82_input)
	_85_ret = _out17
	if !((_85_ret).Is_Failure()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(115,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_82_input = func(_pat_let26_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_86_dt__update__tmp_h1 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let27_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_87_dt__update_hmyTenToTen_h1 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_86_dt__update__tmp_h1).Dtor_MyString(), (_86_dt__update__tmp_h1).Dtor_NonEmptyString(), (_86_dt__update__tmp_h1).Dtor_StringLessThanOrEqualToTen(), (_86_dt__update__tmp_h1).Dtor_MyBlob(), (_86_dt__update__tmp_h1).Dtor_NonEmptyBlob(), (_86_dt__update__tmp_h1).Dtor_BlobLessThanOrEqualToTen(), (_86_dt__update__tmp_h1).Dtor_MyList(), (_86_dt__update__tmp_h1).Dtor_NonEmptyList(), (_86_dt__update__tmp_h1).Dtor_ListLessThanOrEqualToTen(), (_86_dt__update__tmp_h1).Dtor_MyMap(), (_86_dt__update__tmp_h1).Dtor_NonEmptyMap(), (_86_dt__update__tmp_h1).Dtor_MapLessThanOrEqualToTen(), (_86_dt__update__tmp_h1).Dtor_OneToTen(), _87_dt__update_hmyTenToTen_h1, (_86_dt__update__tmp_h1).Dtor_GreaterThanOne(), (_86_dt__update__tmp_h1).Dtor_LessThanTen(), (_86_dt__update__tmp_h1).Dtor_MyUtf8Bytes(), (_86_dt__update__tmp_h1).Dtor_MyListOfUtf8Bytes())
				}(_pat_let27_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceTenToTen(_dafny.IntOfInt64(-1000))))
		}(_pat_let26_0)
	}(_82_input)
	var _out18 Wrappers.Result
	_ = _out18
	_out18 = (client).GetConstraints(_82_input)
	_85_ret = _out18
	if !((_85_ret).Is_Failure()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(119,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_82_input = func(_pat_let28_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_88_dt__update__tmp_h2 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let29_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_89_dt__update_hmyTenToTen_h2 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_88_dt__update__tmp_h2).Dtor_MyString(), (_88_dt__update__tmp_h2).Dtor_NonEmptyString(), (_88_dt__update__tmp_h2).Dtor_StringLessThanOrEqualToTen(), (_88_dt__update__tmp_h2).Dtor_MyBlob(), (_88_dt__update__tmp_h2).Dtor_NonEmptyBlob(), (_88_dt__update__tmp_h2).Dtor_BlobLessThanOrEqualToTen(), (_88_dt__update__tmp_h2).Dtor_MyList(), (_88_dt__update__tmp_h2).Dtor_NonEmptyList(), (_88_dt__update__tmp_h2).Dtor_ListLessThanOrEqualToTen(), (_88_dt__update__tmp_h2).Dtor_MyMap(), (_88_dt__update__tmp_h2).Dtor_NonEmptyMap(), (_88_dt__update__tmp_h2).Dtor_MapLessThanOrEqualToTen(), (_88_dt__update__tmp_h2).Dtor_OneToTen(), _89_dt__update_hmyTenToTen_h2, (_88_dt__update__tmp_h2).Dtor_GreaterThanOne(), (_88_dt__update__tmp_h2).Dtor_LessThanTen(), (_88_dt__update__tmp_h2).Dtor_MyUtf8Bytes(), (_88_dt__update__tmp_h2).Dtor_MyListOfUtf8Bytes())
				}(_pat_let29_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceTenToTen(_dafny.IntOfInt64(-11))))
		}(_pat_let28_0)
	}(_82_input)
	var _out19 Wrappers.Result
	_ = _out19
	_out19 = (client).GetConstraints(_82_input)
	_85_ret = _out19
	if !((_85_ret).Is_Failure()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(123,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_82_input = func(_pat_let30_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_90_dt__update__tmp_h3 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let31_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_91_dt__update_hmyTenToTen_h3 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_90_dt__update__tmp_h3).Dtor_MyString(), (_90_dt__update__tmp_h3).Dtor_NonEmptyString(), (_90_dt__update__tmp_h3).Dtor_StringLessThanOrEqualToTen(), (_90_dt__update__tmp_h3).Dtor_MyBlob(), (_90_dt__update__tmp_h3).Dtor_NonEmptyBlob(), (_90_dt__update__tmp_h3).Dtor_BlobLessThanOrEqualToTen(), (_90_dt__update__tmp_h3).Dtor_MyList(), (_90_dt__update__tmp_h3).Dtor_NonEmptyList(), (_90_dt__update__tmp_h3).Dtor_ListLessThanOrEqualToTen(), (_90_dt__update__tmp_h3).Dtor_MyMap(), (_90_dt__update__tmp_h3).Dtor_NonEmptyMap(), (_90_dt__update__tmp_h3).Dtor_MapLessThanOrEqualToTen(), (_90_dt__update__tmp_h3).Dtor_OneToTen(), _91_dt__update_hmyTenToTen_h3, (_90_dt__update__tmp_h3).Dtor_GreaterThanOne(), (_90_dt__update__tmp_h3).Dtor_LessThanTen(), (_90_dt__update__tmp_h3).Dtor_MyUtf8Bytes(), (_90_dt__update__tmp_h3).Dtor_MyListOfUtf8Bytes())
				}(_pat_let31_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceTenToTen(_dafny.IntOfInt64(-10))))
		}(_pat_let30_0)
	}(_82_input)
	var _out20 Wrappers.Result
	_ = _out20
	_out20 = (client).GetConstraints(_82_input)
	_85_ret = _out20
	if !((_85_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(127,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_82_input = func(_pat_let32_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_92_dt__update__tmp_h4 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let33_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_93_dt__update_hmyTenToTen_h4 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_92_dt__update__tmp_h4).Dtor_MyString(), (_92_dt__update__tmp_h4).Dtor_NonEmptyString(), (_92_dt__update__tmp_h4).Dtor_StringLessThanOrEqualToTen(), (_92_dt__update__tmp_h4).Dtor_MyBlob(), (_92_dt__update__tmp_h4).Dtor_NonEmptyBlob(), (_92_dt__update__tmp_h4).Dtor_BlobLessThanOrEqualToTen(), (_92_dt__update__tmp_h4).Dtor_MyList(), (_92_dt__update__tmp_h4).Dtor_NonEmptyList(), (_92_dt__update__tmp_h4).Dtor_ListLessThanOrEqualToTen(), (_92_dt__update__tmp_h4).Dtor_MyMap(), (_92_dt__update__tmp_h4).Dtor_NonEmptyMap(), (_92_dt__update__tmp_h4).Dtor_MapLessThanOrEqualToTen(), (_92_dt__update__tmp_h4).Dtor_OneToTen(), _93_dt__update_hmyTenToTen_h4, (_92_dt__update__tmp_h4).Dtor_GreaterThanOne(), (_92_dt__update__tmp_h4).Dtor_LessThanTen(), (_92_dt__update__tmp_h4).Dtor_MyUtf8Bytes(), (_92_dt__update__tmp_h4).Dtor_MyListOfUtf8Bytes())
				}(_pat_let33_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceTenToTen(_dafny.Zero)))
		}(_pat_let32_0)
	}(_82_input)
	var _out21 Wrappers.Result
	_ = _out21
	_out21 = (client).GetConstraints(_82_input)
	_85_ret = _out21
	if !((_85_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(131,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_82_input = func(_pat_let34_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_94_dt__update__tmp_h5 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let35_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_95_dt__update_hmyTenToTen_h5 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_94_dt__update__tmp_h5).Dtor_MyString(), (_94_dt__update__tmp_h5).Dtor_NonEmptyString(), (_94_dt__update__tmp_h5).Dtor_StringLessThanOrEqualToTen(), (_94_dt__update__tmp_h5).Dtor_MyBlob(), (_94_dt__update__tmp_h5).Dtor_NonEmptyBlob(), (_94_dt__update__tmp_h5).Dtor_BlobLessThanOrEqualToTen(), (_94_dt__update__tmp_h5).Dtor_MyList(), (_94_dt__update__tmp_h5).Dtor_NonEmptyList(), (_94_dt__update__tmp_h5).Dtor_ListLessThanOrEqualToTen(), (_94_dt__update__tmp_h5).Dtor_MyMap(), (_94_dt__update__tmp_h5).Dtor_NonEmptyMap(), (_94_dt__update__tmp_h5).Dtor_MapLessThanOrEqualToTen(), (_94_dt__update__tmp_h5).Dtor_OneToTen(), _95_dt__update_hmyTenToTen_h5, (_94_dt__update__tmp_h5).Dtor_GreaterThanOne(), (_94_dt__update__tmp_h5).Dtor_LessThanTen(), (_94_dt__update__tmp_h5).Dtor_MyUtf8Bytes(), (_94_dt__update__tmp_h5).Dtor_MyListOfUtf8Bytes())
				}(_pat_let35_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceTenToTen(_dafny.IntOfInt64(10))))
		}(_pat_let34_0)
	}(_82_input)
	var _out22 Wrappers.Result
	_ = _out22
	_out22 = (client).GetConstraints(_82_input)
	_85_ret = _out22
	if !((_85_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(135,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_82_input = func(_pat_let36_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_96_dt__update__tmp_h6 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let37_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_97_dt__update_hmyTenToTen_h6 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_96_dt__update__tmp_h6).Dtor_MyString(), (_96_dt__update__tmp_h6).Dtor_NonEmptyString(), (_96_dt__update__tmp_h6).Dtor_StringLessThanOrEqualToTen(), (_96_dt__update__tmp_h6).Dtor_MyBlob(), (_96_dt__update__tmp_h6).Dtor_NonEmptyBlob(), (_96_dt__update__tmp_h6).Dtor_BlobLessThanOrEqualToTen(), (_96_dt__update__tmp_h6).Dtor_MyList(), (_96_dt__update__tmp_h6).Dtor_NonEmptyList(), (_96_dt__update__tmp_h6).Dtor_ListLessThanOrEqualToTen(), (_96_dt__update__tmp_h6).Dtor_MyMap(), (_96_dt__update__tmp_h6).Dtor_NonEmptyMap(), (_96_dt__update__tmp_h6).Dtor_MapLessThanOrEqualToTen(), (_96_dt__update__tmp_h6).Dtor_OneToTen(), _97_dt__update_hmyTenToTen_h6, (_96_dt__update__tmp_h6).Dtor_GreaterThanOne(), (_96_dt__update__tmp_h6).Dtor_LessThanTen(), (_96_dt__update__tmp_h6).Dtor_MyUtf8Bytes(), (_96_dt__update__tmp_h6).Dtor_MyListOfUtf8Bytes())
				}(_pat_let37_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceTenToTen(_dafny.IntOfInt64(11))))
		}(_pat_let36_0)
	}(_82_input)
	var _out23 Wrappers.Result
	_ = _out23
	_out23 = (client).GetConstraints(_82_input)
	_85_ret = _out23
	if !((_85_ret).Is_Failure()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(139,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
}
func (_static *CompanionStruct_Default___) TestGetConstraintWithLessThanTen(client simpleconstraintsinternaldafnytypes.ISimpleConstraintsClient) {
	var _98_input simpleconstraintsinternaldafnytypes.GetConstraintsInput
	_ = _98_input
	_98_input = Helpers.Companion_Default___.GetValidInput()
	_98_input = func(_pat_let38_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_99_dt__update__tmp_h0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let39_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_100_dt__update_hLessThanTen_h0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_99_dt__update__tmp_h0).Dtor_MyString(), (_99_dt__update__tmp_h0).Dtor_NonEmptyString(), (_99_dt__update__tmp_h0).Dtor_StringLessThanOrEqualToTen(), (_99_dt__update__tmp_h0).Dtor_MyBlob(), (_99_dt__update__tmp_h0).Dtor_NonEmptyBlob(), (_99_dt__update__tmp_h0).Dtor_BlobLessThanOrEqualToTen(), (_99_dt__update__tmp_h0).Dtor_MyList(), (_99_dt__update__tmp_h0).Dtor_NonEmptyList(), (_99_dt__update__tmp_h0).Dtor_ListLessThanOrEqualToTen(), (_99_dt__update__tmp_h0).Dtor_MyMap(), (_99_dt__update__tmp_h0).Dtor_NonEmptyMap(), (_99_dt__update__tmp_h0).Dtor_MapLessThanOrEqualToTen(), (_99_dt__update__tmp_h0).Dtor_OneToTen(), (_99_dt__update__tmp_h0).Dtor_myTenToTen(), (_99_dt__update__tmp_h0).Dtor_GreaterThanOne(), _100_dt__update_hLessThanTen_h0, (_99_dt__update__tmp_h0).Dtor_MyUtf8Bytes(), (_99_dt__update__tmp_h0).Dtor_MyListOfUtf8Bytes())
				}(_pat_let39_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceLessThanTen(_dafny.IntOfInt64(1000))))
		}(_pat_let38_0)
	}(_98_input)
	var _101_ret Wrappers.Result
	_ = _101_ret
	var _out24 Wrappers.Result
	_ = _out24
	_out24 = (client).GetConstraints(_98_input)
	_101_ret = _out24
	if !((_101_ret).Is_Failure()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(151,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_98_input = func(_pat_let40_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_102_dt__update__tmp_h1 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let41_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_103_dt__update_hLessThanTen_h1 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_102_dt__update__tmp_h1).Dtor_MyString(), (_102_dt__update__tmp_h1).Dtor_NonEmptyString(), (_102_dt__update__tmp_h1).Dtor_StringLessThanOrEqualToTen(), (_102_dt__update__tmp_h1).Dtor_MyBlob(), (_102_dt__update__tmp_h1).Dtor_NonEmptyBlob(), (_102_dt__update__tmp_h1).Dtor_BlobLessThanOrEqualToTen(), (_102_dt__update__tmp_h1).Dtor_MyList(), (_102_dt__update__tmp_h1).Dtor_NonEmptyList(), (_102_dt__update__tmp_h1).Dtor_ListLessThanOrEqualToTen(), (_102_dt__update__tmp_h1).Dtor_MyMap(), (_102_dt__update__tmp_h1).Dtor_NonEmptyMap(), (_102_dt__update__tmp_h1).Dtor_MapLessThanOrEqualToTen(), (_102_dt__update__tmp_h1).Dtor_OneToTen(), (_102_dt__update__tmp_h1).Dtor_myTenToTen(), (_102_dt__update__tmp_h1).Dtor_GreaterThanOne(), _103_dt__update_hLessThanTen_h1, (_102_dt__update__tmp_h1).Dtor_MyUtf8Bytes(), (_102_dt__update__tmp_h1).Dtor_MyListOfUtf8Bytes())
				}(_pat_let41_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceLessThanTen(_dafny.IntOfInt64(-1000))))
		}(_pat_let40_0)
	}(_98_input)
	var _out25 Wrappers.Result
	_ = _out25
	_out25 = (client).GetConstraints(_98_input)
	_101_ret = _out25
	if !((_101_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(155,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_98_input = func(_pat_let42_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_104_dt__update__tmp_h2 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let43_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_105_dt__update_hLessThanTen_h2 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_104_dt__update__tmp_h2).Dtor_MyString(), (_104_dt__update__tmp_h2).Dtor_NonEmptyString(), (_104_dt__update__tmp_h2).Dtor_StringLessThanOrEqualToTen(), (_104_dt__update__tmp_h2).Dtor_MyBlob(), (_104_dt__update__tmp_h2).Dtor_NonEmptyBlob(), (_104_dt__update__tmp_h2).Dtor_BlobLessThanOrEqualToTen(), (_104_dt__update__tmp_h2).Dtor_MyList(), (_104_dt__update__tmp_h2).Dtor_NonEmptyList(), (_104_dt__update__tmp_h2).Dtor_ListLessThanOrEqualToTen(), (_104_dt__update__tmp_h2).Dtor_MyMap(), (_104_dt__update__tmp_h2).Dtor_NonEmptyMap(), (_104_dt__update__tmp_h2).Dtor_MapLessThanOrEqualToTen(), (_104_dt__update__tmp_h2).Dtor_OneToTen(), (_104_dt__update__tmp_h2).Dtor_myTenToTen(), (_104_dt__update__tmp_h2).Dtor_GreaterThanOne(), _105_dt__update_hLessThanTen_h2, (_104_dt__update__tmp_h2).Dtor_MyUtf8Bytes(), (_104_dt__update__tmp_h2).Dtor_MyListOfUtf8Bytes())
				}(_pat_let43_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceLessThanTen(_dafny.Zero)))
		}(_pat_let42_0)
	}(_98_input)
	var _out26 Wrappers.Result
	_ = _out26
	_out26 = (client).GetConstraints(_98_input)
	_101_ret = _out26
	if !((_101_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(159,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_98_input = func(_pat_let44_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_106_dt__update__tmp_h3 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let45_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_107_dt__update_hLessThanTen_h3 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_106_dt__update__tmp_h3).Dtor_MyString(), (_106_dt__update__tmp_h3).Dtor_NonEmptyString(), (_106_dt__update__tmp_h3).Dtor_StringLessThanOrEqualToTen(), (_106_dt__update__tmp_h3).Dtor_MyBlob(), (_106_dt__update__tmp_h3).Dtor_NonEmptyBlob(), (_106_dt__update__tmp_h3).Dtor_BlobLessThanOrEqualToTen(), (_106_dt__update__tmp_h3).Dtor_MyList(), (_106_dt__update__tmp_h3).Dtor_NonEmptyList(), (_106_dt__update__tmp_h3).Dtor_ListLessThanOrEqualToTen(), (_106_dt__update__tmp_h3).Dtor_MyMap(), (_106_dt__update__tmp_h3).Dtor_NonEmptyMap(), (_106_dt__update__tmp_h3).Dtor_MapLessThanOrEqualToTen(), (_106_dt__update__tmp_h3).Dtor_OneToTen(), (_106_dt__update__tmp_h3).Dtor_myTenToTen(), (_106_dt__update__tmp_h3).Dtor_GreaterThanOne(), _107_dt__update_hLessThanTen_h3, (_106_dt__update__tmp_h3).Dtor_MyUtf8Bytes(), (_106_dt__update__tmp_h3).Dtor_MyListOfUtf8Bytes())
				}(_pat_let45_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceLessThanTen(_dafny.One)))
		}(_pat_let44_0)
	}(_98_input)
	var _out27 Wrappers.Result
	_ = _out27
	_out27 = (client).GetConstraints(_98_input)
	_101_ret = _out27
	if !((_101_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(163,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_98_input = func(_pat_let46_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_108_dt__update__tmp_h4 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let47_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_109_dt__update_hLessThanTen_h4 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_108_dt__update__tmp_h4).Dtor_MyString(), (_108_dt__update__tmp_h4).Dtor_NonEmptyString(), (_108_dt__update__tmp_h4).Dtor_StringLessThanOrEqualToTen(), (_108_dt__update__tmp_h4).Dtor_MyBlob(), (_108_dt__update__tmp_h4).Dtor_NonEmptyBlob(), (_108_dt__update__tmp_h4).Dtor_BlobLessThanOrEqualToTen(), (_108_dt__update__tmp_h4).Dtor_MyList(), (_108_dt__update__tmp_h4).Dtor_NonEmptyList(), (_108_dt__update__tmp_h4).Dtor_ListLessThanOrEqualToTen(), (_108_dt__update__tmp_h4).Dtor_MyMap(), (_108_dt__update__tmp_h4).Dtor_NonEmptyMap(), (_108_dt__update__tmp_h4).Dtor_MapLessThanOrEqualToTen(), (_108_dt__update__tmp_h4).Dtor_OneToTen(), (_108_dt__update__tmp_h4).Dtor_myTenToTen(), (_108_dt__update__tmp_h4).Dtor_GreaterThanOne(), _109_dt__update_hLessThanTen_h4, (_108_dt__update__tmp_h4).Dtor_MyUtf8Bytes(), (_108_dt__update__tmp_h4).Dtor_MyListOfUtf8Bytes())
				}(_pat_let47_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceLessThanTen(_dafny.IntOfInt64(10))))
		}(_pat_let46_0)
	}(_98_input)
	var _out28 Wrappers.Result
	_ = _out28
	_out28 = (client).GetConstraints(_98_input)
	_101_ret = _out28
	if !((_101_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(167,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_98_input = func(_pat_let48_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_110_dt__update__tmp_h5 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let49_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_111_dt__update_hLessThanTen_h5 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_110_dt__update__tmp_h5).Dtor_MyString(), (_110_dt__update__tmp_h5).Dtor_NonEmptyString(), (_110_dt__update__tmp_h5).Dtor_StringLessThanOrEqualToTen(), (_110_dt__update__tmp_h5).Dtor_MyBlob(), (_110_dt__update__tmp_h5).Dtor_NonEmptyBlob(), (_110_dt__update__tmp_h5).Dtor_BlobLessThanOrEqualToTen(), (_110_dt__update__tmp_h5).Dtor_MyList(), (_110_dt__update__tmp_h5).Dtor_NonEmptyList(), (_110_dt__update__tmp_h5).Dtor_ListLessThanOrEqualToTen(), (_110_dt__update__tmp_h5).Dtor_MyMap(), (_110_dt__update__tmp_h5).Dtor_NonEmptyMap(), (_110_dt__update__tmp_h5).Dtor_MapLessThanOrEqualToTen(), (_110_dt__update__tmp_h5).Dtor_OneToTen(), (_110_dt__update__tmp_h5).Dtor_myTenToTen(), (_110_dt__update__tmp_h5).Dtor_GreaterThanOne(), _111_dt__update_hLessThanTen_h5, (_110_dt__update__tmp_h5).Dtor_MyUtf8Bytes(), (_110_dt__update__tmp_h5).Dtor_MyListOfUtf8Bytes())
				}(_pat_let49_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceLessThanTen(_dafny.IntOfInt64(11))))
		}(_pat_let48_0)
	}(_98_input)
	var _out29 Wrappers.Result
	_ = _out29
	_out29 = (client).GetConstraints(_98_input)
	_101_ret = _out29
	if !((_101_ret).Is_Failure()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(171,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
}
func (_static *CompanionStruct_Default___) TestGetConstraintWithNonEmptyString(client simpleconstraintsinternaldafnytypes.ISimpleConstraintsClient) {
	var _112_input simpleconstraintsinternaldafnytypes.GetConstraintsInput
	_ = _112_input
	_112_input = Helpers.Companion_Default___.GetValidInput()
	_112_input = func(_pat_let50_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_113_dt__update__tmp_h0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let51_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_114_dt__update_hNonEmptyString_h0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_113_dt__update__tmp_h0).Dtor_MyString(), _114_dt__update_hNonEmptyString_h0, (_113_dt__update__tmp_h0).Dtor_StringLessThanOrEqualToTen(), (_113_dt__update__tmp_h0).Dtor_MyBlob(), (_113_dt__update__tmp_h0).Dtor_NonEmptyBlob(), (_113_dt__update__tmp_h0).Dtor_BlobLessThanOrEqualToTen(), (_113_dt__update__tmp_h0).Dtor_MyList(), (_113_dt__update__tmp_h0).Dtor_NonEmptyList(), (_113_dt__update__tmp_h0).Dtor_ListLessThanOrEqualToTen(), (_113_dt__update__tmp_h0).Dtor_MyMap(), (_113_dt__update__tmp_h0).Dtor_NonEmptyMap(), (_113_dt__update__tmp_h0).Dtor_MapLessThanOrEqualToTen(), (_113_dt__update__tmp_h0).Dtor_OneToTen(), (_113_dt__update__tmp_h0).Dtor_myTenToTen(), (_113_dt__update__tmp_h0).Dtor_GreaterThanOne(), (_113_dt__update__tmp_h0).Dtor_LessThanTen(), (_113_dt__update__tmp_h0).Dtor_MyUtf8Bytes(), (_113_dt__update__tmp_h0).Dtor_MyListOfUtf8Bytes())
				}(_pat_let51_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceNonEmptyString(_dafny.SeqOfString("this string is really way too long"))))
		}(_pat_let50_0)
	}(_112_input)
	var _115_ret Wrappers.Result
	_ = _115_ret
	var _out30 Wrappers.Result
	_ = _out30
	_out30 = (client).GetConstraints(_112_input)
	_115_ret = _out30
	if !((_115_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(183,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_112_input = func(_pat_let52_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_116_dt__update__tmp_h1 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let53_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_117_dt__update_hNonEmptyString_h1 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_116_dt__update__tmp_h1).Dtor_MyString(), _117_dt__update_hNonEmptyString_h1, (_116_dt__update__tmp_h1).Dtor_StringLessThanOrEqualToTen(), (_116_dt__update__tmp_h1).Dtor_MyBlob(), (_116_dt__update__tmp_h1).Dtor_NonEmptyBlob(), (_116_dt__update__tmp_h1).Dtor_BlobLessThanOrEqualToTen(), (_116_dt__update__tmp_h1).Dtor_MyList(), (_116_dt__update__tmp_h1).Dtor_NonEmptyList(), (_116_dt__update__tmp_h1).Dtor_ListLessThanOrEqualToTen(), (_116_dt__update__tmp_h1).Dtor_MyMap(), (_116_dt__update__tmp_h1).Dtor_NonEmptyMap(), (_116_dt__update__tmp_h1).Dtor_MapLessThanOrEqualToTen(), (_116_dt__update__tmp_h1).Dtor_OneToTen(), (_116_dt__update__tmp_h1).Dtor_myTenToTen(), (_116_dt__update__tmp_h1).Dtor_GreaterThanOne(), (_116_dt__update__tmp_h1).Dtor_LessThanTen(), (_116_dt__update__tmp_h1).Dtor_MyUtf8Bytes(), (_116_dt__update__tmp_h1).Dtor_MyListOfUtf8Bytes())
				}(_pat_let53_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceNonEmptyString(_dafny.SeqOfString("12345678901"))))
		}(_pat_let52_0)
	}(_112_input)
	var _out31 Wrappers.Result
	_ = _out31
	_out31 = (client).GetConstraints(_112_input)
	_115_ret = _out31
	if !((_115_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(187,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_112_input = func(_pat_let54_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_118_dt__update__tmp_h2 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let55_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_119_dt__update_hNonEmptyString_h2 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_118_dt__update__tmp_h2).Dtor_MyString(), _119_dt__update_hNonEmptyString_h2, (_118_dt__update__tmp_h2).Dtor_StringLessThanOrEqualToTen(), (_118_dt__update__tmp_h2).Dtor_MyBlob(), (_118_dt__update__tmp_h2).Dtor_NonEmptyBlob(), (_118_dt__update__tmp_h2).Dtor_BlobLessThanOrEqualToTen(), (_118_dt__update__tmp_h2).Dtor_MyList(), (_118_dt__update__tmp_h2).Dtor_NonEmptyList(), (_118_dt__update__tmp_h2).Dtor_ListLessThanOrEqualToTen(), (_118_dt__update__tmp_h2).Dtor_MyMap(), (_118_dt__update__tmp_h2).Dtor_NonEmptyMap(), (_118_dt__update__tmp_h2).Dtor_MapLessThanOrEqualToTen(), (_118_dt__update__tmp_h2).Dtor_OneToTen(), (_118_dt__update__tmp_h2).Dtor_myTenToTen(), (_118_dt__update__tmp_h2).Dtor_GreaterThanOne(), (_118_dt__update__tmp_h2).Dtor_LessThanTen(), (_118_dt__update__tmp_h2).Dtor_MyUtf8Bytes(), (_118_dt__update__tmp_h2).Dtor_MyListOfUtf8Bytes())
				}(_pat_let55_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceNonEmptyString(_dafny.SeqOfString("1234567890"))))
		}(_pat_let54_0)
	}(_112_input)
	var _out32 Wrappers.Result
	_ = _out32
	_out32 = (client).GetConstraints(_112_input)
	_115_ret = _out32
	if !((_115_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(191,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_112_input = func(_pat_let56_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_120_dt__update__tmp_h3 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let57_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_121_dt__update_hNonEmptyString_h3 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_120_dt__update__tmp_h3).Dtor_MyString(), _121_dt__update_hNonEmptyString_h3, (_120_dt__update__tmp_h3).Dtor_StringLessThanOrEqualToTen(), (_120_dt__update__tmp_h3).Dtor_MyBlob(), (_120_dt__update__tmp_h3).Dtor_NonEmptyBlob(), (_120_dt__update__tmp_h3).Dtor_BlobLessThanOrEqualToTen(), (_120_dt__update__tmp_h3).Dtor_MyList(), (_120_dt__update__tmp_h3).Dtor_NonEmptyList(), (_120_dt__update__tmp_h3).Dtor_ListLessThanOrEqualToTen(), (_120_dt__update__tmp_h3).Dtor_MyMap(), (_120_dt__update__tmp_h3).Dtor_NonEmptyMap(), (_120_dt__update__tmp_h3).Dtor_MapLessThanOrEqualToTen(), (_120_dt__update__tmp_h3).Dtor_OneToTen(), (_120_dt__update__tmp_h3).Dtor_myTenToTen(), (_120_dt__update__tmp_h3).Dtor_GreaterThanOne(), (_120_dt__update__tmp_h3).Dtor_LessThanTen(), (_120_dt__update__tmp_h3).Dtor_MyUtf8Bytes(), (_120_dt__update__tmp_h3).Dtor_MyListOfUtf8Bytes())
				}(_pat_let57_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceNonEmptyString(_dafny.SeqOfString("1"))))
		}(_pat_let56_0)
	}(_112_input)
	var _out33 Wrappers.Result
	_ = _out33
	_out33 = (client).GetConstraints(_112_input)
	_115_ret = _out33
	if !((_115_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(195,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_112_input = func(_pat_let58_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_122_dt__update__tmp_h4 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let59_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_123_dt__update_hNonEmptyString_h4 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_122_dt__update__tmp_h4).Dtor_MyString(), _123_dt__update_hNonEmptyString_h4, (_122_dt__update__tmp_h4).Dtor_StringLessThanOrEqualToTen(), (_122_dt__update__tmp_h4).Dtor_MyBlob(), (_122_dt__update__tmp_h4).Dtor_NonEmptyBlob(), (_122_dt__update__tmp_h4).Dtor_BlobLessThanOrEqualToTen(), (_122_dt__update__tmp_h4).Dtor_MyList(), (_122_dt__update__tmp_h4).Dtor_NonEmptyList(), (_122_dt__update__tmp_h4).Dtor_ListLessThanOrEqualToTen(), (_122_dt__update__tmp_h4).Dtor_MyMap(), (_122_dt__update__tmp_h4).Dtor_NonEmptyMap(), (_122_dt__update__tmp_h4).Dtor_MapLessThanOrEqualToTen(), (_122_dt__update__tmp_h4).Dtor_OneToTen(), (_122_dt__update__tmp_h4).Dtor_myTenToTen(), (_122_dt__update__tmp_h4).Dtor_GreaterThanOne(), (_122_dt__update__tmp_h4).Dtor_LessThanTen(), (_122_dt__update__tmp_h4).Dtor_MyUtf8Bytes(), (_122_dt__update__tmp_h4).Dtor_MyListOfUtf8Bytes())
				}(_pat_let59_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceNonEmptyString(_dafny.SeqOfString(""))))
		}(_pat_let58_0)
	}(_112_input)
	var _out34 Wrappers.Result
	_ = _out34
	_out34 = (client).GetConstraints(_112_input)
	_115_ret = _out34
	if !((_115_ret).Is_Failure()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(199,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
}
func (_static *CompanionStruct_Default___) TestGetConstraintWithStringLessThanOrEqualToTen(client simpleconstraintsinternaldafnytypes.ISimpleConstraintsClient) {
	var _124_input simpleconstraintsinternaldafnytypes.GetConstraintsInput
	_ = _124_input
	_124_input = Helpers.Companion_Default___.GetValidInput()
	_124_input = func(_pat_let60_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_125_dt__update__tmp_h0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let61_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_126_dt__update_hStringLessThanOrEqualToTen_h0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_125_dt__update__tmp_h0).Dtor_MyString(), (_125_dt__update__tmp_h0).Dtor_NonEmptyString(), _126_dt__update_hStringLessThanOrEqualToTen_h0, (_125_dt__update__tmp_h0).Dtor_MyBlob(), (_125_dt__update__tmp_h0).Dtor_NonEmptyBlob(), (_125_dt__update__tmp_h0).Dtor_BlobLessThanOrEqualToTen(), (_125_dt__update__tmp_h0).Dtor_MyList(), (_125_dt__update__tmp_h0).Dtor_NonEmptyList(), (_125_dt__update__tmp_h0).Dtor_ListLessThanOrEqualToTen(), (_125_dt__update__tmp_h0).Dtor_MyMap(), (_125_dt__update__tmp_h0).Dtor_NonEmptyMap(), (_125_dt__update__tmp_h0).Dtor_MapLessThanOrEqualToTen(), (_125_dt__update__tmp_h0).Dtor_OneToTen(), (_125_dt__update__tmp_h0).Dtor_myTenToTen(), (_125_dt__update__tmp_h0).Dtor_GreaterThanOne(), (_125_dt__update__tmp_h0).Dtor_LessThanTen(), (_125_dt__update__tmp_h0).Dtor_MyUtf8Bytes(), (_125_dt__update__tmp_h0).Dtor_MyListOfUtf8Bytes())
				}(_pat_let61_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceStringLessThanOrEqualToTen(_dafny.SeqOfString("this string is really way too long"))))
		}(_pat_let60_0)
	}(_124_input)
	var _127_ret Wrappers.Result
	_ = _127_ret
	var _out35 Wrappers.Result
	_ = _out35
	_out35 = (client).GetConstraints(_124_input)
	_127_ret = _out35
	if !((_127_ret).Is_Failure()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(211,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_124_input = func(_pat_let62_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_128_dt__update__tmp_h1 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let63_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_129_dt__update_hStringLessThanOrEqualToTen_h1 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_128_dt__update__tmp_h1).Dtor_MyString(), (_128_dt__update__tmp_h1).Dtor_NonEmptyString(), _129_dt__update_hStringLessThanOrEqualToTen_h1, (_128_dt__update__tmp_h1).Dtor_MyBlob(), (_128_dt__update__tmp_h1).Dtor_NonEmptyBlob(), (_128_dt__update__tmp_h1).Dtor_BlobLessThanOrEqualToTen(), (_128_dt__update__tmp_h1).Dtor_MyList(), (_128_dt__update__tmp_h1).Dtor_NonEmptyList(), (_128_dt__update__tmp_h1).Dtor_ListLessThanOrEqualToTen(), (_128_dt__update__tmp_h1).Dtor_MyMap(), (_128_dt__update__tmp_h1).Dtor_NonEmptyMap(), (_128_dt__update__tmp_h1).Dtor_MapLessThanOrEqualToTen(), (_128_dt__update__tmp_h1).Dtor_OneToTen(), (_128_dt__update__tmp_h1).Dtor_myTenToTen(), (_128_dt__update__tmp_h1).Dtor_GreaterThanOne(), (_128_dt__update__tmp_h1).Dtor_LessThanTen(), (_128_dt__update__tmp_h1).Dtor_MyUtf8Bytes(), (_128_dt__update__tmp_h1).Dtor_MyListOfUtf8Bytes())
				}(_pat_let63_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceStringLessThanOrEqualToTen(_dafny.SeqOfString("12345678901"))))
		}(_pat_let62_0)
	}(_124_input)
	var _out36 Wrappers.Result
	_ = _out36
	_out36 = (client).GetConstraints(_124_input)
	_127_ret = _out36
	if !((_127_ret).Is_Failure()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(215,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_124_input = func(_pat_let64_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_130_dt__update__tmp_h2 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let65_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_131_dt__update_hStringLessThanOrEqualToTen_h2 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_130_dt__update__tmp_h2).Dtor_MyString(), (_130_dt__update__tmp_h2).Dtor_NonEmptyString(), _131_dt__update_hStringLessThanOrEqualToTen_h2, (_130_dt__update__tmp_h2).Dtor_MyBlob(), (_130_dt__update__tmp_h2).Dtor_NonEmptyBlob(), (_130_dt__update__tmp_h2).Dtor_BlobLessThanOrEqualToTen(), (_130_dt__update__tmp_h2).Dtor_MyList(), (_130_dt__update__tmp_h2).Dtor_NonEmptyList(), (_130_dt__update__tmp_h2).Dtor_ListLessThanOrEqualToTen(), (_130_dt__update__tmp_h2).Dtor_MyMap(), (_130_dt__update__tmp_h2).Dtor_NonEmptyMap(), (_130_dt__update__tmp_h2).Dtor_MapLessThanOrEqualToTen(), (_130_dt__update__tmp_h2).Dtor_OneToTen(), (_130_dt__update__tmp_h2).Dtor_myTenToTen(), (_130_dt__update__tmp_h2).Dtor_GreaterThanOne(), (_130_dt__update__tmp_h2).Dtor_LessThanTen(), (_130_dt__update__tmp_h2).Dtor_MyUtf8Bytes(), (_130_dt__update__tmp_h2).Dtor_MyListOfUtf8Bytes())
				}(_pat_let65_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceStringLessThanOrEqualToTen(_dafny.SeqOfString("1234567890"))))
		}(_pat_let64_0)
	}(_124_input)
	var _out37 Wrappers.Result
	_ = _out37
	_out37 = (client).GetConstraints(_124_input)
	_127_ret = _out37
	if !((_127_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(219,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_124_input = func(_pat_let66_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_132_dt__update__tmp_h3 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let67_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_133_dt__update_hStringLessThanOrEqualToTen_h3 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_132_dt__update__tmp_h3).Dtor_MyString(), (_132_dt__update__tmp_h3).Dtor_NonEmptyString(), _133_dt__update_hStringLessThanOrEqualToTen_h3, (_132_dt__update__tmp_h3).Dtor_MyBlob(), (_132_dt__update__tmp_h3).Dtor_NonEmptyBlob(), (_132_dt__update__tmp_h3).Dtor_BlobLessThanOrEqualToTen(), (_132_dt__update__tmp_h3).Dtor_MyList(), (_132_dt__update__tmp_h3).Dtor_NonEmptyList(), (_132_dt__update__tmp_h3).Dtor_ListLessThanOrEqualToTen(), (_132_dt__update__tmp_h3).Dtor_MyMap(), (_132_dt__update__tmp_h3).Dtor_NonEmptyMap(), (_132_dt__update__tmp_h3).Dtor_MapLessThanOrEqualToTen(), (_132_dt__update__tmp_h3).Dtor_OneToTen(), (_132_dt__update__tmp_h3).Dtor_myTenToTen(), (_132_dt__update__tmp_h3).Dtor_GreaterThanOne(), (_132_dt__update__tmp_h3).Dtor_LessThanTen(), (_132_dt__update__tmp_h3).Dtor_MyUtf8Bytes(), (_132_dt__update__tmp_h3).Dtor_MyListOfUtf8Bytes())
				}(_pat_let67_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceStringLessThanOrEqualToTen(_dafny.SeqOfString("1"))))
		}(_pat_let66_0)
	}(_124_input)
	var _out38 Wrappers.Result
	_ = _out38
	_out38 = (client).GetConstraints(_124_input)
	_127_ret = _out38
	if !((_127_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(223,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_124_input = func(_pat_let68_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_134_dt__update__tmp_h4 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let69_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_135_dt__update_hStringLessThanOrEqualToTen_h4 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_134_dt__update__tmp_h4).Dtor_MyString(), (_134_dt__update__tmp_h4).Dtor_NonEmptyString(), _135_dt__update_hStringLessThanOrEqualToTen_h4, (_134_dt__update__tmp_h4).Dtor_MyBlob(), (_134_dt__update__tmp_h4).Dtor_NonEmptyBlob(), (_134_dt__update__tmp_h4).Dtor_BlobLessThanOrEqualToTen(), (_134_dt__update__tmp_h4).Dtor_MyList(), (_134_dt__update__tmp_h4).Dtor_NonEmptyList(), (_134_dt__update__tmp_h4).Dtor_ListLessThanOrEqualToTen(), (_134_dt__update__tmp_h4).Dtor_MyMap(), (_134_dt__update__tmp_h4).Dtor_NonEmptyMap(), (_134_dt__update__tmp_h4).Dtor_MapLessThanOrEqualToTen(), (_134_dt__update__tmp_h4).Dtor_OneToTen(), (_134_dt__update__tmp_h4).Dtor_myTenToTen(), (_134_dt__update__tmp_h4).Dtor_GreaterThanOne(), (_134_dt__update__tmp_h4).Dtor_LessThanTen(), (_134_dt__update__tmp_h4).Dtor_MyUtf8Bytes(), (_134_dt__update__tmp_h4).Dtor_MyListOfUtf8Bytes())
				}(_pat_let69_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceStringLessThanOrEqualToTen(_dafny.SeqOfString(""))))
		}(_pat_let68_0)
	}(_124_input)
	var _out39 Wrappers.Result
	_ = _out39
	_out39 = (client).GetConstraints(_124_input)
	_127_ret = _out39
	if !((_127_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(227,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
}
func (_static *CompanionStruct_Default___) TestGetConstraintWithMyBlob(client simpleconstraintsinternaldafnytypes.ISimpleConstraintsClient) {
	var _136_input simpleconstraintsinternaldafnytypes.GetConstraintsInput
	_ = _136_input
	_136_input = Helpers.Companion_Default___.GetValidInput()
	_136_input = func(_pat_let70_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_137_dt__update__tmp_h0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let71_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_138_dt__update_hMyBlob_h0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_137_dt__update__tmp_h0).Dtor_MyString(), (_137_dt__update__tmp_h0).Dtor_NonEmptyString(), (_137_dt__update__tmp_h0).Dtor_StringLessThanOrEqualToTen(), _138_dt__update_hMyBlob_h0, (_137_dt__update__tmp_h0).Dtor_NonEmptyBlob(), (_137_dt__update__tmp_h0).Dtor_BlobLessThanOrEqualToTen(), (_137_dt__update__tmp_h0).Dtor_MyList(), (_137_dt__update__tmp_h0).Dtor_NonEmptyList(), (_137_dt__update__tmp_h0).Dtor_ListLessThanOrEqualToTen(), (_137_dt__update__tmp_h0).Dtor_MyMap(), (_137_dt__update__tmp_h0).Dtor_NonEmptyMap(), (_137_dt__update__tmp_h0).Dtor_MapLessThanOrEqualToTen(), (_137_dt__update__tmp_h0).Dtor_OneToTen(), (_137_dt__update__tmp_h0).Dtor_myTenToTen(), (_137_dt__update__tmp_h0).Dtor_GreaterThanOne(), (_137_dt__update__tmp_h0).Dtor_LessThanTen(), (_137_dt__update__tmp_h0).Dtor_MyUtf8Bytes(), (_137_dt__update__tmp_h0).Dtor_MyListOfUtf8Bytes())
				}(_pat_let71_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceMyBlob(_dafny.SeqOf())))
		}(_pat_let70_0)
	}(_136_input)
	var _139_ret Wrappers.Result
	_ = _139_ret
	var _out40 Wrappers.Result
	_ = _out40
	_out40 = (client).GetConstraints(_136_input)
	_139_ret = _out40
	if !((_139_ret).Is_Failure()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(239,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_136_input = func(_pat_let72_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_140_dt__update__tmp_h1 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let73_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_141_dt__update_hMyBlob_h1 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_140_dt__update__tmp_h1).Dtor_MyString(), (_140_dt__update__tmp_h1).Dtor_NonEmptyString(), (_140_dt__update__tmp_h1).Dtor_StringLessThanOrEqualToTen(), _141_dt__update_hMyBlob_h1, (_140_dt__update__tmp_h1).Dtor_NonEmptyBlob(), (_140_dt__update__tmp_h1).Dtor_BlobLessThanOrEqualToTen(), (_140_dt__update__tmp_h1).Dtor_MyList(), (_140_dt__update__tmp_h1).Dtor_NonEmptyList(), (_140_dt__update__tmp_h1).Dtor_ListLessThanOrEqualToTen(), (_140_dt__update__tmp_h1).Dtor_MyMap(), (_140_dt__update__tmp_h1).Dtor_NonEmptyMap(), (_140_dt__update__tmp_h1).Dtor_MapLessThanOrEqualToTen(), (_140_dt__update__tmp_h1).Dtor_OneToTen(), (_140_dt__update__tmp_h1).Dtor_myTenToTen(), (_140_dt__update__tmp_h1).Dtor_GreaterThanOne(), (_140_dt__update__tmp_h1).Dtor_LessThanTen(), (_140_dt__update__tmp_h1).Dtor_MyUtf8Bytes(), (_140_dt__update__tmp_h1).Dtor_MyListOfUtf8Bytes())
				}(_pat_let73_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceMyBlob(_dafny.SeqOf(uint8(1)))))
		}(_pat_let72_0)
	}(_136_input)
	var _out41 Wrappers.Result
	_ = _out41
	_out41 = (client).GetConstraints(_136_input)
	_139_ret = _out41
	if !((_139_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(243,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_136_input = func(_pat_let74_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_142_dt__update__tmp_h2 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let75_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_143_dt__update_hMyBlob_h2 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_142_dt__update__tmp_h2).Dtor_MyString(), (_142_dt__update__tmp_h2).Dtor_NonEmptyString(), (_142_dt__update__tmp_h2).Dtor_StringLessThanOrEqualToTen(), _143_dt__update_hMyBlob_h2, (_142_dt__update__tmp_h2).Dtor_NonEmptyBlob(), (_142_dt__update__tmp_h2).Dtor_BlobLessThanOrEqualToTen(), (_142_dt__update__tmp_h2).Dtor_MyList(), (_142_dt__update__tmp_h2).Dtor_NonEmptyList(), (_142_dt__update__tmp_h2).Dtor_ListLessThanOrEqualToTen(), (_142_dt__update__tmp_h2).Dtor_MyMap(), (_142_dt__update__tmp_h2).Dtor_NonEmptyMap(), (_142_dt__update__tmp_h2).Dtor_MapLessThanOrEqualToTen(), (_142_dt__update__tmp_h2).Dtor_OneToTen(), (_142_dt__update__tmp_h2).Dtor_myTenToTen(), (_142_dt__update__tmp_h2).Dtor_GreaterThanOne(), (_142_dt__update__tmp_h2).Dtor_LessThanTen(), (_142_dt__update__tmp_h2).Dtor_MyUtf8Bytes(), (_142_dt__update__tmp_h2).Dtor_MyListOfUtf8Bytes())
				}(_pat_let75_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceMyBlob(_dafny.SeqOf(uint8(1), uint8(2), uint8(3), uint8(4), uint8(5), uint8(6), uint8(7), uint8(8), uint8(9), uint8(0)))))
		}(_pat_let74_0)
	}(_136_input)
	var _out42 Wrappers.Result
	_ = _out42
	_out42 = (client).GetConstraints(_136_input)
	_139_ret = _out42
	if !((_139_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(247,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_136_input = func(_pat_let76_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_144_dt__update__tmp_h3 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let77_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_145_dt__update_hMyBlob_h3 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_144_dt__update__tmp_h3).Dtor_MyString(), (_144_dt__update__tmp_h3).Dtor_NonEmptyString(), (_144_dt__update__tmp_h3).Dtor_StringLessThanOrEqualToTen(), _145_dt__update_hMyBlob_h3, (_144_dt__update__tmp_h3).Dtor_NonEmptyBlob(), (_144_dt__update__tmp_h3).Dtor_BlobLessThanOrEqualToTen(), (_144_dt__update__tmp_h3).Dtor_MyList(), (_144_dt__update__tmp_h3).Dtor_NonEmptyList(), (_144_dt__update__tmp_h3).Dtor_ListLessThanOrEqualToTen(), (_144_dt__update__tmp_h3).Dtor_MyMap(), (_144_dt__update__tmp_h3).Dtor_NonEmptyMap(), (_144_dt__update__tmp_h3).Dtor_MapLessThanOrEqualToTen(), (_144_dt__update__tmp_h3).Dtor_OneToTen(), (_144_dt__update__tmp_h3).Dtor_myTenToTen(), (_144_dt__update__tmp_h3).Dtor_GreaterThanOne(), (_144_dt__update__tmp_h3).Dtor_LessThanTen(), (_144_dt__update__tmp_h3).Dtor_MyUtf8Bytes(), (_144_dt__update__tmp_h3).Dtor_MyListOfUtf8Bytes())
				}(_pat_let77_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceMyBlob(_dafny.SeqOf(uint8(1), uint8(2), uint8(3), uint8(4), uint8(5), uint8(6), uint8(7), uint8(8), uint8(9), uint8(0), uint8(1)))))
		}(_pat_let76_0)
	}(_136_input)
	var _out43 Wrappers.Result
	_ = _out43
	_out43 = (client).GetConstraints(_136_input)
	_139_ret = _out43
	if !((_139_ret).Is_Failure()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(251,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
}
func (_static *CompanionStruct_Default___) TestGetConstraintWithNonEmptyBlob(client simpleconstraintsinternaldafnytypes.ISimpleConstraintsClient) {
	var _146_input simpleconstraintsinternaldafnytypes.GetConstraintsInput
	_ = _146_input
	_146_input = Helpers.Companion_Default___.GetValidInput()
	_146_input = func(_pat_let78_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_147_dt__update__tmp_h0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let79_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_148_dt__update_hNonEmptyBlob_h0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_147_dt__update__tmp_h0).Dtor_MyString(), (_147_dt__update__tmp_h0).Dtor_NonEmptyString(), (_147_dt__update__tmp_h0).Dtor_StringLessThanOrEqualToTen(), (_147_dt__update__tmp_h0).Dtor_MyBlob(), _148_dt__update_hNonEmptyBlob_h0, (_147_dt__update__tmp_h0).Dtor_BlobLessThanOrEqualToTen(), (_147_dt__update__tmp_h0).Dtor_MyList(), (_147_dt__update__tmp_h0).Dtor_NonEmptyList(), (_147_dt__update__tmp_h0).Dtor_ListLessThanOrEqualToTen(), (_147_dt__update__tmp_h0).Dtor_MyMap(), (_147_dt__update__tmp_h0).Dtor_NonEmptyMap(), (_147_dt__update__tmp_h0).Dtor_MapLessThanOrEqualToTen(), (_147_dt__update__tmp_h0).Dtor_OneToTen(), (_147_dt__update__tmp_h0).Dtor_myTenToTen(), (_147_dt__update__tmp_h0).Dtor_GreaterThanOne(), (_147_dt__update__tmp_h0).Dtor_LessThanTen(), (_147_dt__update__tmp_h0).Dtor_MyUtf8Bytes(), (_147_dt__update__tmp_h0).Dtor_MyListOfUtf8Bytes())
				}(_pat_let79_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceNonEmptyBlob(_dafny.SeqOf())))
		}(_pat_let78_0)
	}(_146_input)
	var _149_ret Wrappers.Result
	_ = _149_ret
	var _out44 Wrappers.Result
	_ = _out44
	_out44 = (client).GetConstraints(_146_input)
	_149_ret = _out44
	if !((_149_ret).Is_Failure()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(263,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_146_input = func(_pat_let80_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_150_dt__update__tmp_h1 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let81_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_151_dt__update_hNonEmptyBlob_h1 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_150_dt__update__tmp_h1).Dtor_MyString(), (_150_dt__update__tmp_h1).Dtor_NonEmptyString(), (_150_dt__update__tmp_h1).Dtor_StringLessThanOrEqualToTen(), (_150_dt__update__tmp_h1).Dtor_MyBlob(), _151_dt__update_hNonEmptyBlob_h1, (_150_dt__update__tmp_h1).Dtor_BlobLessThanOrEqualToTen(), (_150_dt__update__tmp_h1).Dtor_MyList(), (_150_dt__update__tmp_h1).Dtor_NonEmptyList(), (_150_dt__update__tmp_h1).Dtor_ListLessThanOrEqualToTen(), (_150_dt__update__tmp_h1).Dtor_MyMap(), (_150_dt__update__tmp_h1).Dtor_NonEmptyMap(), (_150_dt__update__tmp_h1).Dtor_MapLessThanOrEqualToTen(), (_150_dt__update__tmp_h1).Dtor_OneToTen(), (_150_dt__update__tmp_h1).Dtor_myTenToTen(), (_150_dt__update__tmp_h1).Dtor_GreaterThanOne(), (_150_dt__update__tmp_h1).Dtor_LessThanTen(), (_150_dt__update__tmp_h1).Dtor_MyUtf8Bytes(), (_150_dt__update__tmp_h1).Dtor_MyListOfUtf8Bytes())
				}(_pat_let81_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceNonEmptyBlob(_dafny.SeqOf(uint8(1)))))
		}(_pat_let80_0)
	}(_146_input)
	var _out45 Wrappers.Result
	_ = _out45
	_out45 = (client).GetConstraints(_146_input)
	_149_ret = _out45
	if !((_149_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(267,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_146_input = func(_pat_let82_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_152_dt__update__tmp_h2 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let83_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_153_dt__update_hNonEmptyBlob_h2 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_152_dt__update__tmp_h2).Dtor_MyString(), (_152_dt__update__tmp_h2).Dtor_NonEmptyString(), (_152_dt__update__tmp_h2).Dtor_StringLessThanOrEqualToTen(), (_152_dt__update__tmp_h2).Dtor_MyBlob(), _153_dt__update_hNonEmptyBlob_h2, (_152_dt__update__tmp_h2).Dtor_BlobLessThanOrEqualToTen(), (_152_dt__update__tmp_h2).Dtor_MyList(), (_152_dt__update__tmp_h2).Dtor_NonEmptyList(), (_152_dt__update__tmp_h2).Dtor_ListLessThanOrEqualToTen(), (_152_dt__update__tmp_h2).Dtor_MyMap(), (_152_dt__update__tmp_h2).Dtor_NonEmptyMap(), (_152_dt__update__tmp_h2).Dtor_MapLessThanOrEqualToTen(), (_152_dt__update__tmp_h2).Dtor_OneToTen(), (_152_dt__update__tmp_h2).Dtor_myTenToTen(), (_152_dt__update__tmp_h2).Dtor_GreaterThanOne(), (_152_dt__update__tmp_h2).Dtor_LessThanTen(), (_152_dt__update__tmp_h2).Dtor_MyUtf8Bytes(), (_152_dt__update__tmp_h2).Dtor_MyListOfUtf8Bytes())
				}(_pat_let83_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceNonEmptyBlob(_dafny.SeqOf(uint8(1), uint8(2), uint8(3), uint8(4), uint8(5), uint8(6), uint8(7), uint8(8), uint8(9), uint8(0)))))
		}(_pat_let82_0)
	}(_146_input)
	var _out46 Wrappers.Result
	_ = _out46
	_out46 = (client).GetConstraints(_146_input)
	_149_ret = _out46
	if !((_149_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(271,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
}
func (_static *CompanionStruct_Default___) TestGetConstraintWithBlobLessThanOrEqualToTen(client simpleconstraintsinternaldafnytypes.ISimpleConstraintsClient) {
	var _154_input simpleconstraintsinternaldafnytypes.GetConstraintsInput
	_ = _154_input
	_154_input = Helpers.Companion_Default___.GetValidInput()
	_154_input = func(_pat_let84_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_155_dt__update__tmp_h0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let85_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_156_dt__update_hBlobLessThanOrEqualToTen_h0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_155_dt__update__tmp_h0).Dtor_MyString(), (_155_dt__update__tmp_h0).Dtor_NonEmptyString(), (_155_dt__update__tmp_h0).Dtor_StringLessThanOrEqualToTen(), (_155_dt__update__tmp_h0).Dtor_MyBlob(), (_155_dt__update__tmp_h0).Dtor_NonEmptyBlob(), _156_dt__update_hBlobLessThanOrEqualToTen_h0, (_155_dt__update__tmp_h0).Dtor_MyList(), (_155_dt__update__tmp_h0).Dtor_NonEmptyList(), (_155_dt__update__tmp_h0).Dtor_ListLessThanOrEqualToTen(), (_155_dt__update__tmp_h0).Dtor_MyMap(), (_155_dt__update__tmp_h0).Dtor_NonEmptyMap(), (_155_dt__update__tmp_h0).Dtor_MapLessThanOrEqualToTen(), (_155_dt__update__tmp_h0).Dtor_OneToTen(), (_155_dt__update__tmp_h0).Dtor_myTenToTen(), (_155_dt__update__tmp_h0).Dtor_GreaterThanOne(), (_155_dt__update__tmp_h0).Dtor_LessThanTen(), (_155_dt__update__tmp_h0).Dtor_MyUtf8Bytes(), (_155_dt__update__tmp_h0).Dtor_MyListOfUtf8Bytes())
				}(_pat_let85_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceBlobLessThanOrEqualToTen(_dafny.SeqOf())))
		}(_pat_let84_0)
	}(_154_input)
	var _157_ret Wrappers.Result
	_ = _157_ret
	var _out47 Wrappers.Result
	_ = _out47
	_out47 = (client).GetConstraints(_154_input)
	_157_ret = _out47
	if !((_157_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(283,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_154_input = func(_pat_let86_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_158_dt__update__tmp_h1 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let87_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_159_dt__update_hBlobLessThanOrEqualToTen_h1 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_158_dt__update__tmp_h1).Dtor_MyString(), (_158_dt__update__tmp_h1).Dtor_NonEmptyString(), (_158_dt__update__tmp_h1).Dtor_StringLessThanOrEqualToTen(), (_158_dt__update__tmp_h1).Dtor_MyBlob(), (_158_dt__update__tmp_h1).Dtor_NonEmptyBlob(), _159_dt__update_hBlobLessThanOrEqualToTen_h1, (_158_dt__update__tmp_h1).Dtor_MyList(), (_158_dt__update__tmp_h1).Dtor_NonEmptyList(), (_158_dt__update__tmp_h1).Dtor_ListLessThanOrEqualToTen(), (_158_dt__update__tmp_h1).Dtor_MyMap(), (_158_dt__update__tmp_h1).Dtor_NonEmptyMap(), (_158_dt__update__tmp_h1).Dtor_MapLessThanOrEqualToTen(), (_158_dt__update__tmp_h1).Dtor_OneToTen(), (_158_dt__update__tmp_h1).Dtor_myTenToTen(), (_158_dt__update__tmp_h1).Dtor_GreaterThanOne(), (_158_dt__update__tmp_h1).Dtor_LessThanTen(), (_158_dt__update__tmp_h1).Dtor_MyUtf8Bytes(), (_158_dt__update__tmp_h1).Dtor_MyListOfUtf8Bytes())
				}(_pat_let87_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceBlobLessThanOrEqualToTen(_dafny.SeqOf(uint8(1)))))
		}(_pat_let86_0)
	}(_154_input)
	var _out48 Wrappers.Result
	_ = _out48
	_out48 = (client).GetConstraints(_154_input)
	_157_ret = _out48
	if !((_157_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(287,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_154_input = func(_pat_let88_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_160_dt__update__tmp_h2 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let89_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_161_dt__update_hBlobLessThanOrEqualToTen_h2 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_160_dt__update__tmp_h2).Dtor_MyString(), (_160_dt__update__tmp_h2).Dtor_NonEmptyString(), (_160_dt__update__tmp_h2).Dtor_StringLessThanOrEqualToTen(), (_160_dt__update__tmp_h2).Dtor_MyBlob(), (_160_dt__update__tmp_h2).Dtor_NonEmptyBlob(), _161_dt__update_hBlobLessThanOrEqualToTen_h2, (_160_dt__update__tmp_h2).Dtor_MyList(), (_160_dt__update__tmp_h2).Dtor_NonEmptyList(), (_160_dt__update__tmp_h2).Dtor_ListLessThanOrEqualToTen(), (_160_dt__update__tmp_h2).Dtor_MyMap(), (_160_dt__update__tmp_h2).Dtor_NonEmptyMap(), (_160_dt__update__tmp_h2).Dtor_MapLessThanOrEqualToTen(), (_160_dt__update__tmp_h2).Dtor_OneToTen(), (_160_dt__update__tmp_h2).Dtor_myTenToTen(), (_160_dt__update__tmp_h2).Dtor_GreaterThanOne(), (_160_dt__update__tmp_h2).Dtor_LessThanTen(), (_160_dt__update__tmp_h2).Dtor_MyUtf8Bytes(), (_160_dt__update__tmp_h2).Dtor_MyListOfUtf8Bytes())
				}(_pat_let89_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceBlobLessThanOrEqualToTen(_dafny.SeqOf(uint8(1), uint8(2), uint8(3), uint8(4), uint8(5), uint8(6), uint8(7), uint8(8), uint8(9), uint8(0)))))
		}(_pat_let88_0)
	}(_154_input)
	var _out49 Wrappers.Result
	_ = _out49
	_out49 = (client).GetConstraints(_154_input)
	_157_ret = _out49
	if !((_157_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(291,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_154_input = func(_pat_let90_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_162_dt__update__tmp_h3 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let91_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_163_dt__update_hBlobLessThanOrEqualToTen_h3 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_162_dt__update__tmp_h3).Dtor_MyString(), (_162_dt__update__tmp_h3).Dtor_NonEmptyString(), (_162_dt__update__tmp_h3).Dtor_StringLessThanOrEqualToTen(), (_162_dt__update__tmp_h3).Dtor_MyBlob(), (_162_dt__update__tmp_h3).Dtor_NonEmptyBlob(), _163_dt__update_hBlobLessThanOrEqualToTen_h3, (_162_dt__update__tmp_h3).Dtor_MyList(), (_162_dt__update__tmp_h3).Dtor_NonEmptyList(), (_162_dt__update__tmp_h3).Dtor_ListLessThanOrEqualToTen(), (_162_dt__update__tmp_h3).Dtor_MyMap(), (_162_dt__update__tmp_h3).Dtor_NonEmptyMap(), (_162_dt__update__tmp_h3).Dtor_MapLessThanOrEqualToTen(), (_162_dt__update__tmp_h3).Dtor_OneToTen(), (_162_dt__update__tmp_h3).Dtor_myTenToTen(), (_162_dt__update__tmp_h3).Dtor_GreaterThanOne(), (_162_dt__update__tmp_h3).Dtor_LessThanTen(), (_162_dt__update__tmp_h3).Dtor_MyUtf8Bytes(), (_162_dt__update__tmp_h3).Dtor_MyListOfUtf8Bytes())
				}(_pat_let91_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceBlobLessThanOrEqualToTen(_dafny.SeqOf(uint8(1), uint8(2), uint8(3), uint8(4), uint8(5), uint8(6), uint8(7), uint8(8), uint8(9), uint8(0), uint8(1)))))
		}(_pat_let90_0)
	}(_154_input)
	var _out50 Wrappers.Result
	_ = _out50
	_out50 = (client).GetConstraints(_154_input)
	_157_ret = _out50
	if !((_157_ret).Is_Failure()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(295,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
}
func (_static *CompanionStruct_Default___) TestGetConstraintWithMyList(client simpleconstraintsinternaldafnytypes.ISimpleConstraintsClient) {
	var _164_input simpleconstraintsinternaldafnytypes.GetConstraintsInput
	_ = _164_input
	_164_input = Helpers.Companion_Default___.GetValidInput()
	_164_input = func(_pat_let92_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_165_dt__update__tmp_h0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let93_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_166_dt__update_hMyList_h0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_165_dt__update__tmp_h0).Dtor_MyString(), (_165_dt__update__tmp_h0).Dtor_NonEmptyString(), (_165_dt__update__tmp_h0).Dtor_StringLessThanOrEqualToTen(), (_165_dt__update__tmp_h0).Dtor_MyBlob(), (_165_dt__update__tmp_h0).Dtor_NonEmptyBlob(), (_165_dt__update__tmp_h0).Dtor_BlobLessThanOrEqualToTen(), _166_dt__update_hMyList_h0, (_165_dt__update__tmp_h0).Dtor_NonEmptyList(), (_165_dt__update__tmp_h0).Dtor_ListLessThanOrEqualToTen(), (_165_dt__update__tmp_h0).Dtor_MyMap(), (_165_dt__update__tmp_h0).Dtor_NonEmptyMap(), (_165_dt__update__tmp_h0).Dtor_MapLessThanOrEqualToTen(), (_165_dt__update__tmp_h0).Dtor_OneToTen(), (_165_dt__update__tmp_h0).Dtor_myTenToTen(), (_165_dt__update__tmp_h0).Dtor_GreaterThanOne(), (_165_dt__update__tmp_h0).Dtor_LessThanTen(), (_165_dt__update__tmp_h0).Dtor_MyUtf8Bytes(), (_165_dt__update__tmp_h0).Dtor_MyListOfUtf8Bytes())
				}(_pat_let93_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceMyList(_dafny.SeqOf())))
		}(_pat_let92_0)
	}(_164_input)
	var _167_ret Wrappers.Result
	_ = _167_ret
	var _out51 Wrappers.Result
	_ = _out51
	_out51 = (client).GetConstraints(_164_input)
	_167_ret = _out51
	if !((_167_ret).Is_Failure()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(307,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_164_input = func(_pat_let94_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_168_dt__update__tmp_h1 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let95_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_169_dt__update_hMyList_h1 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_168_dt__update__tmp_h1).Dtor_MyString(), (_168_dt__update__tmp_h1).Dtor_NonEmptyString(), (_168_dt__update__tmp_h1).Dtor_StringLessThanOrEqualToTen(), (_168_dt__update__tmp_h1).Dtor_MyBlob(), (_168_dt__update__tmp_h1).Dtor_NonEmptyBlob(), (_168_dt__update__tmp_h1).Dtor_BlobLessThanOrEqualToTen(), _169_dt__update_hMyList_h1, (_168_dt__update__tmp_h1).Dtor_NonEmptyList(), (_168_dt__update__tmp_h1).Dtor_ListLessThanOrEqualToTen(), (_168_dt__update__tmp_h1).Dtor_MyMap(), (_168_dt__update__tmp_h1).Dtor_NonEmptyMap(), (_168_dt__update__tmp_h1).Dtor_MapLessThanOrEqualToTen(), (_168_dt__update__tmp_h1).Dtor_OneToTen(), (_168_dt__update__tmp_h1).Dtor_myTenToTen(), (_168_dt__update__tmp_h1).Dtor_GreaterThanOne(), (_168_dt__update__tmp_h1).Dtor_LessThanTen(), (_168_dt__update__tmp_h1).Dtor_MyUtf8Bytes(), (_168_dt__update__tmp_h1).Dtor_MyListOfUtf8Bytes())
				}(_pat_let95_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceMyList(_dafny.SeqOf(_dafny.SeqOfString("1")))))
		}(_pat_let94_0)
	}(_164_input)
	var _out52 Wrappers.Result
	_ = _out52
	_out52 = (client).GetConstraints(_164_input)
	_167_ret = _out52
	if !((_167_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(311,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_164_input = func(_pat_let96_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_170_dt__update__tmp_h2 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let97_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_171_dt__update_hMyList_h2 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_170_dt__update__tmp_h2).Dtor_MyString(), (_170_dt__update__tmp_h2).Dtor_NonEmptyString(), (_170_dt__update__tmp_h2).Dtor_StringLessThanOrEqualToTen(), (_170_dt__update__tmp_h2).Dtor_MyBlob(), (_170_dt__update__tmp_h2).Dtor_NonEmptyBlob(), (_170_dt__update__tmp_h2).Dtor_BlobLessThanOrEqualToTen(), _171_dt__update_hMyList_h2, (_170_dt__update__tmp_h2).Dtor_NonEmptyList(), (_170_dt__update__tmp_h2).Dtor_ListLessThanOrEqualToTen(), (_170_dt__update__tmp_h2).Dtor_MyMap(), (_170_dt__update__tmp_h2).Dtor_NonEmptyMap(), (_170_dt__update__tmp_h2).Dtor_MapLessThanOrEqualToTen(), (_170_dt__update__tmp_h2).Dtor_OneToTen(), (_170_dt__update__tmp_h2).Dtor_myTenToTen(), (_170_dt__update__tmp_h2).Dtor_GreaterThanOne(), (_170_dt__update__tmp_h2).Dtor_LessThanTen(), (_170_dt__update__tmp_h2).Dtor_MyUtf8Bytes(), (_170_dt__update__tmp_h2).Dtor_MyListOfUtf8Bytes())
				}(_pat_let97_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceMyList(_dafny.SeqOf(_dafny.SeqOfString("1"), _dafny.SeqOfString("2"), _dafny.SeqOfString("3"), _dafny.SeqOfString("4"), _dafny.SeqOfString("5"), _dafny.SeqOfString("6"), _dafny.SeqOfString("7"), _dafny.SeqOfString("8"), _dafny.SeqOfString("9"), _dafny.SeqOfString("0")))))
		}(_pat_let96_0)
	}(_164_input)
	var _out53 Wrappers.Result
	_ = _out53
	_out53 = (client).GetConstraints(_164_input)
	_167_ret = _out53
	if !((_167_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(315,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_164_input = func(_pat_let98_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_172_dt__update__tmp_h3 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let99_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_173_dt__update_hMyList_h3 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_172_dt__update__tmp_h3).Dtor_MyString(), (_172_dt__update__tmp_h3).Dtor_NonEmptyString(), (_172_dt__update__tmp_h3).Dtor_StringLessThanOrEqualToTen(), (_172_dt__update__tmp_h3).Dtor_MyBlob(), (_172_dt__update__tmp_h3).Dtor_NonEmptyBlob(), (_172_dt__update__tmp_h3).Dtor_BlobLessThanOrEqualToTen(), _173_dt__update_hMyList_h3, (_172_dt__update__tmp_h3).Dtor_NonEmptyList(), (_172_dt__update__tmp_h3).Dtor_ListLessThanOrEqualToTen(), (_172_dt__update__tmp_h3).Dtor_MyMap(), (_172_dt__update__tmp_h3).Dtor_NonEmptyMap(), (_172_dt__update__tmp_h3).Dtor_MapLessThanOrEqualToTen(), (_172_dt__update__tmp_h3).Dtor_OneToTen(), (_172_dt__update__tmp_h3).Dtor_myTenToTen(), (_172_dt__update__tmp_h3).Dtor_GreaterThanOne(), (_172_dt__update__tmp_h3).Dtor_LessThanTen(), (_172_dt__update__tmp_h3).Dtor_MyUtf8Bytes(), (_172_dt__update__tmp_h3).Dtor_MyListOfUtf8Bytes())
				}(_pat_let99_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceMyList(_dafny.SeqOf(_dafny.SeqOfString("1"), _dafny.SeqOfString("2"), _dafny.SeqOfString("3"), _dafny.SeqOfString("4"), _dafny.SeqOfString("5"), _dafny.SeqOfString("6"), _dafny.SeqOfString("7"), _dafny.SeqOfString("8"), _dafny.SeqOfString("9"), _dafny.SeqOfString("0"), _dafny.SeqOfString("a")))))
		}(_pat_let98_0)
	}(_164_input)
	var _out54 Wrappers.Result
	_ = _out54
	_out54 = (client).GetConstraints(_164_input)
	_167_ret = _out54
	if !((_167_ret).Is_Failure()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(319,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
}
func (_static *CompanionStruct_Default___) TestGetConstraintWithNonEmptyList(client simpleconstraintsinternaldafnytypes.ISimpleConstraintsClient) {
	var _174_input simpleconstraintsinternaldafnytypes.GetConstraintsInput
	_ = _174_input
	_174_input = Helpers.Companion_Default___.GetValidInput()
	_174_input = func(_pat_let100_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_175_dt__update__tmp_h0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let101_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_176_dt__update_hNonEmptyList_h0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_175_dt__update__tmp_h0).Dtor_MyString(), (_175_dt__update__tmp_h0).Dtor_NonEmptyString(), (_175_dt__update__tmp_h0).Dtor_StringLessThanOrEqualToTen(), (_175_dt__update__tmp_h0).Dtor_MyBlob(), (_175_dt__update__tmp_h0).Dtor_NonEmptyBlob(), (_175_dt__update__tmp_h0).Dtor_BlobLessThanOrEqualToTen(), (_175_dt__update__tmp_h0).Dtor_MyList(), _176_dt__update_hNonEmptyList_h0, (_175_dt__update__tmp_h0).Dtor_ListLessThanOrEqualToTen(), (_175_dt__update__tmp_h0).Dtor_MyMap(), (_175_dt__update__tmp_h0).Dtor_NonEmptyMap(), (_175_dt__update__tmp_h0).Dtor_MapLessThanOrEqualToTen(), (_175_dt__update__tmp_h0).Dtor_OneToTen(), (_175_dt__update__tmp_h0).Dtor_myTenToTen(), (_175_dt__update__tmp_h0).Dtor_GreaterThanOne(), (_175_dt__update__tmp_h0).Dtor_LessThanTen(), (_175_dt__update__tmp_h0).Dtor_MyUtf8Bytes(), (_175_dt__update__tmp_h0).Dtor_MyListOfUtf8Bytes())
				}(_pat_let101_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceNonEmptyList(_dafny.SeqOf())))
		}(_pat_let100_0)
	}(_174_input)
	var _177_ret Wrappers.Result
	_ = _177_ret
	var _out55 Wrappers.Result
	_ = _out55
	_out55 = (client).GetConstraints(_174_input)
	_177_ret = _out55
	if !((_177_ret).Is_Failure()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(331,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_174_input = func(_pat_let102_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_178_dt__update__tmp_h1 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let103_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_179_dt__update_hNonEmptyList_h1 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_178_dt__update__tmp_h1).Dtor_MyString(), (_178_dt__update__tmp_h1).Dtor_NonEmptyString(), (_178_dt__update__tmp_h1).Dtor_StringLessThanOrEqualToTen(), (_178_dt__update__tmp_h1).Dtor_MyBlob(), (_178_dt__update__tmp_h1).Dtor_NonEmptyBlob(), (_178_dt__update__tmp_h1).Dtor_BlobLessThanOrEqualToTen(), (_178_dt__update__tmp_h1).Dtor_MyList(), _179_dt__update_hNonEmptyList_h1, (_178_dt__update__tmp_h1).Dtor_ListLessThanOrEqualToTen(), (_178_dt__update__tmp_h1).Dtor_MyMap(), (_178_dt__update__tmp_h1).Dtor_NonEmptyMap(), (_178_dt__update__tmp_h1).Dtor_MapLessThanOrEqualToTen(), (_178_dt__update__tmp_h1).Dtor_OneToTen(), (_178_dt__update__tmp_h1).Dtor_myTenToTen(), (_178_dt__update__tmp_h1).Dtor_GreaterThanOne(), (_178_dt__update__tmp_h1).Dtor_LessThanTen(), (_178_dt__update__tmp_h1).Dtor_MyUtf8Bytes(), (_178_dt__update__tmp_h1).Dtor_MyListOfUtf8Bytes())
				}(_pat_let103_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceNonEmptyList(_dafny.SeqOf(_dafny.SeqOfString("1")))))
		}(_pat_let102_0)
	}(_174_input)
	var _out56 Wrappers.Result
	_ = _out56
	_out56 = (client).GetConstraints(_174_input)
	_177_ret = _out56
	if !((_177_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(335,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_174_input = func(_pat_let104_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_180_dt__update__tmp_h2 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let105_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_181_dt__update_hNonEmptyList_h2 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_180_dt__update__tmp_h2).Dtor_MyString(), (_180_dt__update__tmp_h2).Dtor_NonEmptyString(), (_180_dt__update__tmp_h2).Dtor_StringLessThanOrEqualToTen(), (_180_dt__update__tmp_h2).Dtor_MyBlob(), (_180_dt__update__tmp_h2).Dtor_NonEmptyBlob(), (_180_dt__update__tmp_h2).Dtor_BlobLessThanOrEqualToTen(), (_180_dt__update__tmp_h2).Dtor_MyList(), _181_dt__update_hNonEmptyList_h2, (_180_dt__update__tmp_h2).Dtor_ListLessThanOrEqualToTen(), (_180_dt__update__tmp_h2).Dtor_MyMap(), (_180_dt__update__tmp_h2).Dtor_NonEmptyMap(), (_180_dt__update__tmp_h2).Dtor_MapLessThanOrEqualToTen(), (_180_dt__update__tmp_h2).Dtor_OneToTen(), (_180_dt__update__tmp_h2).Dtor_myTenToTen(), (_180_dt__update__tmp_h2).Dtor_GreaterThanOne(), (_180_dt__update__tmp_h2).Dtor_LessThanTen(), (_180_dt__update__tmp_h2).Dtor_MyUtf8Bytes(), (_180_dt__update__tmp_h2).Dtor_MyListOfUtf8Bytes())
				}(_pat_let105_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceNonEmptyList(_dafny.SeqOf(_dafny.SeqOfString("1"), _dafny.SeqOfString("2"), _dafny.SeqOfString("3"), _dafny.SeqOfString("4"), _dafny.SeqOfString("5"), _dafny.SeqOfString("6"), _dafny.SeqOfString("7"), _dafny.SeqOfString("8"), _dafny.SeqOfString("9"), _dafny.SeqOfString("0")))))
		}(_pat_let104_0)
	}(_174_input)
	var _out57 Wrappers.Result
	_ = _out57
	_out57 = (client).GetConstraints(_174_input)
	_177_ret = _out57
	if !((_177_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(339,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_174_input = func(_pat_let106_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_182_dt__update__tmp_h3 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let107_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_183_dt__update_hNonEmptyList_h3 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_182_dt__update__tmp_h3).Dtor_MyString(), (_182_dt__update__tmp_h3).Dtor_NonEmptyString(), (_182_dt__update__tmp_h3).Dtor_StringLessThanOrEqualToTen(), (_182_dt__update__tmp_h3).Dtor_MyBlob(), (_182_dt__update__tmp_h3).Dtor_NonEmptyBlob(), (_182_dt__update__tmp_h3).Dtor_BlobLessThanOrEqualToTen(), (_182_dt__update__tmp_h3).Dtor_MyList(), _183_dt__update_hNonEmptyList_h3, (_182_dt__update__tmp_h3).Dtor_ListLessThanOrEqualToTen(), (_182_dt__update__tmp_h3).Dtor_MyMap(), (_182_dt__update__tmp_h3).Dtor_NonEmptyMap(), (_182_dt__update__tmp_h3).Dtor_MapLessThanOrEqualToTen(), (_182_dt__update__tmp_h3).Dtor_OneToTen(), (_182_dt__update__tmp_h3).Dtor_myTenToTen(), (_182_dt__update__tmp_h3).Dtor_GreaterThanOne(), (_182_dt__update__tmp_h3).Dtor_LessThanTen(), (_182_dt__update__tmp_h3).Dtor_MyUtf8Bytes(), (_182_dt__update__tmp_h3).Dtor_MyListOfUtf8Bytes())
				}(_pat_let107_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceNonEmptyList(_dafny.SeqOf(_dafny.SeqOfString("1"), _dafny.SeqOfString("2"), _dafny.SeqOfString("3"), _dafny.SeqOfString("4"), _dafny.SeqOfString("5"), _dafny.SeqOfString("6"), _dafny.SeqOfString("7"), _dafny.SeqOfString("8"), _dafny.SeqOfString("9"), _dafny.SeqOfString("0"), _dafny.SeqOfString("a")))))
		}(_pat_let106_0)
	}(_174_input)
	var _out58 Wrappers.Result
	_ = _out58
	_out58 = (client).GetConstraints(_174_input)
	_177_ret = _out58
	if !((_177_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(343,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
}
func (_static *CompanionStruct_Default___) TestGetConstraintWithListLessThanOrEqualToTen(client simpleconstraintsinternaldafnytypes.ISimpleConstraintsClient) {
	var _184_input simpleconstraintsinternaldafnytypes.GetConstraintsInput
	_ = _184_input
	_184_input = Helpers.Companion_Default___.GetValidInput()
	_184_input = func(_pat_let108_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_185_dt__update__tmp_h0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let109_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_186_dt__update_hListLessThanOrEqualToTen_h0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_185_dt__update__tmp_h0).Dtor_MyString(), (_185_dt__update__tmp_h0).Dtor_NonEmptyString(), (_185_dt__update__tmp_h0).Dtor_StringLessThanOrEqualToTen(), (_185_dt__update__tmp_h0).Dtor_MyBlob(), (_185_dt__update__tmp_h0).Dtor_NonEmptyBlob(), (_185_dt__update__tmp_h0).Dtor_BlobLessThanOrEqualToTen(), (_185_dt__update__tmp_h0).Dtor_MyList(), (_185_dt__update__tmp_h0).Dtor_NonEmptyList(), _186_dt__update_hListLessThanOrEqualToTen_h0, (_185_dt__update__tmp_h0).Dtor_MyMap(), (_185_dt__update__tmp_h0).Dtor_NonEmptyMap(), (_185_dt__update__tmp_h0).Dtor_MapLessThanOrEqualToTen(), (_185_dt__update__tmp_h0).Dtor_OneToTen(), (_185_dt__update__tmp_h0).Dtor_myTenToTen(), (_185_dt__update__tmp_h0).Dtor_GreaterThanOne(), (_185_dt__update__tmp_h0).Dtor_LessThanTen(), (_185_dt__update__tmp_h0).Dtor_MyUtf8Bytes(), (_185_dt__update__tmp_h0).Dtor_MyListOfUtf8Bytes())
				}(_pat_let109_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceListLessThanOrEqualToTen(_dafny.SeqOf())))
		}(_pat_let108_0)
	}(_184_input)
	var _187_ret Wrappers.Result
	_ = _187_ret
	var _out59 Wrappers.Result
	_ = _out59
	_out59 = (client).GetConstraints(_184_input)
	_187_ret = _out59
	if !((_187_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(355,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_184_input = func(_pat_let110_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_188_dt__update__tmp_h1 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let111_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_189_dt__update_hListLessThanOrEqualToTen_h1 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_188_dt__update__tmp_h1).Dtor_MyString(), (_188_dt__update__tmp_h1).Dtor_NonEmptyString(), (_188_dt__update__tmp_h1).Dtor_StringLessThanOrEqualToTen(), (_188_dt__update__tmp_h1).Dtor_MyBlob(), (_188_dt__update__tmp_h1).Dtor_NonEmptyBlob(), (_188_dt__update__tmp_h1).Dtor_BlobLessThanOrEqualToTen(), (_188_dt__update__tmp_h1).Dtor_MyList(), (_188_dt__update__tmp_h1).Dtor_NonEmptyList(), _189_dt__update_hListLessThanOrEqualToTen_h1, (_188_dt__update__tmp_h1).Dtor_MyMap(), (_188_dt__update__tmp_h1).Dtor_NonEmptyMap(), (_188_dt__update__tmp_h1).Dtor_MapLessThanOrEqualToTen(), (_188_dt__update__tmp_h1).Dtor_OneToTen(), (_188_dt__update__tmp_h1).Dtor_myTenToTen(), (_188_dt__update__tmp_h1).Dtor_GreaterThanOne(), (_188_dt__update__tmp_h1).Dtor_LessThanTen(), (_188_dt__update__tmp_h1).Dtor_MyUtf8Bytes(), (_188_dt__update__tmp_h1).Dtor_MyListOfUtf8Bytes())
				}(_pat_let111_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceListLessThanOrEqualToTen(_dafny.SeqOf(_dafny.SeqOfString("1")))))
		}(_pat_let110_0)
	}(_184_input)
	var _out60 Wrappers.Result
	_ = _out60
	_out60 = (client).GetConstraints(_184_input)
	_187_ret = _out60
	if !((_187_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(359,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_184_input = func(_pat_let112_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_190_dt__update__tmp_h2 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let113_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_191_dt__update_hListLessThanOrEqualToTen_h2 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_190_dt__update__tmp_h2).Dtor_MyString(), (_190_dt__update__tmp_h2).Dtor_NonEmptyString(), (_190_dt__update__tmp_h2).Dtor_StringLessThanOrEqualToTen(), (_190_dt__update__tmp_h2).Dtor_MyBlob(), (_190_dt__update__tmp_h2).Dtor_NonEmptyBlob(), (_190_dt__update__tmp_h2).Dtor_BlobLessThanOrEqualToTen(), (_190_dt__update__tmp_h2).Dtor_MyList(), (_190_dt__update__tmp_h2).Dtor_NonEmptyList(), _191_dt__update_hListLessThanOrEqualToTen_h2, (_190_dt__update__tmp_h2).Dtor_MyMap(), (_190_dt__update__tmp_h2).Dtor_NonEmptyMap(), (_190_dt__update__tmp_h2).Dtor_MapLessThanOrEqualToTen(), (_190_dt__update__tmp_h2).Dtor_OneToTen(), (_190_dt__update__tmp_h2).Dtor_myTenToTen(), (_190_dt__update__tmp_h2).Dtor_GreaterThanOne(), (_190_dt__update__tmp_h2).Dtor_LessThanTen(), (_190_dt__update__tmp_h2).Dtor_MyUtf8Bytes(), (_190_dt__update__tmp_h2).Dtor_MyListOfUtf8Bytes())
				}(_pat_let113_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceListLessThanOrEqualToTen(_dafny.SeqOf(_dafny.SeqOfString("1"), _dafny.SeqOfString("2"), _dafny.SeqOfString("3"), _dafny.SeqOfString("4"), _dafny.SeqOfString("5"), _dafny.SeqOfString("6"), _dafny.SeqOfString("7"), _dafny.SeqOfString("8"), _dafny.SeqOfString("9"), _dafny.SeqOfString("0")))))
		}(_pat_let112_0)
	}(_184_input)
	var _out61 Wrappers.Result
	_ = _out61
	_out61 = (client).GetConstraints(_184_input)
	_187_ret = _out61
	if !((_187_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(363,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_184_input = func(_pat_let114_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_192_dt__update__tmp_h3 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let115_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_193_dt__update_hListLessThanOrEqualToTen_h3 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_192_dt__update__tmp_h3).Dtor_MyString(), (_192_dt__update__tmp_h3).Dtor_NonEmptyString(), (_192_dt__update__tmp_h3).Dtor_StringLessThanOrEqualToTen(), (_192_dt__update__tmp_h3).Dtor_MyBlob(), (_192_dt__update__tmp_h3).Dtor_NonEmptyBlob(), (_192_dt__update__tmp_h3).Dtor_BlobLessThanOrEqualToTen(), (_192_dt__update__tmp_h3).Dtor_MyList(), (_192_dt__update__tmp_h3).Dtor_NonEmptyList(), _193_dt__update_hListLessThanOrEqualToTen_h3, (_192_dt__update__tmp_h3).Dtor_MyMap(), (_192_dt__update__tmp_h3).Dtor_NonEmptyMap(), (_192_dt__update__tmp_h3).Dtor_MapLessThanOrEqualToTen(), (_192_dt__update__tmp_h3).Dtor_OneToTen(), (_192_dt__update__tmp_h3).Dtor_myTenToTen(), (_192_dt__update__tmp_h3).Dtor_GreaterThanOne(), (_192_dt__update__tmp_h3).Dtor_LessThanTen(), (_192_dt__update__tmp_h3).Dtor_MyUtf8Bytes(), (_192_dt__update__tmp_h3).Dtor_MyListOfUtf8Bytes())
				}(_pat_let115_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceListLessThanOrEqualToTen(_dafny.SeqOf(_dafny.SeqOfString("1"), _dafny.SeqOfString("2"), _dafny.SeqOfString("3"), _dafny.SeqOfString("4"), _dafny.SeqOfString("5"), _dafny.SeqOfString("6"), _dafny.SeqOfString("7"), _dafny.SeqOfString("8"), _dafny.SeqOfString("9"), _dafny.SeqOfString("0"), _dafny.SeqOfString("a")))))
		}(_pat_let114_0)
	}(_184_input)
	var _out62 Wrappers.Result
	_ = _out62
	_out62 = (client).GetConstraints(_184_input)
	_187_ret = _out62
	if !((_187_ret).Is_Failure()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(367,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
}
func (_static *CompanionStruct_Default___) TestGetConstraintWithMyMap(client simpleconstraintsinternaldafnytypes.ISimpleConstraintsClient) {
	var _194_input simpleconstraintsinternaldafnytypes.GetConstraintsInput
	_ = _194_input
	_194_input = Helpers.Companion_Default___.GetValidInput()
	_194_input = func(_pat_let116_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_195_dt__update__tmp_h0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let117_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_196_dt__update_hMyMap_h0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_195_dt__update__tmp_h0).Dtor_MyString(), (_195_dt__update__tmp_h0).Dtor_NonEmptyString(), (_195_dt__update__tmp_h0).Dtor_StringLessThanOrEqualToTen(), (_195_dt__update__tmp_h0).Dtor_MyBlob(), (_195_dt__update__tmp_h0).Dtor_NonEmptyBlob(), (_195_dt__update__tmp_h0).Dtor_BlobLessThanOrEqualToTen(), (_195_dt__update__tmp_h0).Dtor_MyList(), (_195_dt__update__tmp_h0).Dtor_NonEmptyList(), (_195_dt__update__tmp_h0).Dtor_ListLessThanOrEqualToTen(), _196_dt__update_hMyMap_h0, (_195_dt__update__tmp_h0).Dtor_NonEmptyMap(), (_195_dt__update__tmp_h0).Dtor_MapLessThanOrEqualToTen(), (_195_dt__update__tmp_h0).Dtor_OneToTen(), (_195_dt__update__tmp_h0).Dtor_myTenToTen(), (_195_dt__update__tmp_h0).Dtor_GreaterThanOne(), (_195_dt__update__tmp_h0).Dtor_LessThanTen(), (_195_dt__update__tmp_h0).Dtor_MyUtf8Bytes(), (_195_dt__update__tmp_h0).Dtor_MyListOfUtf8Bytes())
				}(_pat_let117_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceMyMap(_dafny.NewMapBuilder().ToMap())))
		}(_pat_let116_0)
	}(_194_input)
	var _197_ret Wrappers.Result
	_ = _197_ret
	var _out63 Wrappers.Result
	_ = _out63
	_out63 = (client).GetConstraints(_194_input)
	_197_ret = _out63
	if !((_197_ret).Is_Failure()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(379,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_194_input = func(_pat_let118_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_198_dt__update__tmp_h1 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let119_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_199_dt__update_hMyMap_h1 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_198_dt__update__tmp_h1).Dtor_MyString(), (_198_dt__update__tmp_h1).Dtor_NonEmptyString(), (_198_dt__update__tmp_h1).Dtor_StringLessThanOrEqualToTen(), (_198_dt__update__tmp_h1).Dtor_MyBlob(), (_198_dt__update__tmp_h1).Dtor_NonEmptyBlob(), (_198_dt__update__tmp_h1).Dtor_BlobLessThanOrEqualToTen(), (_198_dt__update__tmp_h1).Dtor_MyList(), (_198_dt__update__tmp_h1).Dtor_NonEmptyList(), (_198_dt__update__tmp_h1).Dtor_ListLessThanOrEqualToTen(), _199_dt__update_hMyMap_h1, (_198_dt__update__tmp_h1).Dtor_NonEmptyMap(), (_198_dt__update__tmp_h1).Dtor_MapLessThanOrEqualToTen(), (_198_dt__update__tmp_h1).Dtor_OneToTen(), (_198_dt__update__tmp_h1).Dtor_myTenToTen(), (_198_dt__update__tmp_h1).Dtor_GreaterThanOne(), (_198_dt__update__tmp_h1).Dtor_LessThanTen(), (_198_dt__update__tmp_h1).Dtor_MyUtf8Bytes(), (_198_dt__update__tmp_h1).Dtor_MyListOfUtf8Bytes())
				}(_pat_let119_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceMyMap(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOfString("1"), _dafny.SeqOfString("a")))))
		}(_pat_let118_0)
	}(_194_input)
	var _out64 Wrappers.Result
	_ = _out64
	_out64 = (client).GetConstraints(_194_input)
	_197_ret = _out64
	if !((_197_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(383,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_194_input = func(_pat_let120_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_200_dt__update__tmp_h2 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let121_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_201_dt__update_hMyMap_h2 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_200_dt__update__tmp_h2).Dtor_MyString(), (_200_dt__update__tmp_h2).Dtor_NonEmptyString(), (_200_dt__update__tmp_h2).Dtor_StringLessThanOrEqualToTen(), (_200_dt__update__tmp_h2).Dtor_MyBlob(), (_200_dt__update__tmp_h2).Dtor_NonEmptyBlob(), (_200_dt__update__tmp_h2).Dtor_BlobLessThanOrEqualToTen(), (_200_dt__update__tmp_h2).Dtor_MyList(), (_200_dt__update__tmp_h2).Dtor_NonEmptyList(), (_200_dt__update__tmp_h2).Dtor_ListLessThanOrEqualToTen(), _201_dt__update_hMyMap_h2, (_200_dt__update__tmp_h2).Dtor_NonEmptyMap(), (_200_dt__update__tmp_h2).Dtor_MapLessThanOrEqualToTen(), (_200_dt__update__tmp_h2).Dtor_OneToTen(), (_200_dt__update__tmp_h2).Dtor_myTenToTen(), (_200_dt__update__tmp_h2).Dtor_GreaterThanOne(), (_200_dt__update__tmp_h2).Dtor_LessThanTen(), (_200_dt__update__tmp_h2).Dtor_MyUtf8Bytes(), (_200_dt__update__tmp_h2).Dtor_MyListOfUtf8Bytes())
				}(_pat_let121_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceMyMap(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOfString("1"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("2"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("3"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("4"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("5"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("6"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("7"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("8"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("9"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("0"), _dafny.SeqOfString("a")))))
		}(_pat_let120_0)
	}(_194_input)
	var _out65 Wrappers.Result
	_ = _out65
	_out65 = (client).GetConstraints(_194_input)
	_197_ret = _out65
	if !((_197_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(387,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_194_input = func(_pat_let122_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_202_dt__update__tmp_h3 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let123_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_203_dt__update_hMyMap_h3 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_202_dt__update__tmp_h3).Dtor_MyString(), (_202_dt__update__tmp_h3).Dtor_NonEmptyString(), (_202_dt__update__tmp_h3).Dtor_StringLessThanOrEqualToTen(), (_202_dt__update__tmp_h3).Dtor_MyBlob(), (_202_dt__update__tmp_h3).Dtor_NonEmptyBlob(), (_202_dt__update__tmp_h3).Dtor_BlobLessThanOrEqualToTen(), (_202_dt__update__tmp_h3).Dtor_MyList(), (_202_dt__update__tmp_h3).Dtor_NonEmptyList(), (_202_dt__update__tmp_h3).Dtor_ListLessThanOrEqualToTen(), _203_dt__update_hMyMap_h3, (_202_dt__update__tmp_h3).Dtor_NonEmptyMap(), (_202_dt__update__tmp_h3).Dtor_MapLessThanOrEqualToTen(), (_202_dt__update__tmp_h3).Dtor_OneToTen(), (_202_dt__update__tmp_h3).Dtor_myTenToTen(), (_202_dt__update__tmp_h3).Dtor_GreaterThanOne(), (_202_dt__update__tmp_h3).Dtor_LessThanTen(), (_202_dt__update__tmp_h3).Dtor_MyUtf8Bytes(), (_202_dt__update__tmp_h3).Dtor_MyListOfUtf8Bytes())
				}(_pat_let123_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceMyMap(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOfString("1"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("2"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("3"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("4"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("5"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("6"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("7"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("8"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("9"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("0"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("a"), _dafny.SeqOfString("a")))))
		}(_pat_let122_0)
	}(_194_input)
	var _out66 Wrappers.Result
	_ = _out66
	_out66 = (client).GetConstraints(_194_input)
	_197_ret = _out66
	if !((_197_ret).Is_Failure()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(391,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
}
func (_static *CompanionStruct_Default___) TestGetConstraintWithNonEmptyMap(client simpleconstraintsinternaldafnytypes.ISimpleConstraintsClient) {
	var _204_input simpleconstraintsinternaldafnytypes.GetConstraintsInput
	_ = _204_input
	_204_input = Helpers.Companion_Default___.GetValidInput()
	_204_input = func(_pat_let124_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_205_dt__update__tmp_h0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let125_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_206_dt__update_hNonEmptyMap_h0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_205_dt__update__tmp_h0).Dtor_MyString(), (_205_dt__update__tmp_h0).Dtor_NonEmptyString(), (_205_dt__update__tmp_h0).Dtor_StringLessThanOrEqualToTen(), (_205_dt__update__tmp_h0).Dtor_MyBlob(), (_205_dt__update__tmp_h0).Dtor_NonEmptyBlob(), (_205_dt__update__tmp_h0).Dtor_BlobLessThanOrEqualToTen(), (_205_dt__update__tmp_h0).Dtor_MyList(), (_205_dt__update__tmp_h0).Dtor_NonEmptyList(), (_205_dt__update__tmp_h0).Dtor_ListLessThanOrEqualToTen(), (_205_dt__update__tmp_h0).Dtor_MyMap(), _206_dt__update_hNonEmptyMap_h0, (_205_dt__update__tmp_h0).Dtor_MapLessThanOrEqualToTen(), (_205_dt__update__tmp_h0).Dtor_OneToTen(), (_205_dt__update__tmp_h0).Dtor_myTenToTen(), (_205_dt__update__tmp_h0).Dtor_GreaterThanOne(), (_205_dt__update__tmp_h0).Dtor_LessThanTen(), (_205_dt__update__tmp_h0).Dtor_MyUtf8Bytes(), (_205_dt__update__tmp_h0).Dtor_MyListOfUtf8Bytes())
				}(_pat_let125_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceNonEmptyMap(_dafny.NewMapBuilder().ToMap())))
		}(_pat_let124_0)
	}(_204_input)
	var _207_ret Wrappers.Result
	_ = _207_ret
	var _out67 Wrappers.Result
	_ = _out67
	_out67 = (client).GetConstraints(_204_input)
	_207_ret = _out67
	if !((_207_ret).Is_Failure()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(403,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_204_input = func(_pat_let126_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_208_dt__update__tmp_h1 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let127_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_209_dt__update_hNonEmptyMap_h1 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_208_dt__update__tmp_h1).Dtor_MyString(), (_208_dt__update__tmp_h1).Dtor_NonEmptyString(), (_208_dt__update__tmp_h1).Dtor_StringLessThanOrEqualToTen(), (_208_dt__update__tmp_h1).Dtor_MyBlob(), (_208_dt__update__tmp_h1).Dtor_NonEmptyBlob(), (_208_dt__update__tmp_h1).Dtor_BlobLessThanOrEqualToTen(), (_208_dt__update__tmp_h1).Dtor_MyList(), (_208_dt__update__tmp_h1).Dtor_NonEmptyList(), (_208_dt__update__tmp_h1).Dtor_ListLessThanOrEqualToTen(), (_208_dt__update__tmp_h1).Dtor_MyMap(), _209_dt__update_hNonEmptyMap_h1, (_208_dt__update__tmp_h1).Dtor_MapLessThanOrEqualToTen(), (_208_dt__update__tmp_h1).Dtor_OneToTen(), (_208_dt__update__tmp_h1).Dtor_myTenToTen(), (_208_dt__update__tmp_h1).Dtor_GreaterThanOne(), (_208_dt__update__tmp_h1).Dtor_LessThanTen(), (_208_dt__update__tmp_h1).Dtor_MyUtf8Bytes(), (_208_dt__update__tmp_h1).Dtor_MyListOfUtf8Bytes())
				}(_pat_let127_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceNonEmptyMap(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOfString("1"), _dafny.SeqOfString("a")))))
		}(_pat_let126_0)
	}(_204_input)
	var _out68 Wrappers.Result
	_ = _out68
	_out68 = (client).GetConstraints(_204_input)
	_207_ret = _out68
	if !((_207_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(407,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_204_input = func(_pat_let128_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_210_dt__update__tmp_h2 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let129_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_211_dt__update_hNonEmptyMap_h2 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_210_dt__update__tmp_h2).Dtor_MyString(), (_210_dt__update__tmp_h2).Dtor_NonEmptyString(), (_210_dt__update__tmp_h2).Dtor_StringLessThanOrEqualToTen(), (_210_dt__update__tmp_h2).Dtor_MyBlob(), (_210_dt__update__tmp_h2).Dtor_NonEmptyBlob(), (_210_dt__update__tmp_h2).Dtor_BlobLessThanOrEqualToTen(), (_210_dt__update__tmp_h2).Dtor_MyList(), (_210_dt__update__tmp_h2).Dtor_NonEmptyList(), (_210_dt__update__tmp_h2).Dtor_ListLessThanOrEqualToTen(), (_210_dt__update__tmp_h2).Dtor_MyMap(), _211_dt__update_hNonEmptyMap_h2, (_210_dt__update__tmp_h2).Dtor_MapLessThanOrEqualToTen(), (_210_dt__update__tmp_h2).Dtor_OneToTen(), (_210_dt__update__tmp_h2).Dtor_myTenToTen(), (_210_dt__update__tmp_h2).Dtor_GreaterThanOne(), (_210_dt__update__tmp_h2).Dtor_LessThanTen(), (_210_dt__update__tmp_h2).Dtor_MyUtf8Bytes(), (_210_dt__update__tmp_h2).Dtor_MyListOfUtf8Bytes())
				}(_pat_let129_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceNonEmptyMap(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOfString("1"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("2"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("3"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("4"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("5"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("6"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("7"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("8"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("9"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("0"), _dafny.SeqOfString("a")))))
		}(_pat_let128_0)
	}(_204_input)
	var _out69 Wrappers.Result
	_ = _out69
	_out69 = (client).GetConstraints(_204_input)
	_207_ret = _out69
	if !((_207_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(411,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_204_input = func(_pat_let130_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_212_dt__update__tmp_h3 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let131_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_213_dt__update_hNonEmptyMap_h3 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_212_dt__update__tmp_h3).Dtor_MyString(), (_212_dt__update__tmp_h3).Dtor_NonEmptyString(), (_212_dt__update__tmp_h3).Dtor_StringLessThanOrEqualToTen(), (_212_dt__update__tmp_h3).Dtor_MyBlob(), (_212_dt__update__tmp_h3).Dtor_NonEmptyBlob(), (_212_dt__update__tmp_h3).Dtor_BlobLessThanOrEqualToTen(), (_212_dt__update__tmp_h3).Dtor_MyList(), (_212_dt__update__tmp_h3).Dtor_NonEmptyList(), (_212_dt__update__tmp_h3).Dtor_ListLessThanOrEqualToTen(), (_212_dt__update__tmp_h3).Dtor_MyMap(), _213_dt__update_hNonEmptyMap_h3, (_212_dt__update__tmp_h3).Dtor_MapLessThanOrEqualToTen(), (_212_dt__update__tmp_h3).Dtor_OneToTen(), (_212_dt__update__tmp_h3).Dtor_myTenToTen(), (_212_dt__update__tmp_h3).Dtor_GreaterThanOne(), (_212_dt__update__tmp_h3).Dtor_LessThanTen(), (_212_dt__update__tmp_h3).Dtor_MyUtf8Bytes(), (_212_dt__update__tmp_h3).Dtor_MyListOfUtf8Bytes())
				}(_pat_let131_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceNonEmptyMap(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOfString("1"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("2"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("3"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("4"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("5"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("6"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("7"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("8"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("9"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("0"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("a"), _dafny.SeqOfString("a")))))
		}(_pat_let130_0)
	}(_204_input)
	var _out70 Wrappers.Result
	_ = _out70
	_out70 = (client).GetConstraints(_204_input)
	_207_ret = _out70
	if !((_207_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(415,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
}
func (_static *CompanionStruct_Default___) TestGetConstraintWithMapLessThanOrEqualToTen(client simpleconstraintsinternaldafnytypes.ISimpleConstraintsClient) {
	var _214_input simpleconstraintsinternaldafnytypes.GetConstraintsInput
	_ = _214_input
	_214_input = Helpers.Companion_Default___.GetValidInput()
	_214_input = func(_pat_let132_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_215_dt__update__tmp_h0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let133_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_216_dt__update_hMapLessThanOrEqualToTen_h0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_215_dt__update__tmp_h0).Dtor_MyString(), (_215_dt__update__tmp_h0).Dtor_NonEmptyString(), (_215_dt__update__tmp_h0).Dtor_StringLessThanOrEqualToTen(), (_215_dt__update__tmp_h0).Dtor_MyBlob(), (_215_dt__update__tmp_h0).Dtor_NonEmptyBlob(), (_215_dt__update__tmp_h0).Dtor_BlobLessThanOrEqualToTen(), (_215_dt__update__tmp_h0).Dtor_MyList(), (_215_dt__update__tmp_h0).Dtor_NonEmptyList(), (_215_dt__update__tmp_h0).Dtor_ListLessThanOrEqualToTen(), (_215_dt__update__tmp_h0).Dtor_MyMap(), (_215_dt__update__tmp_h0).Dtor_NonEmptyMap(), _216_dt__update_hMapLessThanOrEqualToTen_h0, (_215_dt__update__tmp_h0).Dtor_OneToTen(), (_215_dt__update__tmp_h0).Dtor_myTenToTen(), (_215_dt__update__tmp_h0).Dtor_GreaterThanOne(), (_215_dt__update__tmp_h0).Dtor_LessThanTen(), (_215_dt__update__tmp_h0).Dtor_MyUtf8Bytes(), (_215_dt__update__tmp_h0).Dtor_MyListOfUtf8Bytes())
				}(_pat_let133_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceMapLessThanOrEqualToTen(_dafny.NewMapBuilder().ToMap())))
		}(_pat_let132_0)
	}(_214_input)
	var _217_ret Wrappers.Result
	_ = _217_ret
	var _out71 Wrappers.Result
	_ = _out71
	_out71 = (client).GetConstraints(_214_input)
	_217_ret = _out71
	if !((_217_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(427,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_214_input = func(_pat_let134_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_218_dt__update__tmp_h1 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let135_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_219_dt__update_hMapLessThanOrEqualToTen_h1 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_218_dt__update__tmp_h1).Dtor_MyString(), (_218_dt__update__tmp_h1).Dtor_NonEmptyString(), (_218_dt__update__tmp_h1).Dtor_StringLessThanOrEqualToTen(), (_218_dt__update__tmp_h1).Dtor_MyBlob(), (_218_dt__update__tmp_h1).Dtor_NonEmptyBlob(), (_218_dt__update__tmp_h1).Dtor_BlobLessThanOrEqualToTen(), (_218_dt__update__tmp_h1).Dtor_MyList(), (_218_dt__update__tmp_h1).Dtor_NonEmptyList(), (_218_dt__update__tmp_h1).Dtor_ListLessThanOrEqualToTen(), (_218_dt__update__tmp_h1).Dtor_MyMap(), (_218_dt__update__tmp_h1).Dtor_NonEmptyMap(), _219_dt__update_hMapLessThanOrEqualToTen_h1, (_218_dt__update__tmp_h1).Dtor_OneToTen(), (_218_dt__update__tmp_h1).Dtor_myTenToTen(), (_218_dt__update__tmp_h1).Dtor_GreaterThanOne(), (_218_dt__update__tmp_h1).Dtor_LessThanTen(), (_218_dt__update__tmp_h1).Dtor_MyUtf8Bytes(), (_218_dt__update__tmp_h1).Dtor_MyListOfUtf8Bytes())
				}(_pat_let135_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceMapLessThanOrEqualToTen(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOfString("1"), _dafny.SeqOfString("a")))))
		}(_pat_let134_0)
	}(_214_input)
	var _out72 Wrappers.Result
	_ = _out72
	_out72 = (client).GetConstraints(_214_input)
	_217_ret = _out72
	if !((_217_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(431,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_214_input = func(_pat_let136_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_220_dt__update__tmp_h2 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let137_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_221_dt__update_hMapLessThanOrEqualToTen_h2 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_220_dt__update__tmp_h2).Dtor_MyString(), (_220_dt__update__tmp_h2).Dtor_NonEmptyString(), (_220_dt__update__tmp_h2).Dtor_StringLessThanOrEqualToTen(), (_220_dt__update__tmp_h2).Dtor_MyBlob(), (_220_dt__update__tmp_h2).Dtor_NonEmptyBlob(), (_220_dt__update__tmp_h2).Dtor_BlobLessThanOrEqualToTen(), (_220_dt__update__tmp_h2).Dtor_MyList(), (_220_dt__update__tmp_h2).Dtor_NonEmptyList(), (_220_dt__update__tmp_h2).Dtor_ListLessThanOrEqualToTen(), (_220_dt__update__tmp_h2).Dtor_MyMap(), (_220_dt__update__tmp_h2).Dtor_NonEmptyMap(), _221_dt__update_hMapLessThanOrEqualToTen_h2, (_220_dt__update__tmp_h2).Dtor_OneToTen(), (_220_dt__update__tmp_h2).Dtor_myTenToTen(), (_220_dt__update__tmp_h2).Dtor_GreaterThanOne(), (_220_dt__update__tmp_h2).Dtor_LessThanTen(), (_220_dt__update__tmp_h2).Dtor_MyUtf8Bytes(), (_220_dt__update__tmp_h2).Dtor_MyListOfUtf8Bytes())
				}(_pat_let137_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceMapLessThanOrEqualToTen(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOfString("1"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("2"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("3"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("4"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("5"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("6"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("7"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("8"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("9"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("0"), _dafny.SeqOfString("a")))))
		}(_pat_let136_0)
	}(_214_input)
	var _out73 Wrappers.Result
	_ = _out73
	_out73 = (client).GetConstraints(_214_input)
	_217_ret = _out73
	if !((_217_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(435,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_214_input = func(_pat_let138_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_222_dt__update__tmp_h3 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let139_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_223_dt__update_hMapLessThanOrEqualToTen_h3 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_222_dt__update__tmp_h3).Dtor_MyString(), (_222_dt__update__tmp_h3).Dtor_NonEmptyString(), (_222_dt__update__tmp_h3).Dtor_StringLessThanOrEqualToTen(), (_222_dt__update__tmp_h3).Dtor_MyBlob(), (_222_dt__update__tmp_h3).Dtor_NonEmptyBlob(), (_222_dt__update__tmp_h3).Dtor_BlobLessThanOrEqualToTen(), (_222_dt__update__tmp_h3).Dtor_MyList(), (_222_dt__update__tmp_h3).Dtor_NonEmptyList(), (_222_dt__update__tmp_h3).Dtor_ListLessThanOrEqualToTen(), (_222_dt__update__tmp_h3).Dtor_MyMap(), (_222_dt__update__tmp_h3).Dtor_NonEmptyMap(), _223_dt__update_hMapLessThanOrEqualToTen_h3, (_222_dt__update__tmp_h3).Dtor_OneToTen(), (_222_dt__update__tmp_h3).Dtor_myTenToTen(), (_222_dt__update__tmp_h3).Dtor_GreaterThanOne(), (_222_dt__update__tmp_h3).Dtor_LessThanTen(), (_222_dt__update__tmp_h3).Dtor_MyUtf8Bytes(), (_222_dt__update__tmp_h3).Dtor_MyListOfUtf8Bytes())
				}(_pat_let139_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceMapLessThanOrEqualToTen(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOfString("1"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("2"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("3"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("4"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("5"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("6"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("7"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("8"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("9"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("0"), _dafny.SeqOfString("a")).UpdateUnsafe(_dafny.SeqOfString("a"), _dafny.SeqOfString("a")))))
		}(_pat_let138_0)
	}(_214_input)
	var _out74 Wrappers.Result
	_ = _out74
	_out74 = (client).GetConstraints(_214_input)
	_217_ret = _out74
	if !((_217_ret).Is_Failure()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(439,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
}
func (_static *CompanionStruct_Default___) TestGetConstraintWithGreaterThanOne(client simpleconstraintsinternaldafnytypes.ISimpleConstraintsClient) {
	var _224_input simpleconstraintsinternaldafnytypes.GetConstraintsInput
	_ = _224_input
	_224_input = Helpers.Companion_Default___.GetValidInput()
	_224_input = func(_pat_let140_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_225_dt__update__tmp_h0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let141_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_226_dt__update_hGreaterThanOne_h0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_225_dt__update__tmp_h0).Dtor_MyString(), (_225_dt__update__tmp_h0).Dtor_NonEmptyString(), (_225_dt__update__tmp_h0).Dtor_StringLessThanOrEqualToTen(), (_225_dt__update__tmp_h0).Dtor_MyBlob(), (_225_dt__update__tmp_h0).Dtor_NonEmptyBlob(), (_225_dt__update__tmp_h0).Dtor_BlobLessThanOrEqualToTen(), (_225_dt__update__tmp_h0).Dtor_MyList(), (_225_dt__update__tmp_h0).Dtor_NonEmptyList(), (_225_dt__update__tmp_h0).Dtor_ListLessThanOrEqualToTen(), (_225_dt__update__tmp_h0).Dtor_MyMap(), (_225_dt__update__tmp_h0).Dtor_NonEmptyMap(), (_225_dt__update__tmp_h0).Dtor_MapLessThanOrEqualToTen(), (_225_dt__update__tmp_h0).Dtor_OneToTen(), (_225_dt__update__tmp_h0).Dtor_myTenToTen(), _226_dt__update_hGreaterThanOne_h0, (_225_dt__update__tmp_h0).Dtor_LessThanTen(), (_225_dt__update__tmp_h0).Dtor_MyUtf8Bytes(), (_225_dt__update__tmp_h0).Dtor_MyListOfUtf8Bytes())
				}(_pat_let141_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceGreaterThanOne(_dafny.IntOfInt64(1000))))
		}(_pat_let140_0)
	}(_224_input)
	var _227_ret Wrappers.Result
	_ = _227_ret
	var _out75 Wrappers.Result
	_ = _out75
	_out75 = (client).GetConstraints(_224_input)
	_227_ret = _out75
	if !((_227_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(451,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_224_input = func(_pat_let142_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_228_dt__update__tmp_h1 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let143_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_229_dt__update_hGreaterThanOne_h1 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_228_dt__update__tmp_h1).Dtor_MyString(), (_228_dt__update__tmp_h1).Dtor_NonEmptyString(), (_228_dt__update__tmp_h1).Dtor_StringLessThanOrEqualToTen(), (_228_dt__update__tmp_h1).Dtor_MyBlob(), (_228_dt__update__tmp_h1).Dtor_NonEmptyBlob(), (_228_dt__update__tmp_h1).Dtor_BlobLessThanOrEqualToTen(), (_228_dt__update__tmp_h1).Dtor_MyList(), (_228_dt__update__tmp_h1).Dtor_NonEmptyList(), (_228_dt__update__tmp_h1).Dtor_ListLessThanOrEqualToTen(), (_228_dt__update__tmp_h1).Dtor_MyMap(), (_228_dt__update__tmp_h1).Dtor_NonEmptyMap(), (_228_dt__update__tmp_h1).Dtor_MapLessThanOrEqualToTen(), (_228_dt__update__tmp_h1).Dtor_OneToTen(), (_228_dt__update__tmp_h1).Dtor_myTenToTen(), _229_dt__update_hGreaterThanOne_h1, (_228_dt__update__tmp_h1).Dtor_LessThanTen(), (_228_dt__update__tmp_h1).Dtor_MyUtf8Bytes(), (_228_dt__update__tmp_h1).Dtor_MyListOfUtf8Bytes())
				}(_pat_let143_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceGreaterThanOne(_dafny.IntOfInt64(-1000))))
		}(_pat_let142_0)
	}(_224_input)
	var _out76 Wrappers.Result
	_ = _out76
	_out76 = (client).GetConstraints(_224_input)
	_227_ret = _out76
	if !((_227_ret).Is_Failure()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(455,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_224_input = func(_pat_let144_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_230_dt__update__tmp_h2 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let145_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_231_dt__update_hGreaterThanOne_h2 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_230_dt__update__tmp_h2).Dtor_MyString(), (_230_dt__update__tmp_h2).Dtor_NonEmptyString(), (_230_dt__update__tmp_h2).Dtor_StringLessThanOrEqualToTen(), (_230_dt__update__tmp_h2).Dtor_MyBlob(), (_230_dt__update__tmp_h2).Dtor_NonEmptyBlob(), (_230_dt__update__tmp_h2).Dtor_BlobLessThanOrEqualToTen(), (_230_dt__update__tmp_h2).Dtor_MyList(), (_230_dt__update__tmp_h2).Dtor_NonEmptyList(), (_230_dt__update__tmp_h2).Dtor_ListLessThanOrEqualToTen(), (_230_dt__update__tmp_h2).Dtor_MyMap(), (_230_dt__update__tmp_h2).Dtor_NonEmptyMap(), (_230_dt__update__tmp_h2).Dtor_MapLessThanOrEqualToTen(), (_230_dt__update__tmp_h2).Dtor_OneToTen(), (_230_dt__update__tmp_h2).Dtor_myTenToTen(), _231_dt__update_hGreaterThanOne_h2, (_230_dt__update__tmp_h2).Dtor_LessThanTen(), (_230_dt__update__tmp_h2).Dtor_MyUtf8Bytes(), (_230_dt__update__tmp_h2).Dtor_MyListOfUtf8Bytes())
				}(_pat_let145_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceGreaterThanOne(_dafny.Zero)))
		}(_pat_let144_0)
	}(_224_input)
	var _out77 Wrappers.Result
	_ = _out77
	_out77 = (client).GetConstraints(_224_input)
	_227_ret = _out77
	if !((_227_ret).Is_Failure()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(459,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_224_input = func(_pat_let146_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_232_dt__update__tmp_h3 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let147_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_233_dt__update_hGreaterThanOne_h3 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_232_dt__update__tmp_h3).Dtor_MyString(), (_232_dt__update__tmp_h3).Dtor_NonEmptyString(), (_232_dt__update__tmp_h3).Dtor_StringLessThanOrEqualToTen(), (_232_dt__update__tmp_h3).Dtor_MyBlob(), (_232_dt__update__tmp_h3).Dtor_NonEmptyBlob(), (_232_dt__update__tmp_h3).Dtor_BlobLessThanOrEqualToTen(), (_232_dt__update__tmp_h3).Dtor_MyList(), (_232_dt__update__tmp_h3).Dtor_NonEmptyList(), (_232_dt__update__tmp_h3).Dtor_ListLessThanOrEqualToTen(), (_232_dt__update__tmp_h3).Dtor_MyMap(), (_232_dt__update__tmp_h3).Dtor_NonEmptyMap(), (_232_dt__update__tmp_h3).Dtor_MapLessThanOrEqualToTen(), (_232_dt__update__tmp_h3).Dtor_OneToTen(), (_232_dt__update__tmp_h3).Dtor_myTenToTen(), _233_dt__update_hGreaterThanOne_h3, (_232_dt__update__tmp_h3).Dtor_LessThanTen(), (_232_dt__update__tmp_h3).Dtor_MyUtf8Bytes(), (_232_dt__update__tmp_h3).Dtor_MyListOfUtf8Bytes())
				}(_pat_let147_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceGreaterThanOne(_dafny.One)))
		}(_pat_let146_0)
	}(_224_input)
	var _out78 Wrappers.Result
	_ = _out78
	_out78 = (client).GetConstraints(_224_input)
	_227_ret = _out78
	if !((_227_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(463,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
}
func (_static *CompanionStruct_Default___) TestGetConstraintWithUtf8Bytes(client simpleconstraintsinternaldafnytypes.ISimpleConstraintsClient) {
	var _234_input simpleconstraintsinternaldafnytypes.GetConstraintsInput
	_ = _234_input
	_234_input = Helpers.Companion_Default___.GetValidInput()
	_234_input = func(_pat_let148_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_235_dt__update__tmp_h0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let149_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_236_dt__update_hMyUtf8Bytes_h0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_235_dt__update__tmp_h0).Dtor_MyString(), (_235_dt__update__tmp_h0).Dtor_NonEmptyString(), (_235_dt__update__tmp_h0).Dtor_StringLessThanOrEqualToTen(), (_235_dt__update__tmp_h0).Dtor_MyBlob(), (_235_dt__update__tmp_h0).Dtor_NonEmptyBlob(), (_235_dt__update__tmp_h0).Dtor_BlobLessThanOrEqualToTen(), (_235_dt__update__tmp_h0).Dtor_MyList(), (_235_dt__update__tmp_h0).Dtor_NonEmptyList(), (_235_dt__update__tmp_h0).Dtor_ListLessThanOrEqualToTen(), (_235_dt__update__tmp_h0).Dtor_MyMap(), (_235_dt__update__tmp_h0).Dtor_NonEmptyMap(), (_235_dt__update__tmp_h0).Dtor_MapLessThanOrEqualToTen(), (_235_dt__update__tmp_h0).Dtor_OneToTen(), (_235_dt__update__tmp_h0).Dtor_myTenToTen(), (_235_dt__update__tmp_h0).Dtor_GreaterThanOne(), (_235_dt__update__tmp_h0).Dtor_LessThanTen(), _236_dt__update_hMyUtf8Bytes_h0, (_235_dt__update__tmp_h0).Dtor_MyListOfUtf8Bytes())
				}(_pat_let149_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceUtf8Bytes(_dafny.SeqOf())))
		}(_pat_let148_0)
	}(_234_input)
	var _237_ret Wrappers.Result
	_ = _237_ret
	var _out79 Wrappers.Result
	_ = _out79
	_out79 = (client).GetConstraints(_234_input)
	_237_ret = _out79
	if !((_237_ret).Is_Failure()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(475,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_234_input = func(_pat_let150_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_238_dt__update__tmp_h1 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let151_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_239_dt__update_hMyUtf8Bytes_h1 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_238_dt__update__tmp_h1).Dtor_MyString(), (_238_dt__update__tmp_h1).Dtor_NonEmptyString(), (_238_dt__update__tmp_h1).Dtor_StringLessThanOrEqualToTen(), (_238_dt__update__tmp_h1).Dtor_MyBlob(), (_238_dt__update__tmp_h1).Dtor_NonEmptyBlob(), (_238_dt__update__tmp_h1).Dtor_BlobLessThanOrEqualToTen(), (_238_dt__update__tmp_h1).Dtor_MyList(), (_238_dt__update__tmp_h1).Dtor_NonEmptyList(), (_238_dt__update__tmp_h1).Dtor_ListLessThanOrEqualToTen(), (_238_dt__update__tmp_h1).Dtor_MyMap(), (_238_dt__update__tmp_h1).Dtor_NonEmptyMap(), (_238_dt__update__tmp_h1).Dtor_MapLessThanOrEqualToTen(), (_238_dt__update__tmp_h1).Dtor_OneToTen(), (_238_dt__update__tmp_h1).Dtor_myTenToTen(), (_238_dt__update__tmp_h1).Dtor_GreaterThanOne(), (_238_dt__update__tmp_h1).Dtor_LessThanTen(), _239_dt__update_hMyUtf8Bytes_h1, (_238_dt__update__tmp_h1).Dtor_MyListOfUtf8Bytes())
				}(_pat_let151_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceUtf8Bytes(_dafny.SeqOf(uint8(1)))))
		}(_pat_let150_0)
	}(_234_input)
	var _out80 Wrappers.Result
	_ = _out80
	_out80 = (client).GetConstraints(_234_input)
	_237_ret = _out80
	if !((_237_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(479,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_234_input = func(_pat_let152_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_240_dt__update__tmp_h2 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let153_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_241_dt__update_hMyUtf8Bytes_h2 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_240_dt__update__tmp_h2).Dtor_MyString(), (_240_dt__update__tmp_h2).Dtor_NonEmptyString(), (_240_dt__update__tmp_h2).Dtor_StringLessThanOrEqualToTen(), (_240_dt__update__tmp_h2).Dtor_MyBlob(), (_240_dt__update__tmp_h2).Dtor_NonEmptyBlob(), (_240_dt__update__tmp_h2).Dtor_BlobLessThanOrEqualToTen(), (_240_dt__update__tmp_h2).Dtor_MyList(), (_240_dt__update__tmp_h2).Dtor_NonEmptyList(), (_240_dt__update__tmp_h2).Dtor_ListLessThanOrEqualToTen(), (_240_dt__update__tmp_h2).Dtor_MyMap(), (_240_dt__update__tmp_h2).Dtor_NonEmptyMap(), (_240_dt__update__tmp_h2).Dtor_MapLessThanOrEqualToTen(), (_240_dt__update__tmp_h2).Dtor_OneToTen(), (_240_dt__update__tmp_h2).Dtor_myTenToTen(), (_240_dt__update__tmp_h2).Dtor_GreaterThanOne(), (_240_dt__update__tmp_h2).Dtor_LessThanTen(), _241_dt__update_hMyUtf8Bytes_h2, (_240_dt__update__tmp_h2).Dtor_MyListOfUtf8Bytes())
				}(_pat_let153_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceUtf8Bytes(_dafny.SeqOf(uint8(1), uint8(2), uint8(3), uint8(4), uint8(5), uint8(6), uint8(7), uint8(8), uint8(9), uint8(0)))))
		}(_pat_let152_0)
	}(_234_input)
	var _out81 Wrappers.Result
	_ = _out81
	_out81 = (client).GetConstraints(_234_input)
	_237_ret = _out81
	if !((_237_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(483,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_234_input = func(_pat_let154_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_242_dt__update__tmp_h3 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let155_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_243_dt__update_hMyUtf8Bytes_h3 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_242_dt__update__tmp_h3).Dtor_MyString(), (_242_dt__update__tmp_h3).Dtor_NonEmptyString(), (_242_dt__update__tmp_h3).Dtor_StringLessThanOrEqualToTen(), (_242_dt__update__tmp_h3).Dtor_MyBlob(), (_242_dt__update__tmp_h3).Dtor_NonEmptyBlob(), (_242_dt__update__tmp_h3).Dtor_BlobLessThanOrEqualToTen(), (_242_dt__update__tmp_h3).Dtor_MyList(), (_242_dt__update__tmp_h3).Dtor_NonEmptyList(), (_242_dt__update__tmp_h3).Dtor_ListLessThanOrEqualToTen(), (_242_dt__update__tmp_h3).Dtor_MyMap(), (_242_dt__update__tmp_h3).Dtor_NonEmptyMap(), (_242_dt__update__tmp_h3).Dtor_MapLessThanOrEqualToTen(), (_242_dt__update__tmp_h3).Dtor_OneToTen(), (_242_dt__update__tmp_h3).Dtor_myTenToTen(), (_242_dt__update__tmp_h3).Dtor_GreaterThanOne(), (_242_dt__update__tmp_h3).Dtor_LessThanTen(), _243_dt__update_hMyUtf8Bytes_h3, (_242_dt__update__tmp_h3).Dtor_MyListOfUtf8Bytes())
				}(_pat_let155_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceUtf8Bytes(_dafny.SeqOf(uint8(1), uint8(2), uint8(3), uint8(4), uint8(5), uint8(6), uint8(7), uint8(8), uint8(9), uint8(0), uint8(1)))))
		}(_pat_let154_0)
	}(_234_input)
	var _out82 Wrappers.Result
	_ = _out82
	_out82 = (client).GetConstraints(_234_input)
	_237_ret = _out82
	if !((_237_ret).Is_Failure()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(487,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	_234_input = func(_pat_let156_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_244_dt__update__tmp_h4 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let157_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_245_dt__update_hMyUtf8Bytes_h4 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_244_dt__update__tmp_h4).Dtor_MyString(), (_244_dt__update__tmp_h4).Dtor_NonEmptyString(), (_244_dt__update__tmp_h4).Dtor_StringLessThanOrEqualToTen(), (_244_dt__update__tmp_h4).Dtor_MyBlob(), (_244_dt__update__tmp_h4).Dtor_NonEmptyBlob(), (_244_dt__update__tmp_h4).Dtor_BlobLessThanOrEqualToTen(), (_244_dt__update__tmp_h4).Dtor_MyList(), (_244_dt__update__tmp_h4).Dtor_NonEmptyList(), (_244_dt__update__tmp_h4).Dtor_ListLessThanOrEqualToTen(), (_244_dt__update__tmp_h4).Dtor_MyMap(), (_244_dt__update__tmp_h4).Dtor_NonEmptyMap(), (_244_dt__update__tmp_h4).Dtor_MapLessThanOrEqualToTen(), (_244_dt__update__tmp_h4).Dtor_OneToTen(), (_244_dt__update__tmp_h4).Dtor_myTenToTen(), (_244_dt__update__tmp_h4).Dtor_GreaterThanOne(), (_244_dt__update__tmp_h4).Dtor_LessThanTen(), _245_dt__update_hMyUtf8Bytes_h4, (_244_dt__update__tmp_h4).Dtor_MyListOfUtf8Bytes())
				}(_pat_let157_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceUtf8Bytes(_dafny.SeqOf(uint8(255), uint8(255), uint8(255)))))
		}(_pat_let156_0)
	}(_234_input)
	var _out83 Wrappers.Result
	_ = _out83
	_out83 = (client).GetConstraints(_234_input)
	_237_ret = _out83
	if !((_237_ret).Is_Failure()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(492,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	var _246_one _dafny.Sequence
	_ = _246_one
	_246_one = _dafny.SeqOf(uint8(240), uint8(168), uint8(137), uint8(159))
	var _247_two _dafny.Sequence
	_ = _247_two
	_247_two = _dafny.SeqOf(uint8(194), uint8(163))
	var _pat_let_tv0 = _246_one
	_ = _pat_let_tv0
	_234_input = func(_pat_let158_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_248_dt__update__tmp_h5 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let159_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_249_dt__update_hMyUtf8Bytes_h5 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_248_dt__update__tmp_h5).Dtor_MyString(), (_248_dt__update__tmp_h5).Dtor_NonEmptyString(), (_248_dt__update__tmp_h5).Dtor_StringLessThanOrEqualToTen(), (_248_dt__update__tmp_h5).Dtor_MyBlob(), (_248_dt__update__tmp_h5).Dtor_NonEmptyBlob(), (_248_dt__update__tmp_h5).Dtor_BlobLessThanOrEqualToTen(), (_248_dt__update__tmp_h5).Dtor_MyList(), (_248_dt__update__tmp_h5).Dtor_NonEmptyList(), (_248_dt__update__tmp_h5).Dtor_ListLessThanOrEqualToTen(), (_248_dt__update__tmp_h5).Dtor_MyMap(), (_248_dt__update__tmp_h5).Dtor_NonEmptyMap(), (_248_dt__update__tmp_h5).Dtor_MapLessThanOrEqualToTen(), (_248_dt__update__tmp_h5).Dtor_OneToTen(), (_248_dt__update__tmp_h5).Dtor_myTenToTen(), (_248_dt__update__tmp_h5).Dtor_GreaterThanOne(), (_248_dt__update__tmp_h5).Dtor_LessThanTen(), _249_dt__update_hMyUtf8Bytes_h5, (_248_dt__update__tmp_h5).Dtor_MyListOfUtf8Bytes())
				}(_pat_let159_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceUtf8Bytes(_pat_let_tv0)))
		}(_pat_let158_0)
	}(_234_input)
	var _out84 Wrappers.Result
	_ = _out84
	_out84 = (client).GetConstraints(_234_input)
	_237_ret = _out84
	if !((_237_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(498,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	var _pat_let_tv1 = _246_one
	_ = _pat_let_tv1
	var _pat_let_tv2 = _246_one
	_ = _pat_let_tv2
	_234_input = func(_pat_let160_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_250_dt__update__tmp_h6 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let161_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_251_dt__update_hMyUtf8Bytes_h6 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_250_dt__update__tmp_h6).Dtor_MyString(), (_250_dt__update__tmp_h6).Dtor_NonEmptyString(), (_250_dt__update__tmp_h6).Dtor_StringLessThanOrEqualToTen(), (_250_dt__update__tmp_h6).Dtor_MyBlob(), (_250_dt__update__tmp_h6).Dtor_NonEmptyBlob(), (_250_dt__update__tmp_h6).Dtor_BlobLessThanOrEqualToTen(), (_250_dt__update__tmp_h6).Dtor_MyList(), (_250_dt__update__tmp_h6).Dtor_NonEmptyList(), (_250_dt__update__tmp_h6).Dtor_ListLessThanOrEqualToTen(), (_250_dt__update__tmp_h6).Dtor_MyMap(), (_250_dt__update__tmp_h6).Dtor_NonEmptyMap(), (_250_dt__update__tmp_h6).Dtor_MapLessThanOrEqualToTen(), (_250_dt__update__tmp_h6).Dtor_OneToTen(), (_250_dt__update__tmp_h6).Dtor_myTenToTen(), (_250_dt__update__tmp_h6).Dtor_GreaterThanOne(), (_250_dt__update__tmp_h6).Dtor_LessThanTen(), _251_dt__update_hMyUtf8Bytes_h6, (_250_dt__update__tmp_h6).Dtor_MyListOfUtf8Bytes())
				}(_pat_let161_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceUtf8Bytes(_dafny.Companion_Sequence_.Concatenate(_pat_let_tv1, _pat_let_tv2))))
		}(_pat_let160_0)
	}(_234_input)
	var _out85 Wrappers.Result
	_ = _out85
	_out85 = (client).GetConstraints(_234_input)
	_237_ret = _out85
	if !((_237_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(502,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	var _pat_let_tv3 = _246_one
	_ = _pat_let_tv3
	var _pat_let_tv4 = _246_one
	_ = _pat_let_tv4
	var _pat_let_tv5 = _246_one
	_ = _pat_let_tv5
	_234_input = func(_pat_let162_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_252_dt__update__tmp_h7 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let163_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_253_dt__update_hMyUtf8Bytes_h7 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_252_dt__update__tmp_h7).Dtor_MyString(), (_252_dt__update__tmp_h7).Dtor_NonEmptyString(), (_252_dt__update__tmp_h7).Dtor_StringLessThanOrEqualToTen(), (_252_dt__update__tmp_h7).Dtor_MyBlob(), (_252_dt__update__tmp_h7).Dtor_NonEmptyBlob(), (_252_dt__update__tmp_h7).Dtor_BlobLessThanOrEqualToTen(), (_252_dt__update__tmp_h7).Dtor_MyList(), (_252_dt__update__tmp_h7).Dtor_NonEmptyList(), (_252_dt__update__tmp_h7).Dtor_ListLessThanOrEqualToTen(), (_252_dt__update__tmp_h7).Dtor_MyMap(), (_252_dt__update__tmp_h7).Dtor_NonEmptyMap(), (_252_dt__update__tmp_h7).Dtor_MapLessThanOrEqualToTen(), (_252_dt__update__tmp_h7).Dtor_OneToTen(), (_252_dt__update__tmp_h7).Dtor_myTenToTen(), (_252_dt__update__tmp_h7).Dtor_GreaterThanOne(), (_252_dt__update__tmp_h7).Dtor_LessThanTen(), _253_dt__update_hMyUtf8Bytes_h7, (_252_dt__update__tmp_h7).Dtor_MyListOfUtf8Bytes())
				}(_pat_let163_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceUtf8Bytes(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_pat_let_tv3, _pat_let_tv4), _pat_let_tv5))))
		}(_pat_let162_0)
	}(_234_input)
	var _out86 Wrappers.Result
	_ = _out86
	_out86 = (client).GetConstraints(_234_input)
	_237_ret = _out86
	if !((_237_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(507,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	var _pat_let_tv6 = _247_two
	_ = _pat_let_tv6
	var _pat_let_tv7 = _247_two
	_ = _pat_let_tv7
	var _pat_let_tv8 = _247_two
	_ = _pat_let_tv8
	var _pat_let_tv9 = _247_two
	_ = _pat_let_tv9
	var _pat_let_tv10 = _247_two
	_ = _pat_let_tv10
	var _pat_let_tv11 = _247_two
	_ = _pat_let_tv11
	var _pat_let_tv12 = _247_two
	_ = _pat_let_tv12
	var _pat_let_tv13 = _247_two
	_ = _pat_let_tv13
	var _pat_let_tv14 = _247_two
	_ = _pat_let_tv14
	var _pat_let_tv15 = _247_two
	_ = _pat_let_tv15
	_234_input = func(_pat_let164_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_254_dt__update__tmp_h8 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let165_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_255_dt__update_hMyUtf8Bytes_h8 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_254_dt__update__tmp_h8).Dtor_MyString(), (_254_dt__update__tmp_h8).Dtor_NonEmptyString(), (_254_dt__update__tmp_h8).Dtor_StringLessThanOrEqualToTen(), (_254_dt__update__tmp_h8).Dtor_MyBlob(), (_254_dt__update__tmp_h8).Dtor_NonEmptyBlob(), (_254_dt__update__tmp_h8).Dtor_BlobLessThanOrEqualToTen(), (_254_dt__update__tmp_h8).Dtor_MyList(), (_254_dt__update__tmp_h8).Dtor_NonEmptyList(), (_254_dt__update__tmp_h8).Dtor_ListLessThanOrEqualToTen(), (_254_dt__update__tmp_h8).Dtor_MyMap(), (_254_dt__update__tmp_h8).Dtor_NonEmptyMap(), (_254_dt__update__tmp_h8).Dtor_MapLessThanOrEqualToTen(), (_254_dt__update__tmp_h8).Dtor_OneToTen(), (_254_dt__update__tmp_h8).Dtor_myTenToTen(), (_254_dt__update__tmp_h8).Dtor_GreaterThanOne(), (_254_dt__update__tmp_h8).Dtor_LessThanTen(), _255_dt__update_hMyUtf8Bytes_h8, (_254_dt__update__tmp_h8).Dtor_MyListOfUtf8Bytes())
				}(_pat_let165_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceUtf8Bytes(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_pat_let_tv6, _pat_let_tv7), _pat_let_tv8), _pat_let_tv9), _pat_let_tv10), _pat_let_tv11), _pat_let_tv12), _pat_let_tv13), _pat_let_tv14), _pat_let_tv15))))
		}(_pat_let164_0)
	}(_234_input)
	var _out87 Wrappers.Result
	_ = _out87
	_out87 = (client).GetConstraints(_234_input)
	_237_ret = _out87
	if !((_237_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(511,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	var _pat_let_tv16 = _247_two
	_ = _pat_let_tv16
	var _pat_let_tv17 = _247_two
	_ = _pat_let_tv17
	var _pat_let_tv18 = _247_two
	_ = _pat_let_tv18
	var _pat_let_tv19 = _247_two
	_ = _pat_let_tv19
	var _pat_let_tv20 = _247_two
	_ = _pat_let_tv20
	var _pat_let_tv21 = _247_two
	_ = _pat_let_tv21
	var _pat_let_tv22 = _247_two
	_ = _pat_let_tv22
	var _pat_let_tv23 = _247_two
	_ = _pat_let_tv23
	var _pat_let_tv24 = _247_two
	_ = _pat_let_tv24
	var _pat_let_tv25 = _247_two
	_ = _pat_let_tv25
	var _pat_let_tv26 = _247_two
	_ = _pat_let_tv26
	_234_input = func(_pat_let166_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_256_dt__update__tmp_h9 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let167_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_257_dt__update_hMyUtf8Bytes_h9 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_256_dt__update__tmp_h9).Dtor_MyString(), (_256_dt__update__tmp_h9).Dtor_NonEmptyString(), (_256_dt__update__tmp_h9).Dtor_StringLessThanOrEqualToTen(), (_256_dt__update__tmp_h9).Dtor_MyBlob(), (_256_dt__update__tmp_h9).Dtor_NonEmptyBlob(), (_256_dt__update__tmp_h9).Dtor_BlobLessThanOrEqualToTen(), (_256_dt__update__tmp_h9).Dtor_MyList(), (_256_dt__update__tmp_h9).Dtor_NonEmptyList(), (_256_dt__update__tmp_h9).Dtor_ListLessThanOrEqualToTen(), (_256_dt__update__tmp_h9).Dtor_MyMap(), (_256_dt__update__tmp_h9).Dtor_NonEmptyMap(), (_256_dt__update__tmp_h9).Dtor_MapLessThanOrEqualToTen(), (_256_dt__update__tmp_h9).Dtor_OneToTen(), (_256_dt__update__tmp_h9).Dtor_myTenToTen(), (_256_dt__update__tmp_h9).Dtor_GreaterThanOne(), (_256_dt__update__tmp_h9).Dtor_LessThanTen(), _257_dt__update_hMyUtf8Bytes_h9, (_256_dt__update__tmp_h9).Dtor_MyListOfUtf8Bytes())
				}(_pat_let167_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceUtf8Bytes(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_pat_let_tv16, _pat_let_tv17), _pat_let_tv18), _pat_let_tv19), _pat_let_tv20), _pat_let_tv21), _pat_let_tv22), _pat_let_tv23), _pat_let_tv24), _pat_let_tv25), _pat_let_tv26))))
		}(_pat_let166_0)
	}(_234_input)
	var _out88 Wrappers.Result
	_ = _out88
	_out88 = (client).GetConstraints(_234_input)
	_237_ret = _out88
	if !((_237_ret).Is_Failure()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(515,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
}
func (_static *CompanionStruct_Default___) TestGetConstraintWithListOfUtf8Bytes(client simpleconstraintsinternaldafnytypes.ISimpleConstraintsClient) {
	var _258_bad _dafny.Sequence
	_ = _258_bad
	_258_bad = Helpers.Companion_Default___.ForceUtf8Bytes(_dafny.SeqOf(uint8(255), uint8(255), uint8(255)))
	var _259_good _dafny.Sequence
	_ = _259_good
	_259_good = Helpers.Companion_Default___.ForceUtf8Bytes(_dafny.SeqOf(uint8(1), uint8(2), uint8(3)))
	var _260_input simpleconstraintsinternaldafnytypes.GetConstraintsInput
	_ = _260_input
	_260_input = Helpers.Companion_Default___.GetValidInput()
	_260_input = func(_pat_let168_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_261_dt__update__tmp_h0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let169_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_262_dt__update_hMyListOfUtf8Bytes_h0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_261_dt__update__tmp_h0).Dtor_MyString(), (_261_dt__update__tmp_h0).Dtor_NonEmptyString(), (_261_dt__update__tmp_h0).Dtor_StringLessThanOrEqualToTen(), (_261_dt__update__tmp_h0).Dtor_MyBlob(), (_261_dt__update__tmp_h0).Dtor_NonEmptyBlob(), (_261_dt__update__tmp_h0).Dtor_BlobLessThanOrEqualToTen(), (_261_dt__update__tmp_h0).Dtor_MyList(), (_261_dt__update__tmp_h0).Dtor_NonEmptyList(), (_261_dt__update__tmp_h0).Dtor_ListLessThanOrEqualToTen(), (_261_dt__update__tmp_h0).Dtor_MyMap(), (_261_dt__update__tmp_h0).Dtor_NonEmptyMap(), (_261_dt__update__tmp_h0).Dtor_MapLessThanOrEqualToTen(), (_261_dt__update__tmp_h0).Dtor_OneToTen(), (_261_dt__update__tmp_h0).Dtor_myTenToTen(), (_261_dt__update__tmp_h0).Dtor_GreaterThanOne(), (_261_dt__update__tmp_h0).Dtor_LessThanTen(), (_261_dt__update__tmp_h0).Dtor_MyUtf8Bytes(), _262_dt__update_hMyListOfUtf8Bytes_h0)
				}(_pat_let169_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceListOfUtf8Bytes(_dafny.SeqOf())))
		}(_pat_let168_0)
	}(_260_input)
	var _263_ret Wrappers.Result
	_ = _263_ret
	var _out89 Wrappers.Result
	_ = _out89
	_out89 = (client).GetConstraints(_260_input)
	_263_ret = _out89
	if !((_263_ret).Is_Failure()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(542,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	var _pat_let_tv27 = _259_good
	_ = _pat_let_tv27
	_260_input = func(_pat_let170_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_264_dt__update__tmp_h1 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let171_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_265_dt__update_hMyListOfUtf8Bytes_h1 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_264_dt__update__tmp_h1).Dtor_MyString(), (_264_dt__update__tmp_h1).Dtor_NonEmptyString(), (_264_dt__update__tmp_h1).Dtor_StringLessThanOrEqualToTen(), (_264_dt__update__tmp_h1).Dtor_MyBlob(), (_264_dt__update__tmp_h1).Dtor_NonEmptyBlob(), (_264_dt__update__tmp_h1).Dtor_BlobLessThanOrEqualToTen(), (_264_dt__update__tmp_h1).Dtor_MyList(), (_264_dt__update__tmp_h1).Dtor_NonEmptyList(), (_264_dt__update__tmp_h1).Dtor_ListLessThanOrEqualToTen(), (_264_dt__update__tmp_h1).Dtor_MyMap(), (_264_dt__update__tmp_h1).Dtor_NonEmptyMap(), (_264_dt__update__tmp_h1).Dtor_MapLessThanOrEqualToTen(), (_264_dt__update__tmp_h1).Dtor_OneToTen(), (_264_dt__update__tmp_h1).Dtor_myTenToTen(), (_264_dt__update__tmp_h1).Dtor_GreaterThanOne(), (_264_dt__update__tmp_h1).Dtor_LessThanTen(), (_264_dt__update__tmp_h1).Dtor_MyUtf8Bytes(), _265_dt__update_hMyListOfUtf8Bytes_h1)
				}(_pat_let171_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceListOfUtf8Bytes(_dafny.SeqOf(_pat_let_tv27))))
		}(_pat_let170_0)
	}(_260_input)
	var _out90 Wrappers.Result
	_ = _out90
	_out90 = (client).GetConstraints(_260_input)
	_263_ret = _out90
	if !((_263_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(546,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	var _pat_let_tv28 = _259_good
	_ = _pat_let_tv28
	var _pat_let_tv29 = _259_good
	_ = _pat_let_tv29
	_260_input = func(_pat_let172_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_266_dt__update__tmp_h2 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let173_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_267_dt__update_hMyListOfUtf8Bytes_h2 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_266_dt__update__tmp_h2).Dtor_MyString(), (_266_dt__update__tmp_h2).Dtor_NonEmptyString(), (_266_dt__update__tmp_h2).Dtor_StringLessThanOrEqualToTen(), (_266_dt__update__tmp_h2).Dtor_MyBlob(), (_266_dt__update__tmp_h2).Dtor_NonEmptyBlob(), (_266_dt__update__tmp_h2).Dtor_BlobLessThanOrEqualToTen(), (_266_dt__update__tmp_h2).Dtor_MyList(), (_266_dt__update__tmp_h2).Dtor_NonEmptyList(), (_266_dt__update__tmp_h2).Dtor_ListLessThanOrEqualToTen(), (_266_dt__update__tmp_h2).Dtor_MyMap(), (_266_dt__update__tmp_h2).Dtor_NonEmptyMap(), (_266_dt__update__tmp_h2).Dtor_MapLessThanOrEqualToTen(), (_266_dt__update__tmp_h2).Dtor_OneToTen(), (_266_dt__update__tmp_h2).Dtor_myTenToTen(), (_266_dt__update__tmp_h2).Dtor_GreaterThanOne(), (_266_dt__update__tmp_h2).Dtor_LessThanTen(), (_266_dt__update__tmp_h2).Dtor_MyUtf8Bytes(), _267_dt__update_hMyListOfUtf8Bytes_h2)
				}(_pat_let173_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceListOfUtf8Bytes(_dafny.SeqOf(_pat_let_tv28, _pat_let_tv29))))
		}(_pat_let172_0)
	}(_260_input)
	var _out91 Wrappers.Result
	_ = _out91
	_out91 = (client).GetConstraints(_260_input)
	_263_ret = _out91
	if !((_263_ret).Is_Success()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(550,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	var _pat_let_tv30 = _259_good
	_ = _pat_let_tv30
	var _pat_let_tv31 = _259_good
	_ = _pat_let_tv31
	var _pat_let_tv32 = _259_good
	_ = _pat_let_tv32
	_260_input = func(_pat_let174_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_268_dt__update__tmp_h3 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let175_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_269_dt__update_hMyListOfUtf8Bytes_h3 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_268_dt__update__tmp_h3).Dtor_MyString(), (_268_dt__update__tmp_h3).Dtor_NonEmptyString(), (_268_dt__update__tmp_h3).Dtor_StringLessThanOrEqualToTen(), (_268_dt__update__tmp_h3).Dtor_MyBlob(), (_268_dt__update__tmp_h3).Dtor_NonEmptyBlob(), (_268_dt__update__tmp_h3).Dtor_BlobLessThanOrEqualToTen(), (_268_dt__update__tmp_h3).Dtor_MyList(), (_268_dt__update__tmp_h3).Dtor_NonEmptyList(), (_268_dt__update__tmp_h3).Dtor_ListLessThanOrEqualToTen(), (_268_dt__update__tmp_h3).Dtor_MyMap(), (_268_dt__update__tmp_h3).Dtor_NonEmptyMap(), (_268_dt__update__tmp_h3).Dtor_MapLessThanOrEqualToTen(), (_268_dt__update__tmp_h3).Dtor_OneToTen(), (_268_dt__update__tmp_h3).Dtor_myTenToTen(), (_268_dt__update__tmp_h3).Dtor_GreaterThanOne(), (_268_dt__update__tmp_h3).Dtor_LessThanTen(), (_268_dt__update__tmp_h3).Dtor_MyUtf8Bytes(), _269_dt__update_hMyListOfUtf8Bytes_h3)
				}(_pat_let175_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceListOfUtf8Bytes(_dafny.SeqOf(_pat_let_tv30, _pat_let_tv31, _pat_let_tv32))))
		}(_pat_let174_0)
	}(_260_input)
	var _out92 Wrappers.Result
	_ = _out92
	_out92 = (client).GetConstraints(_260_input)
	_263_ret = _out92
	if !((_263_ret).Is_Failure()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(554,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	var _pat_let_tv33 = _258_bad
	_ = _pat_let_tv33
	_260_input = func(_pat_let176_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_270_dt__update__tmp_h4 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let177_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_271_dt__update_hMyListOfUtf8Bytes_h4 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_270_dt__update__tmp_h4).Dtor_MyString(), (_270_dt__update__tmp_h4).Dtor_NonEmptyString(), (_270_dt__update__tmp_h4).Dtor_StringLessThanOrEqualToTen(), (_270_dt__update__tmp_h4).Dtor_MyBlob(), (_270_dt__update__tmp_h4).Dtor_NonEmptyBlob(), (_270_dt__update__tmp_h4).Dtor_BlobLessThanOrEqualToTen(), (_270_dt__update__tmp_h4).Dtor_MyList(), (_270_dt__update__tmp_h4).Dtor_NonEmptyList(), (_270_dt__update__tmp_h4).Dtor_ListLessThanOrEqualToTen(), (_270_dt__update__tmp_h4).Dtor_MyMap(), (_270_dt__update__tmp_h4).Dtor_NonEmptyMap(), (_270_dt__update__tmp_h4).Dtor_MapLessThanOrEqualToTen(), (_270_dt__update__tmp_h4).Dtor_OneToTen(), (_270_dt__update__tmp_h4).Dtor_myTenToTen(), (_270_dt__update__tmp_h4).Dtor_GreaterThanOne(), (_270_dt__update__tmp_h4).Dtor_LessThanTen(), (_270_dt__update__tmp_h4).Dtor_MyUtf8Bytes(), _271_dt__update_hMyListOfUtf8Bytes_h4)
				}(_pat_let177_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceListOfUtf8Bytes(_dafny.SeqOf(_pat_let_tv33))))
		}(_pat_let176_0)
	}(_260_input)
	var _out93 Wrappers.Result
	_ = _out93
	_out93 = (client).GetConstraints(_260_input)
	_263_ret = _out93
	if !((_263_ret).Is_Failure()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(558,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	var _pat_let_tv34 = _258_bad
	_ = _pat_let_tv34
	var _pat_let_tv35 = _259_good
	_ = _pat_let_tv35
	_260_input = func(_pat_let178_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_272_dt__update__tmp_h5 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let179_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_273_dt__update_hMyListOfUtf8Bytes_h5 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_272_dt__update__tmp_h5).Dtor_MyString(), (_272_dt__update__tmp_h5).Dtor_NonEmptyString(), (_272_dt__update__tmp_h5).Dtor_StringLessThanOrEqualToTen(), (_272_dt__update__tmp_h5).Dtor_MyBlob(), (_272_dt__update__tmp_h5).Dtor_NonEmptyBlob(), (_272_dt__update__tmp_h5).Dtor_BlobLessThanOrEqualToTen(), (_272_dt__update__tmp_h5).Dtor_MyList(), (_272_dt__update__tmp_h5).Dtor_NonEmptyList(), (_272_dt__update__tmp_h5).Dtor_ListLessThanOrEqualToTen(), (_272_dt__update__tmp_h5).Dtor_MyMap(), (_272_dt__update__tmp_h5).Dtor_NonEmptyMap(), (_272_dt__update__tmp_h5).Dtor_MapLessThanOrEqualToTen(), (_272_dt__update__tmp_h5).Dtor_OneToTen(), (_272_dt__update__tmp_h5).Dtor_myTenToTen(), (_272_dt__update__tmp_h5).Dtor_GreaterThanOne(), (_272_dt__update__tmp_h5).Dtor_LessThanTen(), (_272_dt__update__tmp_h5).Dtor_MyUtf8Bytes(), _273_dt__update_hMyListOfUtf8Bytes_h5)
				}(_pat_let179_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceListOfUtf8Bytes(_dafny.SeqOf(_pat_let_tv34, _pat_let_tv35))))
		}(_pat_let178_0)
	}(_260_input)
	var _out94 Wrappers.Result
	_ = _out94
	_out94 = (client).GetConstraints(_260_input)
	_263_ret = _out94
	if !((_263_ret).Is_Failure()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(562,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	var _pat_let_tv36 = _259_good
	_ = _pat_let_tv36
	var _pat_let_tv37 = _258_bad
	_ = _pat_let_tv37
	_260_input = func(_pat_let180_0 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
		return func(_274_dt__update__tmp_h6 simpleconstraintsinternaldafnytypes.GetConstraintsInput) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
			return func(_pat_let181_0 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
				return func(_275_dt__update_hMyListOfUtf8Bytes_h6 Wrappers.Option) simpleconstraintsinternaldafnytypes.GetConstraintsInput {
					return simpleconstraintsinternaldafnytypes.Companion_GetConstraintsInput_.Create_GetConstraintsInput_((_274_dt__update__tmp_h6).Dtor_MyString(), (_274_dt__update__tmp_h6).Dtor_NonEmptyString(), (_274_dt__update__tmp_h6).Dtor_StringLessThanOrEqualToTen(), (_274_dt__update__tmp_h6).Dtor_MyBlob(), (_274_dt__update__tmp_h6).Dtor_NonEmptyBlob(), (_274_dt__update__tmp_h6).Dtor_BlobLessThanOrEqualToTen(), (_274_dt__update__tmp_h6).Dtor_MyList(), (_274_dt__update__tmp_h6).Dtor_NonEmptyList(), (_274_dt__update__tmp_h6).Dtor_ListLessThanOrEqualToTen(), (_274_dt__update__tmp_h6).Dtor_MyMap(), (_274_dt__update__tmp_h6).Dtor_NonEmptyMap(), (_274_dt__update__tmp_h6).Dtor_MapLessThanOrEqualToTen(), (_274_dt__update__tmp_h6).Dtor_OneToTen(), (_274_dt__update__tmp_h6).Dtor_myTenToTen(), (_274_dt__update__tmp_h6).Dtor_GreaterThanOne(), (_274_dt__update__tmp_h6).Dtor_LessThanTen(), (_274_dt__update__tmp_h6).Dtor_MyUtf8Bytes(), _275_dt__update_hMyListOfUtf8Bytes_h6)
				}(_pat_let181_0)
			}(Wrappers.Companion_Option_.Create_Some_(Helpers.Companion_Default___.ForceListOfUtf8Bytes(_dafny.SeqOf(_pat_let_tv36, _pat_let_tv37))))
		}(_pat_let180_0)
	}(_260_input)
	var _out95 Wrappers.Result
	_ = _out95
	_out95 = (client).GetConstraints(_260_input)
	_263_ret = _out95
	if !((_263_ret).Is_Failure()) {
		panic("test/WrappedSimpleConstraintsTest.dfy(566,4): " + (_dafny.SeqOfString("expectation violation")).String())
	}
}

// End of class Default__
