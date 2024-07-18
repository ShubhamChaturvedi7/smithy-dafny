// Package SimpleConstraintsImpl
// Dafny module SimpleConstraintsImpl compiled into Go

package SimpleConstraintsImpl

import (
	os "os"

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
	return "SimpleConstraintsImpl.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) GetConstraints(config Config, input simpleconstraintsinternaldafnytypes.GetConstraintsInput) Wrappers.Result {
	var output Wrappers.Result = Wrappers.Companion_Result_.Default(simpleconstraintsinternaldafnytypes.Companion_GetConstraintsOutput_.Default())
	_ = output
	var _19_res simpleconstraintsinternaldafnytypes.GetConstraintsOutput
	_ = _19_res
	_19_res = simpleconstraintsinternaldafnytypes.Companion_GetConstraintsOutput_.Create_GetConstraintsOutput_((input).Dtor_MyString(), (input).Dtor_NonEmptyString(), (input).Dtor_StringLessThanOrEqualToTen(), (input).Dtor_MyBlob(), (input).Dtor_NonEmptyBlob(), (input).Dtor_BlobLessThanOrEqualToTen(), (input).Dtor_MyList(), (input).Dtor_NonEmptyList(), (input).Dtor_ListLessThanOrEqualToTen(), (input).Dtor_MyMap(), (input).Dtor_NonEmptyMap(), (input).Dtor_MapLessThanOrEqualToTen(), (input).Dtor_OneToTen(), Wrappers.Companion_Option_.Create_None_(), (input).Dtor_GreaterThanOne(), (input).Dtor_LessThanTen(), (input).Dtor_MyUtf8Bytes(), (input).Dtor_MyListOfUtf8Bytes())
	output = Wrappers.Companion_Result_.Create_Success_(_19_res)
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
			return "SimpleConstraintsImpl.Config.Config"
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
	return "SimpleConstraintsImpl.Config"
}
func (_this Config) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = Config{}

// End of datatype Config
