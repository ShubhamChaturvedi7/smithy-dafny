// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
// Do not modify this file. This file is machine generated, and any changes to it will be overwritten.
#[allow(dead_code)]
pub fn to_dafny(
    value: &aws_sdk_dynamodb::types::CancellationReason,
) -> ::std::rc::Rc<crate::r#software::amazon::cryptography::services::dynamodb::internaldafny::types::CancellationReason>{
  ::std::rc::Rc::new(
    crate::r#software::amazon::cryptography::services::dynamodb::internaldafny::types::CancellationReason::CancellationReason {
        Item:
::std::rc::Rc::new(match &value.item {
    Some(x) => crate::r#_Wrappers_Compile::Option::Some { value :
        ::dafny_runtime::dafny_runtime_conversions::hashmap_to_dafny_map(x,
            |k| dafny_runtime::dafny_runtime_conversions::unicode_chars_false::string_to_dafny_string(&k),
            |v| crate::conversions::attribute_value::to_dafny(&v)
,
        )
    },
    None => crate::r#_Wrappers_Compile::Option::None {}
})
,
 Code: crate::standard_library_conversions::ostring_to_dafny(&value.code),
 Message: crate::standard_library_conversions::ostring_to_dafny(&value.message),
    }
  )
} #[allow(dead_code)]
pub fn from_dafny(
    dafny_value: ::std::rc::Rc<
        crate::r#software::amazon::cryptography::services::dynamodb::internaldafny::types::CancellationReason,
    >,
) -> aws_sdk_dynamodb::types::CancellationReason {
    aws_sdk_dynamodb::types::CancellationReason::builder()
          .set_item(match (*dafny_value.Item()).as_ref() {
    crate::r#_Wrappers_Compile::Option::Some { value } =>
        Some(
            ::dafny_runtime::dafny_runtime_conversions::dafny_map_to_hashmap(value,
                |k| dafny_runtime::dafny_runtime_conversions::unicode_chars_false::dafny_string_to_string(k),
                |v| crate::conversions::attribute_value::from_dafny(v.clone())
,
            )
        ),
    _ => None
}
)
 .set_code(crate::standard_library_conversions::ostring_from_dafny(dafny_value.Code().clone()))
 .set_message(crate::standard_library_conversions::ostring_from_dafny(dafny_value.Message().clone()))
          .build()

}
