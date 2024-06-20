// Dafny program the_program compiled into Go
package main

import (
	os "os"

	SimpleUnionImpl "github.com/Smithy-dafny/TestModels/Union/SimpleUnionImpl"
	SimpleUnionImplTest "github.com/Smithy-dafny/TestModels/Union/SimpleUnionImplTest"
	WrappedSimpleUnionTest "github.com/Smithy-dafny/TestModels/Union/WrappedSimpleUnionTest"
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
var _ WrappedSimpleUnionTest.Dummy__

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
	return "_module.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) Test____Main____(__noArgsParameter _dafny.Sequence) {
	var _14_success bool
	_ = _14_success
	_14_success = true
	_dafny.Print((_dafny.SeqOfString("SimpleUnionImplTest.TestUnion: ")).SetString())
	func() {
		defer func() {
			if r := recover(); r != nil {
				var _15_haltMessage = _dafny.SeqOfString(r.(string))
				{
					_dafny.Print((_dafny.SeqOfString("FAILED\n	")).SetString())
					_dafny.Print((_15_haltMessage).SetString())
					_dafny.Print((_dafny.SeqOfString("\n")).SetString())
					_14_success = false
				}
			}
		}()
		{
			SimpleUnionImplTest.Companion_Default___.TestUnion()
			{
				_dafny.Print((_dafny.SeqOfString("PASSED\n")).SetString())
			}
		}
	}()
	_dafny.Print((_dafny.SeqOfString("WrappedSimpleUnionTest.TestUnion: ")).SetString())
	func() {
		defer func() {
			if r := recover(); r != nil {
				var _16_haltMessage = _dafny.SeqOfString(r.(string))
				{
					_dafny.Print((_dafny.SeqOfString("FAILED\n	")).SetString())
					_dafny.Print((_16_haltMessage).SetString())
					_dafny.Print((_dafny.SeqOfString("\n")).SetString())
					_14_success = false
				}
			}
		}()
		{
			WrappedSimpleUnionTest.Companion_Default___.TestUnion()
			{
				_dafny.Print((_dafny.SeqOfString("PASSED\n")).SetString())
			}
		}
	}()
	if !(_14_success) {
		panic("<stdin>(1,0): " + (_dafny.SeqOfString("Test failures occurred: see above.\n")).String())
	}
}

// End of class Default__
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Test____Main____(_dafny.FromMainArguments(os.Args))
}
