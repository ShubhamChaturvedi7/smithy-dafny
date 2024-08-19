// Code generated by software.amazon.smithy.rust.codegen.smithy-rs. DO NOT EDIT.

use dafny_standard_library::conversion::{
    obool_from_dafny, obool_to_dafny, oint_from_dafny, oint_to_dafny, olong_from_dafny,
    olong_to_dafny, ostring_from_dafny, ostring_to_dafny,
};

#[allow(dead_code)]
pub fn to_dafny(
    config: crate::config::Config,
) -> ::std::rc::Rc<
    crate::implementation_from_dafny::_simple_dconstructor_dinternaldafny_dtypes::SimpleConstructorConfig,
>{
    ::std::rc::Rc::new(
        crate::implementation_from_dafny::_simple_dconstructor_dinternaldafny_dtypes::SimpleConstructorConfig::SimpleConstructorConfig { blobValue: blob_to_dafny(config.blob_value), booleanValue: obool_to_dafny(config.boolean_value), stringValue: ostring_to_dafny(&config.string_value), integerValue: oint_to_dafny(config.integer_value), longValue: olong_to_dafny(config.long_value) })
}

fn blob_to_dafny(
    value: ::std::option::Option<::std::vec::Vec<u8>>,
) -> ::std::rc::Rc<
    crate::implementation_from_dafny::_Wrappers_Compile::Option<::dafny_runtime::Sequence<u8>>,
> {
    let v = match value {
        Some(v) => crate::implementation_from_dafny::_Wrappers_Compile::Option::Some {
            value: ::dafny_runtime::Sequence::from_array(&v),
        },
        None => crate::implementation_from_dafny::_Wrappers_Compile::Option::None {},
    };

    ::std::rc::Rc::new(v)
}

#[allow(dead_code)]
pub fn from_dafny(
    config: ::std::rc::Rc<
        crate::implementation_from_dafny::_simple_dconstructor_dinternaldafny_dtypes::SimpleConstructorConfig,
    >,
) -> crate::config::Config {
    crate::config::Config {
        blob_value: blob_from_dafny(config.blobValue().clone()),
        boolean_value: obool_from_dafny(config.booleanValue().clone()),
        string_value: ostring_from_dafny(config.stringValue().clone()),
        integer_value: oint_from_dafny(config.integerValue().clone()),
        long_value: olong_from_dafny(config.longValue().clone()),
    }
}

fn blob_from_dafny(
    value: ::std::rc::Rc<
        crate::implementation_from_dafny::_Wrappers_Compile::Option<::dafny_runtime::Sequence<u8>>,
    >,
) -> ::std::option::Option<::std::vec::Vec<u8>> {
    match value.as_ref() {
        crate::implementation_from_dafny::_Wrappers_Compile::Option::Some { value } => {
            Some(::std::rc::Rc::try_unwrap(value.to_array()).unwrap_or_else(|rc| (*rc).clone()))
        }
        crate::implementation_from_dafny::_Wrappers_Compile::Option::None {} => None,
        crate::implementation_from_dafny::_Wrappers_Compile::Option::_PhantomVariant(_) => {
            unreachable!()
        }
    }
}
