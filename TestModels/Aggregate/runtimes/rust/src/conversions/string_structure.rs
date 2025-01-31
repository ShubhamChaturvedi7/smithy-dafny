// Code generated by software.amazon.smithy.rust.codegen.smithy-rs. DO NOT EDIT.
#[allow(dead_code)]
pub fn to_dafny(
    value: crate::types::StringStructure,
) -> ::std::rc::Rc<
    crate::r#simple::aggregate::internaldafny::types::StringStructure,
> {
    ::std::rc::Rc::new(crate::r#simple::aggregate::internaldafny::types::StringStructure::StringStructure {
  value: crate::standard_library_conversions::ostring_to_dafny(&value.value),
})
}

pub fn to_dafny_plain(
    value: crate::types::StringStructure,
) -> crate::r#simple::aggregate::internaldafny::types::StringStructure {
    crate::r#simple::aggregate::internaldafny::types::StringStructure::StringStructure {
  value: crate::standard_library_conversions::ostring_to_dafny(&value.value),
}
}

pub fn option_to_dafny(
    value: Option<crate::types::StringStructure>,
) -> ::std::rc::Rc<
    crate::_Wrappers_Compile::Option<
        ::std::rc::Rc<
            crate::r#simple::aggregate::internaldafny::types::StringStructure,
        >,
    >,
> {
    let inner = match value {
        None => crate::_Wrappers_Compile::Option::None {},
        Some(x) => crate::_Wrappers_Compile::Option::Some {
            value: ::std::rc::Rc::new(to_dafny_plain(x)),
        },
    };
    ::std::rc::Rc::new(inner)
}

#[allow(dead_code)]
pub fn from_dafny(
    dafny_value: ::std::rc::Rc<
        crate::r#simple::aggregate::internaldafny::types::StringStructure,
    >,
) -> crate::types::StringStructure {
    plain_from_dafny(&*dafny_value)
}

#[allow(dead_code)]
pub fn plain_from_dafny(
    dafny_value : &crate::r#simple::aggregate::internaldafny::types::StringStructure,
) -> crate::types::StringStructure {
    match dafny_value {
    crate::r#simple::aggregate::internaldafny::types::StringStructure::StringStructure {
      value,
    } =>
    crate::types::StringStructure {
      value: crate::standard_library_conversions::ostring_from_dafny(value.clone()),
    }
}
}

#[allow(dead_code)]
pub fn option_from_dafny(
    dafny_value: ::std::rc::Rc<crate::_Wrappers_Compile::Option<::std::rc::Rc<
        crate::r#simple::aggregate::internaldafny::types::StringStructure,
    >>>,
) -> Option<crate::types::StringStructure> {
    match &*dafny_value {
        crate::_Wrappers_Compile::Option::Some { value } => {
            Some(plain_from_dafny(value))
        }
        _ => None,
    }
}
