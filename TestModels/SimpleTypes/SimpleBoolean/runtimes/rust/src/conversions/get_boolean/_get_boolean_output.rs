// Code generated by software.amazon.smithy.rust.codegen.smithy-rs. DO NOT EDIT.
#[allow(dead_code)]
pub fn to_dafny(
    value: crate::operation::get_boolean::GetBooleanOutput,
) -> ::std::rc::Rc<
    ::simple_boolean_dafny::r#_simple_dtypes_dboolean_dinternaldafny_dtypes::GetBooleanOutput,
> {
    let dafny_value = match value.value {
        Some(b) => ::simple_boolean_dafny::_Wrappers_Compile::Option::Some { value: b },
        None => ::simple_boolean_dafny::_Wrappers_Compile::Option::None {},
    };
    ::std::rc::Rc::new(::simple_boolean_dafny::r#_simple_dtypes_dboolean_dinternaldafny_dtypes::GetBooleanOutput::GetBooleanOutput {
    value: ::std::rc::Rc::new(dafny_value)
  })
}

#[allow(dead_code)]
pub fn from_dafny(
    dafny_value: ::std::rc::Rc<::simple_boolean_dafny::r#_simple_dtypes_dboolean_dinternaldafny_dtypes::GetBooleanOutput>,
) -> crate::operation::get_boolean::GetBooleanOutput {
    let value = if matches!(
        dafny_value.value().as_ref(),
        ::simple_boolean_dafny::_Wrappers_Compile::Option::Some { .. }
    ) {
        Some(dafny_value.value().Extract())
    } else if matches!(
        dafny_value.value().as_ref(),
        ::simple_boolean_dafny::_Wrappers_Compile::Option::None { .. }
    ) {
        None
    } else {
        panic!("Unreachable")
    };
    crate::operation::get_boolean::GetBooleanOutput { value }
}
