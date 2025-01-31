// Code generated by software.amazon.smithy.rust.codegen.smithy-rs. DO NOT EDIT.
#[allow(dead_code)]
pub fn to_dafny(
    value: crate::operation::get_blob_known_value::GetBlobKnownValueInput,
) -> ::std::rc::Rc<crate::r#simple::types::blob::internaldafny::types::GetBlobInput>
{
    let dafny_value = match value.value {
        Some(v) => crate::_Wrappers_Compile::Option::Some {
            value: ::dafny_runtime::dafny_runtime_conversions::vec_to_dafny_sequence(&v, |x| *x),
        },
        None => crate::_Wrappers_Compile::Option::None {},
    };
    ::std::rc::Rc::new(crate::r#simple::types::blob::internaldafny::types::GetBlobInput::GetBlobInput {
    value: ::std::rc::Rc::new(dafny_value)
  })
}

#[allow(dead_code)]
pub fn from_dafny(
    dafny_value: ::std::rc::Rc<
        crate::r#simple::types::blob::internaldafny::types::GetBlobInput,
    >,
) -> crate::operation::get_blob_known_value::GetBlobKnownValueInput {
    let value = if matches!(
        dafny_value.value().as_ref(),
        crate::_Wrappers_Compile::Option::Some { .. }
    ) {
        Some(
            ::std::rc::Rc::try_unwrap(dafny_value.value().Extract().to_array())
                .unwrap_or_else(|rc| (*rc).clone()),
        )
    } else if matches!(
        dafny_value.value().as_ref(),
        crate::_Wrappers_Compile::Option::None { .. }
    ) {
        None
    } else {
        panic!("Unreachable")
    };
    crate::operation::get_blob_known_value::GetBlobKnownValueInput { value }
}
