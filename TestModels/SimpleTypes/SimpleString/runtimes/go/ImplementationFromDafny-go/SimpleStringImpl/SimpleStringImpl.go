// Package SimpleStringImpl
// Dafny module SimpleStringImpl compiled into Go

package SimpleStringImpl

import (
	os "os"

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
	return "SimpleStringImpl.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) GetString(config Config, input simpletypessmithystringinternaldafnytypes.GetStringInput) Wrappers.Result {
	var output Wrappers.Result = Wrappers.Companion_Result_.Default(simpletypessmithystringinternaldafnytypes.Companion_GetStringOutput_.Default())
	_ = output
	var _1_res simpletypessmithystringinternaldafnytypes.GetStringOutput
	_ = _1_res
	_1_res = simpletypessmithystringinternaldafnytypes.Companion_GetStringOutput_.Create_GetStringOutput_((input).Dtor_value())
	output = Wrappers.Companion_Result_.Create_Success_(_1_res)
	return output
	return output
}
func (_static *CompanionStruct_Default___) GetStringKnownValue(config Config, input simpletypessmithystringinternaldafnytypes.GetStringInput) Wrappers.Result {
	var output Wrappers.Result = Wrappers.Companion_Result_.Default(simpletypessmithystringinternaldafnytypes.Companion_GetStringOutput_.Default())
	_ = output
	if !(((input).Dtor_value()).Is_Some()) {
		panic("src/SimpleStringImpl.dfy(28,8): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	if !(_dafny.Companion_Sequence_.Equal(((input).Dtor_value()).Dtor_value().(_dafny.Sequence), _dafny.SeqOfString("TEST_SIMPLE_STRING_KNOWN_VALUE"))) {
		panic("src/SimpleStringImpl.dfy(29,8): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	var _2_res simpletypessmithystringinternaldafnytypes.GetStringOutput
	_ = _2_res
	_2_res = simpletypessmithystringinternaldafnytypes.Companion_GetStringOutput_.Create_GetStringOutput_((input).Dtor_value())
	output = Wrappers.Companion_Result_.Create_Success_(_2_res)
	return output
	return output
}
func (_static *CompanionStruct_Default___) GetStringUTF8(config Config, input simpletypessmithystringinternaldafnytypes.GetStringInput) Wrappers.Result {
	var output Wrappers.Result = Wrappers.Companion_Result_.Default(simpletypessmithystringinternaldafnytypes.Companion_GetStringOutput_.Default())
	_ = output
	if !(((input).Dtor_value()).Is_Some()) {
		panic("src/SimpleStringImpl.dfy(35,8): " + (_dafny.SeqOfString("expectation violation")).String())
	}
	var _3_res simpletypessmithystringinternaldafnytypes.GetStringOutput
	_ = _3_res
	_3_res = simpletypessmithystringinternaldafnytypes.Companion_GetStringOutput_.Create_GetStringOutput_((input).Dtor_value())
	output = Wrappers.Companion_Result_.Create_Success_(_3_res)
	return output
	return output
}

// End of class Default__

// Definition of datatype Config
type Config struct {
	Data_Config_
}

func (_this Config) Get_() Data_Config_ {
	return _this.Data_Config_
}

type Data_Config_ interface {
	isConfig()
}

type CompanionStruct_Config_ struct {
}

var Companion_Config_ = CompanionStruct_Config_{}

type Config_Config struct {
}

func (Config_Config) isConfig() {}

func (CompanionStruct_Config_) Create_Config_() Config {
	return Config{Config_Config{}}
}

func (_this Config) Is_Config() bool {
	_, ok := _this.Get_().(Config_Config)
	return ok
}

func (CompanionStruct_Config_) Default() Config {
	return Companion_Config_.Create_Config_()
}

func (_ CompanionStruct_Config_) AllSingletonConstructors() _dafny.Iterator {
	i := -1
	return func() (interface{}, bool) {
		i++
		switch i {
		case 0:
			return Companion_Config_.Create_Config_(), true
		default:
			return Config{}, false
		}
	}
}

func (_this Config) String() string {
	switch _this.Get_().(type) {
	case nil:
		return "null"
	case Config_Config:
		{
			return "SimpleStringImpl.Config.Config"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this Config) Equals(other Config) bool {
	switch _this.Get_().(type) {
	case Config_Config:
		{
			_, ok := other.Get_().(Config_Config)
			return ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this Config) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(Config)
	return ok && _this.Equals(typed)
}

func Type_Config_() _dafny.TypeDescriptor {
	return type_Config_{}
}

type type_Config_ struct {
}

func (_this type_Config_) Default() interface{} {
	return Companion_Config_.Default()
}

func (_this type_Config_) String() string {
	return "SimpleStringImpl.Config"
}
func (_this Config) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = Config{}

// End of datatype Config
