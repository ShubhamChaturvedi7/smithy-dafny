// Dafny program the_program compiled into Go
package main

import (
	StandardLibrary "StandardLibrary"
	StandardLibraryInterop "StandardLibraryInterop"
	StandardLibrary_UInt "StandardLibrary_UInt"
	Wrappers "Wrappers"
	os "os"

	_System "github.com/dafny-lang/DafnyRuntimeGo/System_"
	_dafny "github.com/dafny-lang/DafnyRuntimeGo/dafny"
	TestUTF8 "github.com/dafny-lang/DafnyStandardLibGo/TestUTF8"
)

var _ = os.Args
var _ _dafny.Dummy__
var _ _System.Dummy__
var _ Wrappers.Dummy__
var _ StandardLibrary_UInt.Dummy__
var _ StandardLibrary.Dummy__
var _ StandardLibraryInterop.Dummy__
var _ TestUTF8.Dummy__

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
	var _57_success bool
	_ = _57_success
	_57_success = true
	_dafny.Print((_dafny.SeqOfString("TestUTF8.TestEncodeHappyCase: ")).SetString())
	func() {
		defer func() {
			if r := recover(); r != nil {
				var _58_haltMessage = _dafny.SeqOfString(r.(string))
				{
					_dafny.Print((_dafny.SeqOfString("FAILED\n	")).SetString())
					_dafny.Print((_58_haltMessage).SetString())
					_dafny.Print((_dafny.SeqOfString("\n")).SetString())
					_57_success = false
				}
			}
		}()
		{
			TestUTF8.Companion_Default___.TestEncodeHappyCase()
			{
				_dafny.Print((_dafny.SeqOfString("PASSED\n")).SetString())
			}
		}
	}()
	_dafny.Print((_dafny.SeqOfString("TestUTF8.TestEncodeInvalidUnicode: ")).SetString())
	func() {
		defer func() {
			if r := recover(); r != nil {
				var _59_haltMessage = _dafny.SeqOfString(r.(string))
				{
					_dafny.Print((_dafny.SeqOfString("FAILED\n	")).SetString())
					_dafny.Print((_59_haltMessage).SetString())
					_dafny.Print((_dafny.SeqOfString("\n")).SetString())
					_57_success = false
				}
			}
		}()
		{
			TestUTF8.Companion_Default___.TestEncodeInvalidUnicode()
			{
				_dafny.Print((_dafny.SeqOfString("PASSED\n")).SetString())
			}
		}
	}()
	_dafny.Print((_dafny.SeqOfString("TestUTF8.TestDecodeHappyCase: ")).SetString())
	func() {
		defer func() {
			if r := recover(); r != nil {
				var _60_haltMessage = _dafny.SeqOfString(r.(string))
				{
					_dafny.Print((_dafny.SeqOfString("FAILED\n	")).SetString())
					_dafny.Print((_60_haltMessage).SetString())
					_dafny.Print((_dafny.SeqOfString("\n")).SetString())
					_57_success = false
				}
			}
		}()
		{
			TestUTF8.Companion_Default___.TestDecodeHappyCase()
			{
				_dafny.Print((_dafny.SeqOfString("PASSED\n")).SetString())
			}
		}
	}()
	_dafny.Print((_dafny.SeqOfString("TestUTF8.TestDecodeInvalidUnicode: ")).SetString())
	func() {
		defer func() {
			if r := recover(); r != nil {
				var _61_haltMessage = _dafny.SeqOfString(r.(string))
				{
					_dafny.Print((_dafny.SeqOfString("FAILED\n	")).SetString())
					_dafny.Print((_61_haltMessage).SetString())
					_dafny.Print((_dafny.SeqOfString("\n")).SetString())
					_57_success = false
				}
			}
		}()
		{
			TestUTF8.Companion_Default___.TestDecodeInvalidUnicode()
			{
				_dafny.Print((_dafny.SeqOfString("PASSED\n")).SetString())
			}
		}
	}()
	_dafny.Print((_dafny.SeqOfString("TestUTF8.Test1Byte: ")).SetString())
	func() {
		defer func() {
			if r := recover(); r != nil {
				var _62_haltMessage = _dafny.SeqOfString(r.(string))
				{
					_dafny.Print((_dafny.SeqOfString("FAILED\n	")).SetString())
					_dafny.Print((_62_haltMessage).SetString())
					_dafny.Print((_dafny.SeqOfString("\n")).SetString())
					_57_success = false
				}
			}
		}()
		{
			TestUTF8.Companion_Default___.Test1Byte()
			{
				_dafny.Print((_dafny.SeqOfString("PASSED\n")).SetString())
			}
		}
	}()
	_dafny.Print((_dafny.SeqOfString("TestUTF8.Test2Bytes: ")).SetString())
	func() {
		defer func() {
			if r := recover(); r != nil {
				var _63_haltMessage = _dafny.SeqOfString(r.(string))
				{
					_dafny.Print((_dafny.SeqOfString("FAILED\n	")).SetString())
					_dafny.Print((_63_haltMessage).SetString())
					_dafny.Print((_dafny.SeqOfString("\n")).SetString())
					_57_success = false
				}
			}
		}()
		{
			TestUTF8.Companion_Default___.Test2Bytes()
			{
				_dafny.Print((_dafny.SeqOfString("PASSED\n")).SetString())
			}
		}
	}()
	_dafny.Print((_dafny.SeqOfString("TestUTF8.Test3Bytes: ")).SetString())
	func() {
		defer func() {
			if r := recover(); r != nil {
				var _64_haltMessage = _dafny.SeqOfString(r.(string))
				{
					_dafny.Print((_dafny.SeqOfString("FAILED\n	")).SetString())
					_dafny.Print((_64_haltMessage).SetString())
					_dafny.Print((_dafny.SeqOfString("\n")).SetString())
					_57_success = false
				}
			}
		}()
		{
			TestUTF8.Companion_Default___.Test3Bytes()
			{
				_dafny.Print((_dafny.SeqOfString("PASSED\n")).SetString())
			}
		}
	}()
	_dafny.Print((_dafny.SeqOfString("TestUTF8.Test4Bytes: ")).SetString())
	func() {
		defer func() {
			if r := recover(); r != nil {
				var _65_haltMessage = _dafny.SeqOfString(r.(string))
				{
					_dafny.Print((_dafny.SeqOfString("FAILED\n	")).SetString())
					_dafny.Print((_65_haltMessage).SetString())
					_dafny.Print((_dafny.SeqOfString("\n")).SetString())
					_57_success = false
				}
			}
		}()
		{
			TestUTF8.Companion_Default___.Test4Bytes()
			{
				_dafny.Print((_dafny.SeqOfString("PASSED\n")).SetString())
			}
		}
	}()
	if !(_57_success) {
		panic("<stdin>(1,0): " + (_dafny.SeqOfString("Test failures occurred: see above.\n")).String())
	}
}

// End of class Default__
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Test____Main____(_dafny.FromMainArguments(os.Args))
}
