// Code generated by software.amazon.smithy.rust.codegen.smithy-rs. DO NOT EDIT.
#[allow(dead_code)]
pub fn to_dafny(
    value: crate::operation::get_blob_known_value::GetBlobKnownValueOutput,
) -> ::std::rc::Rc<crate::r#simple::types::blob::internaldafny::types::GetBlobOutput>
{
    let dafny_value = match value.value {
        Some(v) => crate::_Wrappers_Compile::Option::Some {
            value: dafny_runtime::dafny_runtime_conversions::vec_to_dafny_sequence(&v, |e| *e),
        },
        None => crate::_Wrappers_Compile::Option::None {},
    };
    ::std::rc::Rc::new(crate::r#simple::types::blob::internaldafny::types::GetBlobOutput::GetBlobOutput {
    value: ::std::rc::Rc::new(dafny_value)
  })
}

#[allow(dead_code)]
pub fn from_dafny(
    dafny_value: ::std::rc::Rc<
        crate::r#simple::types::blob::internaldafny::types::GetBlobOutput,
    >,
) -> crate::operation::get_blob_known_value::GetBlobKnownValueOutput {
    let value = if matches!(
        dafny_value.value().as_ref(),
        crate::_Wrappers_Compile::Option::Some { .. }
    ) {
        Some(
            dafny_runtime::dafny_runtime_conversions::dafny_sequence_to_vec(
                &dafny_value.value().Extract(),
                |e| *e,
            ),
        )
    } else if matches!(
        dafny_value.value().as_ref(),
        crate::_Wrappers_Compile::Option::None { .. }
    ) {
        None
    } else {
        panic!("Unreachable")
    };
    crate::operation::get_blob_known_value::GetBlobKnownValueOutput { value }
}
