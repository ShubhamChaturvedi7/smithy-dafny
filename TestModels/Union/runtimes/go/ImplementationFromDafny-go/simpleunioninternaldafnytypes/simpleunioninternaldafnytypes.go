// Package simpleunioninternaldafnytypes
// Dafny module simpleunioninternaldafnytypes compiled into Go

package simpleunioninternaldafnytypes

import (
	os "os"

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

// Definition of datatype DafnyCallEvent
type DafnyCallEvent struct {
	Data_DafnyCallEvent_
}

func (_this DafnyCallEvent) Get_() Data_DafnyCallEvent_ {
	return _this.Data_DafnyCallEvent_
}

type Data_DafnyCallEvent_ interface {
	isDafnyCallEvent()
}

type CompanionStruct_DafnyCallEvent_ struct {
}

var Companion_DafnyCallEvent_ = CompanionStruct_DafnyCallEvent_{}

type DafnyCallEvent_DafnyCallEvent struct {
	Input  interface{}
	Output interface{}
}

func (DafnyCallEvent_DafnyCallEvent) isDafnyCallEvent() {}

func (CompanionStruct_DafnyCallEvent_) Create_DafnyCallEvent_(Input interface{}, Output interface{}) DafnyCallEvent {
	return DafnyCallEvent{DafnyCallEvent_DafnyCallEvent{Input, Output}}
}

func (_this DafnyCallEvent) Is_DafnyCallEvent() bool {
	_, ok := _this.Get_().(DafnyCallEvent_DafnyCallEvent)
	return ok
}

func (CompanionStruct_DafnyCallEvent_) Default(_default_I interface{}, _default_O interface{}) DafnyCallEvent {
	return Companion_DafnyCallEvent_.Create_DafnyCallEvent_(_default_I, _default_O)
}

func (_this DafnyCallEvent) Dtor_input() interface{} {
	return _this.Get_().(DafnyCallEvent_DafnyCallEvent).Input
}

func (_this DafnyCallEvent) Dtor_output() interface{} {
	return _this.Get_().(DafnyCallEvent_DafnyCallEvent).Output
}

func (_this DafnyCallEvent) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case DafnyCallEvent_DafnyCallEvent:
		{
			return "SimpleUnionTypes.DafnyCallEvent.DafnyCallEvent" + "(" + _dafny.String(data.Input) + ", " + _dafny.String(data.Output) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this DafnyCallEvent) Equals(other DafnyCallEvent) bool {
	switch data1 := _this.Get_().(type) {
	case DafnyCallEvent_DafnyCallEvent:
		{
			data2, ok := other.Get_().(DafnyCallEvent_DafnyCallEvent)
			return ok && _dafny.AreEqual(data1.Input, data2.Input) && _dafny.AreEqual(data1.Output, data2.Output)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this DafnyCallEvent) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(DafnyCallEvent)
	return ok && _this.Equals(typed)
}

func Type_DafnyCallEvent_(Type_I_ _dafny.TypeDescriptor, Type_O_ _dafny.TypeDescriptor) _dafny.TypeDescriptor {
	return type_DafnyCallEvent_{Type_I_, Type_O_}
}

type type_DafnyCallEvent_ struct {
	Type_I_ _dafny.TypeDescriptor
	Type_O_ _dafny.TypeDescriptor
}

func (_this type_DafnyCallEvent_) Default() interface{} {
	Type_I_ := _this.Type_I_
	_ = Type_I_
	Type_O_ := _this.Type_O_
	_ = Type_O_
	return Companion_DafnyCallEvent_.Default(Type_I_.Default(), Type_O_.Default())
}

func (_this type_DafnyCallEvent_) String() string {
	return "simpleunioninternaldafnytypes.DafnyCallEvent"
}
func (_this DafnyCallEvent) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = DafnyCallEvent{}

// End of datatype DafnyCallEvent

// Definition of datatype GetKnownValueUnionInput
type GetKnownValueUnionInput struct {
	Data_GetKnownValueUnionInput_
}

func (_this GetKnownValueUnionInput) Get_() Data_GetKnownValueUnionInput_ {
	return _this.Data_GetKnownValueUnionInput_
}

type Data_GetKnownValueUnionInput_ interface {
	isGetKnownValueUnionInput()
}

type CompanionStruct_GetKnownValueUnionInput_ struct {
}

var Companion_GetKnownValueUnionInput_ = CompanionStruct_GetKnownValueUnionInput_{}

type GetKnownValueUnionInput_GetKnownValueUnionInput struct {
	Union Wrappers.Option
}

func (GetKnownValueUnionInput_GetKnownValueUnionInput) isGetKnownValueUnionInput() {}

func (CompanionStruct_GetKnownValueUnionInput_) Create_GetKnownValueUnionInput_(Union Wrappers.Option) GetKnownValueUnionInput {
	return GetKnownValueUnionInput{GetKnownValueUnionInput_GetKnownValueUnionInput{Union}}
}

func (_this GetKnownValueUnionInput) Is_GetKnownValueUnionInput() bool {
	_, ok := _this.Get_().(GetKnownValueUnionInput_GetKnownValueUnionInput)
	return ok
}

func (CompanionStruct_GetKnownValueUnionInput_) Default() GetKnownValueUnionInput {
	return Companion_GetKnownValueUnionInput_.Create_GetKnownValueUnionInput_(Wrappers.Companion_Option_.Default())
}

func (_this GetKnownValueUnionInput) Dtor_union() Wrappers.Option {
	return _this.Get_().(GetKnownValueUnionInput_GetKnownValueUnionInput).Union
}

func (_this GetKnownValueUnionInput) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case GetKnownValueUnionInput_GetKnownValueUnionInput:
		{
			return "SimpleUnionTypes.GetKnownValueUnionInput.GetKnownValueUnionInput" + "(" + _dafny.String(data.Union) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this GetKnownValueUnionInput) Equals(other GetKnownValueUnionInput) bool {
	switch data1 := _this.Get_().(type) {
	case GetKnownValueUnionInput_GetKnownValueUnionInput:
		{
			data2, ok := other.Get_().(GetKnownValueUnionInput_GetKnownValueUnionInput)
			return ok && data1.Union.Equals(data2.Union)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this GetKnownValueUnionInput) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(GetKnownValueUnionInput)
	return ok && _this.Equals(typed)
}

func Type_GetKnownValueUnionInput_() _dafny.TypeDescriptor {
	return type_GetKnownValueUnionInput_{}
}

type type_GetKnownValueUnionInput_ struct {
}

func (_this type_GetKnownValueUnionInput_) Default() interface{} {
	return Companion_GetKnownValueUnionInput_.Default()
}

func (_this type_GetKnownValueUnionInput_) String() string {
	return "simpleunioninternaldafnytypes.GetKnownValueUnionInput"
}
func (_this GetKnownValueUnionInput) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = GetKnownValueUnionInput{}

// End of datatype GetKnownValueUnionInput

// Definition of datatype GetKnownValueUnionOutput
type GetKnownValueUnionOutput struct {
	Data_GetKnownValueUnionOutput_
}

func (_this GetKnownValueUnionOutput) Get_() Data_GetKnownValueUnionOutput_ {
	return _this.Data_GetKnownValueUnionOutput_
}

type Data_GetKnownValueUnionOutput_ interface {
	isGetKnownValueUnionOutput()
}

type CompanionStruct_GetKnownValueUnionOutput_ struct {
}

var Companion_GetKnownValueUnionOutput_ = CompanionStruct_GetKnownValueUnionOutput_{}

type GetKnownValueUnionOutput_GetKnownValueUnionOutput struct {
	Union Wrappers.Option
}

func (GetKnownValueUnionOutput_GetKnownValueUnionOutput) isGetKnownValueUnionOutput() {}

func (CompanionStruct_GetKnownValueUnionOutput_) Create_GetKnownValueUnionOutput_(Union Wrappers.Option) GetKnownValueUnionOutput {
	return GetKnownValueUnionOutput{GetKnownValueUnionOutput_GetKnownValueUnionOutput{Union}}
}

func (_this GetKnownValueUnionOutput) Is_GetKnownValueUnionOutput() bool {
	_, ok := _this.Get_().(GetKnownValueUnionOutput_GetKnownValueUnionOutput)
	return ok
}

func (CompanionStruct_GetKnownValueUnionOutput_) Default() GetKnownValueUnionOutput {
	return Companion_GetKnownValueUnionOutput_.Create_GetKnownValueUnionOutput_(Wrappers.Companion_Option_.Default())
}

func (_this GetKnownValueUnionOutput) Dtor_union() Wrappers.Option {
	return _this.Get_().(GetKnownValueUnionOutput_GetKnownValueUnionOutput).Union
}

func (_this GetKnownValueUnionOutput) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case GetKnownValueUnionOutput_GetKnownValueUnionOutput:
		{
			return "SimpleUnionTypes.GetKnownValueUnionOutput.GetKnownValueUnionOutput" + "(" + _dafny.String(data.Union) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this GetKnownValueUnionOutput) Equals(other GetKnownValueUnionOutput) bool {
	switch data1 := _this.Get_().(type) {
	case GetKnownValueUnionOutput_GetKnownValueUnionOutput:
		{
			data2, ok := other.Get_().(GetKnownValueUnionOutput_GetKnownValueUnionOutput)
			return ok && data1.Union.Equals(data2.Union)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this GetKnownValueUnionOutput) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(GetKnownValueUnionOutput)
	return ok && _this.Equals(typed)
}

func Type_GetKnownValueUnionOutput_() _dafny.TypeDescriptor {
	return type_GetKnownValueUnionOutput_{}
}

type type_GetKnownValueUnionOutput_ struct {
}

func (_this type_GetKnownValueUnionOutput_) Default() interface{} {
	return Companion_GetKnownValueUnionOutput_.Default()
}

func (_this type_GetKnownValueUnionOutput_) String() string {
	return "simpleunioninternaldafnytypes.GetKnownValueUnionOutput"
}
func (_this GetKnownValueUnionOutput) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = GetKnownValueUnionOutput{}

// End of datatype GetKnownValueUnionOutput

// Definition of datatype GetUnionInput
type GetUnionInput struct {
	Data_GetUnionInput_
}

func (_this GetUnionInput) Get_() Data_GetUnionInput_ {
	return _this.Data_GetUnionInput_
}

type Data_GetUnionInput_ interface {
	isGetUnionInput()
}

type CompanionStruct_GetUnionInput_ struct {
}

var Companion_GetUnionInput_ = CompanionStruct_GetUnionInput_{}

type GetUnionInput_GetUnionInput struct {
	Union Wrappers.Option
}

func (GetUnionInput_GetUnionInput) isGetUnionInput() {}

func (CompanionStruct_GetUnionInput_) Create_GetUnionInput_(Union Wrappers.Option) GetUnionInput {
	return GetUnionInput{GetUnionInput_GetUnionInput{Union}}
}

func (_this GetUnionInput) Is_GetUnionInput() bool {
	_, ok := _this.Get_().(GetUnionInput_GetUnionInput)
	return ok
}

func (CompanionStruct_GetUnionInput_) Default() GetUnionInput {
	return Companion_GetUnionInput_.Create_GetUnionInput_(Wrappers.Companion_Option_.Default())
}

func (_this GetUnionInput) Dtor_union() Wrappers.Option {
	return _this.Get_().(GetUnionInput_GetUnionInput).Union
}

func (_this GetUnionInput) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case GetUnionInput_GetUnionInput:
		{
			return "SimpleUnionTypes.GetUnionInput.GetUnionInput" + "(" + _dafny.String(data.Union) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this GetUnionInput) Equals(other GetUnionInput) bool {
	switch data1 := _this.Get_().(type) {
	case GetUnionInput_GetUnionInput:
		{
			data2, ok := other.Get_().(GetUnionInput_GetUnionInput)
			return ok && data1.Union.Equals(data2.Union)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this GetUnionInput) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(GetUnionInput)
	return ok && _this.Equals(typed)
}

func Type_GetUnionInput_() _dafny.TypeDescriptor {
	return type_GetUnionInput_{}
}

type type_GetUnionInput_ struct {
}

func (_this type_GetUnionInput_) Default() interface{} {
	return Companion_GetUnionInput_.Default()
}

func (_this type_GetUnionInput_) String() string {
	return "simpleunioninternaldafnytypes.GetUnionInput"
}
func (_this GetUnionInput) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = GetUnionInput{}

// End of datatype GetUnionInput

// Definition of datatype GetUnionOutput
type GetUnionOutput struct {
	Data_GetUnionOutput_
}

func (_this GetUnionOutput) Get_() Data_GetUnionOutput_ {
	return _this.Data_GetUnionOutput_
}

type Data_GetUnionOutput_ interface {
	isGetUnionOutput()
}

type CompanionStruct_GetUnionOutput_ struct {
}

var Companion_GetUnionOutput_ = CompanionStruct_GetUnionOutput_{}

type GetUnionOutput_GetUnionOutput struct {
	Union Wrappers.Option
}

func (GetUnionOutput_GetUnionOutput) isGetUnionOutput() {}

func (CompanionStruct_GetUnionOutput_) Create_GetUnionOutput_(Union Wrappers.Option) GetUnionOutput {
	return GetUnionOutput{GetUnionOutput_GetUnionOutput{Union}}
}

func (_this GetUnionOutput) Is_GetUnionOutput() bool {
	_, ok := _this.Get_().(GetUnionOutput_GetUnionOutput)
	return ok
}

func (CompanionStruct_GetUnionOutput_) Default() GetUnionOutput {
	return Companion_GetUnionOutput_.Create_GetUnionOutput_(Wrappers.Companion_Option_.Default())
}

func (_this GetUnionOutput) Dtor_union() Wrappers.Option {
	return _this.Get_().(GetUnionOutput_GetUnionOutput).Union
}

func (_this GetUnionOutput) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case GetUnionOutput_GetUnionOutput:
		{
			return "SimpleUnionTypes.GetUnionOutput.GetUnionOutput" + "(" + _dafny.String(data.Union) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this GetUnionOutput) Equals(other GetUnionOutput) bool {
	switch data1 := _this.Get_().(type) {
	case GetUnionOutput_GetUnionOutput:
		{
			data2, ok := other.Get_().(GetUnionOutput_GetUnionOutput)
			return ok && data1.Union.Equals(data2.Union)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this GetUnionOutput) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(GetUnionOutput)
	return ok && _this.Equals(typed)
}

func Type_GetUnionOutput_() _dafny.TypeDescriptor {
	return type_GetUnionOutput_{}
}

type type_GetUnionOutput_ struct {
}

func (_this type_GetUnionOutput_) Default() interface{} {
	return Companion_GetUnionOutput_.Default()
}

func (_this type_GetUnionOutput_) String() string {
	return "simpleunioninternaldafnytypes.GetUnionOutput"
}
func (_this GetUnionOutput) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = GetUnionOutput{}

// End of datatype GetUnionOutput

// Definition of datatype KnownValueUnion
type KnownValueUnion struct {
	Data_KnownValueUnion_
}

func (_this KnownValueUnion) Get_() Data_KnownValueUnion_ {
	return _this.Data_KnownValueUnion_
}

type Data_KnownValueUnion_ interface {
	isKnownValueUnion()
}

type CompanionStruct_KnownValueUnion_ struct {
}

var Companion_KnownValueUnion_ = CompanionStruct_KnownValueUnion_{}

type KnownValueUnion_Value struct {
	Value int32
}

func (KnownValueUnion_Value) isKnownValueUnion() {}

func (CompanionStruct_KnownValueUnion_) Create_Value_(Value int32) KnownValueUnion {
	return KnownValueUnion{KnownValueUnion_Value{Value}}
}

func (_this KnownValueUnion) Is_Value() bool {
	_, ok := _this.Get_().(KnownValueUnion_Value)
	return ok
}

func (CompanionStruct_KnownValueUnion_) Default() KnownValueUnion {
	return Companion_KnownValueUnion_.Create_Value_(int32(0))
}

func (_this KnownValueUnion) Dtor_Value() int32 {
	return _this.Get_().(KnownValueUnion_Value).Value
}

func (_this KnownValueUnion) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case KnownValueUnion_Value:
		{
			return "SimpleUnionTypes.KnownValueUnion.Value" + "(" + _dafny.String(data.Value) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this KnownValueUnion) Equals(other KnownValueUnion) bool {
	switch data1 := _this.Get_().(type) {
	case KnownValueUnion_Value:
		{
			data2, ok := other.Get_().(KnownValueUnion_Value)
			return ok && data1.Value == data2.Value
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this KnownValueUnion) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(KnownValueUnion)
	return ok && _this.Equals(typed)
}

func Type_KnownValueUnion_() _dafny.TypeDescriptor {
	return type_KnownValueUnion_{}
}

type type_KnownValueUnion_ struct {
}

func (_this type_KnownValueUnion_) Default() interface{} {
	return Companion_KnownValueUnion_.Default()
}

func (_this type_KnownValueUnion_) String() string {
	return "simpleunioninternaldafnytypes.KnownValueUnion"
}
func (_this KnownValueUnion) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = KnownValueUnion{}

// End of datatype KnownValueUnion

// Definition of datatype MyUnion
type MyUnion struct {
	Data_MyUnion_
}

func (_this MyUnion) Get_() Data_MyUnion_ {
	return _this.Data_MyUnion_
}

type Data_MyUnion_ interface {
	isMyUnion()
}

type CompanionStruct_MyUnion_ struct {
}

var Companion_MyUnion_ = CompanionStruct_MyUnion_{}

type MyUnion_IntegerValue struct {
	IntegerValue int32
}

func (MyUnion_IntegerValue) isMyUnion() {}

func (CompanionStruct_MyUnion_) Create_IntegerValue_(IntegerValue int32) MyUnion {
	return MyUnion{MyUnion_IntegerValue{IntegerValue}}
}

func (_this MyUnion) Is_IntegerValue() bool {
	_, ok := _this.Get_().(MyUnion_IntegerValue)
	return ok
}

type MyUnion_StringValue struct {
	StringValue _dafny.Sequence
}

func (MyUnion_StringValue) isMyUnion() {}

func (CompanionStruct_MyUnion_) Create_StringValue_(StringValue _dafny.Sequence) MyUnion {
	return MyUnion{MyUnion_StringValue{StringValue}}
}

func (_this MyUnion) Is_StringValue() bool {
	_, ok := _this.Get_().(MyUnion_StringValue)
	return ok
}

func (CompanionStruct_MyUnion_) Default() MyUnion {
	return Companion_MyUnion_.Create_IntegerValue_(int32(0))
}

func (_this MyUnion) Dtor_IntegerValue() int32 {
	return _this.Get_().(MyUnion_IntegerValue).IntegerValue
}

func (_this MyUnion) Dtor_StringValue() _dafny.Sequence {
	return _this.Get_().(MyUnion_StringValue).StringValue
}

func (_this MyUnion) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case MyUnion_IntegerValue:
		{
			return "SimpleUnionTypes.MyUnion.IntegerValue" + "(" + _dafny.String(data.IntegerValue) + ")"
		}
	case MyUnion_StringValue:
		{
			return "SimpleUnionTypes.MyUnion.StringValue" + "(" + _dafny.String(data.StringValue) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this MyUnion) Equals(other MyUnion) bool {
	switch data1 := _this.Get_().(type) {
	case MyUnion_IntegerValue:
		{
			data2, ok := other.Get_().(MyUnion_IntegerValue)
			return ok && data1.IntegerValue == data2.IntegerValue
		}
	case MyUnion_StringValue:
		{
			data2, ok := other.Get_().(MyUnion_StringValue)
			return ok && data1.StringValue.Equals(data2.StringValue)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this MyUnion) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(MyUnion)
	return ok && _this.Equals(typed)
}

func Type_MyUnion_() _dafny.TypeDescriptor {
	return type_MyUnion_{}
}

type type_MyUnion_ struct {
}

func (_this type_MyUnion_) Default() interface{} {
	return Companion_MyUnion_.Default()
}

func (_this type_MyUnion_) String() string {
	return "simpleunioninternaldafnytypes.MyUnion"
}
func (_this MyUnion) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = MyUnion{}

// End of datatype MyUnion

// Definition of class ISimpleUnionClientCallHistory
type ISimpleUnionClientCallHistory struct {
	dummy byte
}

func New_ISimpleUnionClientCallHistory_() *ISimpleUnionClientCallHistory {
	_this := ISimpleUnionClientCallHistory{}

	return &_this
}

type CompanionStruct_ISimpleUnionClientCallHistory_ struct {
}

var Companion_ISimpleUnionClientCallHistory_ = CompanionStruct_ISimpleUnionClientCallHistory_{}

func (_this *ISimpleUnionClientCallHistory) Equals(other *ISimpleUnionClientCallHistory) bool {
	return _this == other
}

func (_this *ISimpleUnionClientCallHistory) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*ISimpleUnionClientCallHistory)
	return ok && _this.Equals(other)
}

func (*ISimpleUnionClientCallHistory) String() string {
	return "simpleunioninternaldafnytypes.ISimpleUnionClientCallHistory"
}

func Type_ISimpleUnionClientCallHistory_() _dafny.TypeDescriptor {
	return type_ISimpleUnionClientCallHistory_{}
}

type type_ISimpleUnionClientCallHistory_ struct {
}

func (_this type_ISimpleUnionClientCallHistory_) Default() interface{} {
	return (*ISimpleUnionClientCallHistory)(nil)
}

func (_this type_ISimpleUnionClientCallHistory_) String() string {
	return "simpleunioninternaldafnytypes.ISimpleUnionClientCallHistory"
}
func (_this *ISimpleUnionClientCallHistory) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &ISimpleUnionClientCallHistory{}

// End of class ISimpleUnionClientCallHistory

// Definition of trait ISimpleUnionClient
type ISimpleUnionClient interface {
	String() string
	GetUnion(input GetUnionInput) Wrappers.Result
	GetKnownValueUnion(input GetKnownValueUnionInput) Wrappers.Result
}
type CompanionStruct_ISimpleUnionClient_ struct {
	TraitID_ *_dafny.TraitID
}

var Companion_ISimpleUnionClient_ = CompanionStruct_ISimpleUnionClient_{
	TraitID_: &_dafny.TraitID{},
}

func (CompanionStruct_ISimpleUnionClient_) CastTo_(x interface{}) ISimpleUnionClient {
	var t ISimpleUnionClient
	t, _ = x.(ISimpleUnionClient)
	return t
}

// End of trait ISimpleUnionClient

// Definition of datatype SimpleUnionConfig
type SimpleUnionConfig struct {
	Data_SimpleUnionConfig_
}

func (_this SimpleUnionConfig) Get_() Data_SimpleUnionConfig_ {
	return _this.Data_SimpleUnionConfig_
}

type Data_SimpleUnionConfig_ interface {
	isSimpleUnionConfig()
}

type CompanionStruct_SimpleUnionConfig_ struct {
}

var Companion_SimpleUnionConfig_ = CompanionStruct_SimpleUnionConfig_{}

type SimpleUnionConfig_SimpleUnionConfig struct {
}

func (SimpleUnionConfig_SimpleUnionConfig) isSimpleUnionConfig() {}

func (CompanionStruct_SimpleUnionConfig_) Create_SimpleUnionConfig_() SimpleUnionConfig {
	return SimpleUnionConfig{SimpleUnionConfig_SimpleUnionConfig{}}
}

func (_this SimpleUnionConfig) Is_SimpleUnionConfig() bool {
	_, ok := _this.Get_().(SimpleUnionConfig_SimpleUnionConfig)
	return ok
}

func (CompanionStruct_SimpleUnionConfig_) Default() SimpleUnionConfig {
	return Companion_SimpleUnionConfig_.Create_SimpleUnionConfig_()
}

func (_ CompanionStruct_SimpleUnionConfig_) AllSingletonConstructors() _dafny.Iterator {
	i := -1
	return func() (interface{}, bool) {
		i++
		switch i {
		case 0:
			return Companion_SimpleUnionConfig_.Create_SimpleUnionConfig_(), true
		default:
			return SimpleUnionConfig{}, false
		}
	}
}

func (_this SimpleUnionConfig) String() string {
	switch _this.Get_().(type) {
	case nil:
		return "null"
	case SimpleUnionConfig_SimpleUnionConfig:
		{
			return "SimpleUnionTypes.SimpleUnionConfig.SimpleUnionConfig"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this SimpleUnionConfig) Equals(other SimpleUnionConfig) bool {
	switch _this.Get_().(type) {
	case SimpleUnionConfig_SimpleUnionConfig:
		{
			_, ok := other.Get_().(SimpleUnionConfig_SimpleUnionConfig)
			return ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this SimpleUnionConfig) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(SimpleUnionConfig)
	return ok && _this.Equals(typed)
}

func Type_SimpleUnionConfig_() _dafny.TypeDescriptor {
	return type_SimpleUnionConfig_{}
}

type type_SimpleUnionConfig_ struct {
}

func (_this type_SimpleUnionConfig_) Default() interface{} {
	return Companion_SimpleUnionConfig_.Default()
}

func (_this type_SimpleUnionConfig_) String() string {
	return "simpleunioninternaldafnytypes.SimpleUnionConfig"
}
func (_this SimpleUnionConfig) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = SimpleUnionConfig{}

// End of datatype SimpleUnionConfig

// Definition of datatype Error
type Error struct {
	Data_Error_
}

func (_this Error) Get_() Data_Error_ {
	return _this.Data_Error_
}

type Data_Error_ interface {
	isError()
}

type CompanionStruct_Error_ struct {
}

var Companion_Error_ = CompanionStruct_Error_{}

type Error_CollectionOfErrors struct {
	List    _dafny.Sequence
	Message _dafny.Sequence
}

func (Error_CollectionOfErrors) isError() {}

func (CompanionStruct_Error_) Create_CollectionOfErrors_(List _dafny.Sequence, Message _dafny.Sequence) Error {
	return Error{Error_CollectionOfErrors{List, Message}}
}

func (_this Error) Is_CollectionOfErrors() bool {
	_, ok := _this.Get_().(Error_CollectionOfErrors)
	return ok
}

type Error_Opaque struct {
	Obj interface{}
}

func (Error_Opaque) isError() {}

func (CompanionStruct_Error_) Create_Opaque_(Obj interface{}) Error {
	return Error{Error_Opaque{Obj}}
}

func (_this Error) Is_Opaque() bool {
	_, ok := _this.Get_().(Error_Opaque)
	return ok
}

func (CompanionStruct_Error_) Default() Error {
	return Companion_Error_.Create_CollectionOfErrors_(_dafny.EmptySeq, _dafny.EmptySeq.SetString())
}

func (_this Error) Dtor_list() _dafny.Sequence {
	return _this.Get_().(Error_CollectionOfErrors).List
}

func (_this Error) Dtor_message() _dafny.Sequence {
	return _this.Get_().(Error_CollectionOfErrors).Message
}

func (_this Error) Dtor_obj() interface{} {
	return _this.Get_().(Error_Opaque).Obj
}

func (_this Error) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case Error_CollectionOfErrors:
		{
			return "SimpleUnionTypes.Error.CollectionOfErrors" + "(" + _dafny.String(data.List) + ", " + _dafny.String(data.Message) + ")"
		}
	case Error_Opaque:
		{
			return "SimpleUnionTypes.Error.Opaque" + "(" + _dafny.String(data.Obj) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this Error) Equals(other Error) bool {
	switch data1 := _this.Get_().(type) {
	case Error_CollectionOfErrors:
		{
			data2, ok := other.Get_().(Error_CollectionOfErrors)
			return ok && data1.List.Equals(data2.List) && data1.Message.Equals(data2.Message)
		}
	case Error_Opaque:
		{
			data2, ok := other.Get_().(Error_Opaque)
			return ok && _dafny.AreEqual(data1.Obj, data2.Obj)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this Error) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(Error)
	return ok && _this.Equals(typed)
}

func Type_Error_() _dafny.TypeDescriptor {
	return type_Error_{}
}

type type_Error_ struct {
}

func (_this type_Error_) Default() interface{} {
	return Companion_Error_.Default()
}

func (_this type_Error_) String() string {
	return "simpleunioninternaldafnytypes.Error"
}
func (_this Error) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = Error{}

// End of datatype Error

// Definition of class OpaqueError
type OpaqueError struct {
}

func New_OpaqueError_() *OpaqueError {
	_this := OpaqueError{}

	return &_this
}

type CompanionStruct_OpaqueError_ struct {
}

var Companion_OpaqueError_ = CompanionStruct_OpaqueError_{}

func (*OpaqueError) String() string {
	return "simpleunioninternaldafnytypes.OpaqueError"
}

// End of class OpaqueError

func Type_OpaqueError_() _dafny.TypeDescriptor {
	return type_OpaqueError_{}
}

type type_OpaqueError_ struct {
}

func (_this type_OpaqueError_) Default() interface{} {
	return Companion_Error_.Default()
}

func (_this type_OpaqueError_) String() string {
	return "simpleunioninternaldafnytypes.OpaqueError"
}
func (_this *CompanionStruct_OpaqueError_) Is_(__source Error) bool {
	var _0_e Error = (__source)
	_ = _0_e
	return (_0_e).Is_Opaque()
}
