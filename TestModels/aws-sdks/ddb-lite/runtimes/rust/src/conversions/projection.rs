// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
// Do not modify this file. This file is machine generated, and any changes to it will be overwritten.
#[allow(dead_code)]
pub fn to_dafny(
    value: &aws_sdk_dynamodb::types::Projection,
) -> ::std::rc::Rc<crate::r#software::amazon::cryptography::services::dynamodb::internaldafny::types::Projection>{
  ::std::rc::Rc::new(
    crate::r#software::amazon::cryptography::services::dynamodb::internaldafny::types::Projection::Projection {
        ProjectionType: ::std::rc::Rc::new(match &value.projection_type {
    Some(x) => crate::_Wrappers_Compile::Option::Some { value: crate::conversions::projection_type::to_dafny(x.clone()) },
    None => crate::_Wrappers_Compile::Option::None { }
})
,
 NonKeyAttributes: ::std::rc::Rc::new(match &value.non_key_attributes {
    Some(x) => crate::r#_Wrappers_Compile::Option::Some { value :
        ::dafny_runtime::dafny_runtime_conversions::vec_to_dafny_sequence(x,
            |e| dafny_runtime::dafny_runtime_conversions::unicode_chars_false::string_to_dafny_string(&e),
        )
    },
    None => crate::r#_Wrappers_Compile::Option::None {}
})
,
    }
  )
} #[allow(dead_code)]
pub fn from_dafny(
    dafny_value: ::std::rc::Rc<
        crate::r#software::amazon::cryptography::services::dynamodb::internaldafny::types::Projection,
    >,
) -> aws_sdk_dynamodb::types::Projection {
    aws_sdk_dynamodb::types::Projection::builder()
          .set_projection_type(match &**dafny_value.ProjectionType() {
    crate::r#_Wrappers_Compile::Option::Some { value } => Some(
        crate::conversions::projection_type::from_dafny(value)
    ),
    _ => None,
}
)
 .set_non_key_attributes(match (*dafny_value.NonKeyAttributes()).as_ref() {
    crate::r#_Wrappers_Compile::Option::Some { value } =>
        Some(
            ::dafny_runtime::dafny_runtime_conversions::dafny_sequence_to_vec(value,
                |e| dafny_runtime::dafny_runtime_conversions::unicode_chars_false::dafny_string_to_string(e),
            )
        ),
    _ => None
}
)
          .build()

}
