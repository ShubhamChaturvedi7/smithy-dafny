// Code generated by software.amazon.smithy.rust.codegen.smithy-rs. DO NOT EDIT.
#[allow(dead_code)]
pub fn to_dafny(
    value: crate::operation::get_double::GetDoubleInput,
) -> ::std::rc::Rc<
    crate::r#simple::types::smithydouble::internaldafny::types::GetDoubleInput,
> {
    let dafny_value = match value.value {
        Some(v) => crate::_Wrappers_Compile::Option::Some {
            value: ::dafny_runtime::Sequence::ArraySequence {
                values: std::rc::Rc::new(f64::to_be_bytes(v).to_vec()),
            },
        },
        None => crate::_Wrappers_Compile::Option::None {},
    };
    ::std::rc::Rc::new(crate::r#simple::types::smithydouble::internaldafny::types::GetDoubleInput::GetDoubleInput {
    value: ::std::rc::Rc::new(dafny_value)
  })
}

#[allow(dead_code)]
pub fn from_dafny(
    dafny_value: ::std::rc::Rc<
        crate::r#simple::types::smithydouble::internaldafny::types::GetDoubleInput,
    >,
) -> crate::operation::get_double::GetDoubleInput {
    let value = if matches!(
        dafny_value.value().as_ref(),
        crate::_Wrappers_Compile::Option::Some { .. }
    ) {
        let my_vec = dafny_runtime::dafny_runtime_conversions::dafny_sequence_to_vec(
            &dafny_value.value().Extract(),
            |x| *x,
        );
        Some(f64::from_be_bytes(
            my_vec.try_into().expect("Error converting Sequence to f64"),
        ))
    } else if matches!(
        dafny_value.value().as_ref(),
        crate::_Wrappers_Compile::Option::None { .. }
    ) {
        None
    } else {
        panic!("Unreachable")
    };
    crate::operation::get_double::GetDoubleInput { value }
}
