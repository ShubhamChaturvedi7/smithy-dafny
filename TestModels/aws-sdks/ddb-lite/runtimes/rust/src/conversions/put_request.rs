// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
// Do not modify this file. This file is machine generated, and any changes to it will be overwritten.
#[allow(dead_code)]
pub fn to_dafny(
    value: &aws_sdk_dynamodb::types::PutRequest,
) -> ::std::rc::Rc<crate::r#software::amazon::cryptography::services::dynamodb::internaldafny::types::PutRequest>{
  ::std::rc::Rc::new(
    crate::r#software::amazon::cryptography::services::dynamodb::internaldafny::types::PutRequest::PutRequest {
        Item: ::dafny_runtime::dafny_runtime_conversions::hashmap_to_dafny_map(&value.item.clone(),
    |k| dafny_runtime::dafny_runtime_conversions::unicode_chars_false::string_to_dafny_string(&k),
    |v| crate::conversions::attribute_value::to_dafny(&v)
,
)
,
    }
  )
} #[allow(dead_code)]
pub fn from_dafny(
    dafny_value: ::std::rc::Rc<
        crate::r#software::amazon::cryptography::services::dynamodb::internaldafny::types::PutRequest,
    >,
) -> aws_sdk_dynamodb::types::PutRequest {
    aws_sdk_dynamodb::types::PutRequest::builder()
          .set_item(Some( ::dafny_runtime::dafny_runtime_conversions::dafny_map_to_hashmap(&dafny_value.Item(),
    |k| dafny_runtime::dafny_runtime_conversions::unicode_chars_false::dafny_string_to_string(k),
    |v| crate::conversions::attribute_value::from_dafny(v.clone())
,
)
 ))
          .build()
          .unwrap()
}
