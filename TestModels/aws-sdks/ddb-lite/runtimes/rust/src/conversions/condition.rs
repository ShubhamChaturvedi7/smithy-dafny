// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
// Do not modify this file. This file is machine generated, and any changes to it will be overwritten.
#[allow(dead_code)]
pub fn to_dafny(
    value: &aws_sdk_dynamodb::types::Condition,
) -> ::std::rc::Rc<crate::r#software::amazon::cryptography::services::dynamodb::internaldafny::types::Condition>{
  ::std::rc::Rc::new(
    crate::r#software::amazon::cryptography::services::dynamodb::internaldafny::types::Condition::Condition {
        AttributeValueList: ::std::rc::Rc::new(match &value.attribute_value_list {
    Some(x) => crate::r#_Wrappers_Compile::Option::Some { value :
        ::dafny_runtime::dafny_runtime_conversions::vec_to_dafny_sequence(x,
            |e| crate::conversions::attribute_value::to_dafny(&e)
,
        )
    },
    None => crate::r#_Wrappers_Compile::Option::None {}
})
,
 ComparisonOperator: crate::conversions::comparison_operator::to_dafny(value.comparison_operator.clone()),
    }
  )
} #[allow(dead_code)]
pub fn from_dafny(
    dafny_value: ::std::rc::Rc<
        crate::r#software::amazon::cryptography::services::dynamodb::internaldafny::types::Condition,
    >,
) -> aws_sdk_dynamodb::types::Condition {
    aws_sdk_dynamodb::types::Condition::builder()
          .set_attribute_value_list(match (*dafny_value.AttributeValueList()).as_ref() {
    crate::r#_Wrappers_Compile::Option::Some { value } =>
        Some(
            ::dafny_runtime::dafny_runtime_conversions::dafny_sequence_to_vec(value,
                |e| crate::conversions::attribute_value::from_dafny(e.clone())
,
            )
        ),
    _ => None
}
)
 .set_comparison_operator(Some( crate::conversions::comparison_operator::from_dafny(dafny_value.ComparisonOperator()) ))
          .build()
          .unwrap()
}
