// Code generated by software.amazon.smithy.rust.codegen.smithy-rs. DO NOT EDIT.
#[allow(dead_code)]
pub fn to_dafny(
    value: crate::operation::get_constructor::GetConstructorOutput,
) -> ::std::rc::Rc<
    crate::simple::constructor::internaldafny::types::GetConstructorOutput,
>{
    let crate::operation::get_constructor::GetConstructorOutput {
        internal_config_string,
        blob_value,
        boolean_value,
        string_value,
        integer_value,
        long_value,
    } = value;

    let internal_config_string = match internal_config_string {
        Some(s) => crate::_Wrappers_Compile::Option::Some { value: dafny_runtime::dafny_runtime_conversions::unicode_chars_false::string_to_dafny_string(&s) },
        None => crate::_Wrappers_Compile::Option::None {},
    };

    let blob_value = match blob_value {
        Some(v) => crate::_Wrappers_Compile::Option::Some {
            value: ::dafny_runtime::dafny_runtime_conversions::vec_to_dafny_sequence(&v, |x| *x),
        },
        None => crate::_Wrappers_Compile::Option::None {},
    };

    let boolean_value = match boolean_value {
        Some(value) => crate::_Wrappers_Compile::Option::Some { value },
        None => crate::_Wrappers_Compile::Option::None {},
    };

    let string_value = match string_value {
        Some(s) => crate::_Wrappers_Compile::Option::Some { value: dafny_runtime::dafny_runtime_conversions::unicode_chars_false::string_to_dafny_string(&s) },
        None => crate::_Wrappers_Compile::Option::None {},
    };

    let integer_value = match integer_value {
        Some(value) => crate::_Wrappers_Compile::Option::Some { value },
        None => crate::_Wrappers_Compile::Option::None {},
    };

    let long_value = match long_value {
        Some(value) => crate::_Wrappers_Compile::Option::Some { value },
        None => crate::_Wrappers_Compile::Option::None {},
    };

    ::std::rc::Rc::new(crate::simple::constructor::internaldafny::types::GetConstructorOutput::GetConstructorOutput {
        internalConfigString: ::std::rc::Rc::new(internal_config_string),
        blobValue: ::std::rc::Rc::new(blob_value),
        booleanValue: ::std::rc::Rc::new(boolean_value),
        stringValue: ::std::rc::Rc::new(string_value),
        integerValue: ::std::rc::Rc::new(integer_value),
        longValue: ::std::rc::Rc::new(long_value),
    })
}

#[allow(dead_code)]
pub fn from_dafny(
    dafny_value: ::std::rc::Rc<
        crate::simple::constructor::internaldafny::types::GetConstructorOutput,
    >,
) -> crate::operation::get_constructor::GetConstructorOutput {
    let internal_config_string = if matches!(
        dafny_value.internalConfigString().as_ref(),
        crate::_Wrappers_Compile::Option::Some { .. }
    ) {
        Some(
            dafny_runtime::dafny_runtime_conversions::unicode_chars_false::dafny_string_to_string(
                &dafny_value.internalConfigString().Extract(),
            ),
        )
    } else if matches!(
        dafny_value.internalConfigString().as_ref(),
        crate::_Wrappers_Compile::Option::None { .. }
    ) {
        None
    } else {
        panic!("Unreachable")
    };

    let blob_value = if matches!(
        dafny_value.blobValue().as_ref(),
        crate::_Wrappers_Compile::Option::Some { .. }
    ) {
        Some(
            dafny_runtime::dafny_runtime_conversions::dafny_sequence_to_vec(
                &dafny_value.blobValue().Extract(),
                |e| *e,
            ),
        )
    } else if matches!(
        dafny_value.blobValue().as_ref(),
        crate::_Wrappers_Compile::Option::None { .. }
    ) {
        None
    } else {
        panic!("Unreachable")
    };

    let boolean_value = if matches!(
        dafny_value.booleanValue().as_ref(),
        crate::_Wrappers_Compile::Option::Some { .. }
    ) {
        Some(dafny_value.booleanValue().Extract())
    } else if matches!(
        dafny_value.booleanValue().as_ref(),
        crate::_Wrappers_Compile::Option::None { .. }
    ) {
        None
    } else {
        panic!("Unreachable")
    };

    let string_value = if matches!(
        dafny_value.stringValue().as_ref(),
        crate::_Wrappers_Compile::Option::Some { .. }
    ) {
        Some(
            dafny_runtime::dafny_runtime_conversions::unicode_chars_false::dafny_string_to_string(
                &dafny_value.stringValue().Extract(),
            ),
        )
    } else if matches!(
        dafny_value.stringValue().as_ref(),
        crate::_Wrappers_Compile::Option::None { .. }
    ) {
        None
    } else {
        panic!("Unreachable")
    };

    let integer_value = if matches!(
        dafny_value.integerValue().as_ref(),
        crate::_Wrappers_Compile::Option::Some { .. }
    ) {
        Some(dafny_value.integerValue().Extract())
    } else if matches!(
        dafny_value.integerValue().as_ref(),
        crate::_Wrappers_Compile::Option::None { .. }
    ) {
        None
    } else {
        panic!("Unreachable")
    };

    let long_value = if matches!(
        dafny_value.longValue().as_ref(),
        crate::_Wrappers_Compile::Option::Some { .. }
    ) {
        Some(dafny_value.longValue().Extract())
    } else if matches!(
        dafny_value.longValue().as_ref(),
        crate::_Wrappers_Compile::Option::None { .. }
    ) {
        None
    } else {
        panic!("Unreachable")
    };

    crate::operation::get_constructor::GetConstructorOutput {
        internal_config_string,
        blob_value,
        boolean_value,
        string_value,
        integer_value,
        long_value,
    }
}
