// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
// Do not modify this file. This file is machine generated, and any changes to it will be overwritten.
#[allow(dead_code)]
pub fn to_dafny(
    value: &aws_sdk_dynamodb::operation::put_item::PutItemInput,
) -> ::std::rc::Rc<
    crate::r#software::amazon::cryptography::services::dynamodb::internaldafny::types::PutItemInput,
>{
    ::std::rc::Rc::new(crate::r#software::amazon::cryptography::services::dynamodb::internaldafny::types::PutItemInput::PutItemInput {
        TableName: crate::standard_library_conversions::ostring_to_dafny(&value.table_name) .Extract(),
 Item: ::dafny_runtime::dafny_runtime_conversions::hashmap_to_dafny_map(&value.item.clone().unwrap(),
    |k| dafny_runtime::dafny_runtime_conversions::unicode_chars_false::string_to_dafny_string(&k),
    |v| crate::conversions::attribute_value::to_dafny(&v)
,
)
,
 Expected:
::std::rc::Rc::new(match &value.expected {
    Some(x) => crate::r#_Wrappers_Compile::Option::Some { value :
        ::dafny_runtime::dafny_runtime_conversions::hashmap_to_dafny_map(x,
            |k| dafny_runtime::dafny_runtime_conversions::unicode_chars_false::string_to_dafny_string(&k),
            |v| crate::conversions::expected_attribute_value::to_dafny(&v)
,
        )
    },
    None => crate::r#_Wrappers_Compile::Option::None {}
})
,
 ReturnValues: ::std::rc::Rc::new(match &value.return_values {
    Some(x) => crate::_Wrappers_Compile::Option::Some { value: crate::conversions::return_value::to_dafny(x.clone()) },
    None => crate::_Wrappers_Compile::Option::None { }
})
,
 ReturnConsumedCapacity: ::std::rc::Rc::new(match &value.return_consumed_capacity {
    Some(x) => crate::_Wrappers_Compile::Option::Some { value: crate::conversions::return_consumed_capacity::to_dafny(x.clone()) },
    None => crate::_Wrappers_Compile::Option::None { }
})
,
 ReturnItemCollectionMetrics: ::std::rc::Rc::new(match &value.return_item_collection_metrics {
    Some(x) => crate::_Wrappers_Compile::Option::Some { value: crate::conversions::return_item_collection_metrics::to_dafny(x.clone()) },
    None => crate::_Wrappers_Compile::Option::None { }
})
,
 ConditionalOperator: ::std::rc::Rc::new(match &value.conditional_operator {
    Some(x) => crate::_Wrappers_Compile::Option::Some { value: crate::conversions::conditional_operator::to_dafny(x.clone()) },
    None => crate::_Wrappers_Compile::Option::None { }
})
,
 ConditionExpression: crate::standard_library_conversions::ostring_to_dafny(&value.condition_expression),
 ExpressionAttributeNames:
::std::rc::Rc::new(match &value.expression_attribute_names {
    Some(x) => crate::r#_Wrappers_Compile::Option::Some { value :
        ::dafny_runtime::dafny_runtime_conversions::hashmap_to_dafny_map(x,
            |k| dafny_runtime::dafny_runtime_conversions::unicode_chars_false::string_to_dafny_string(&k),
            |v| dafny_runtime::dafny_runtime_conversions::unicode_chars_false::string_to_dafny_string(&v),
        )
    },
    None => crate::r#_Wrappers_Compile::Option::None {}
})
,
 ExpressionAttributeValues:
::std::rc::Rc::new(match &value.expression_attribute_values {
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
    })
}
 #[allow(dead_code)]
pub fn from_dafny(
    dafny_value: ::std::rc::Rc<
        crate::r#software::amazon::cryptography::services::dynamodb::internaldafny::types::PutItemInput,
    >,
    client: aws_sdk_dynamodb::Client,
) -> aws_sdk_dynamodb::operation::put_item::builders::PutItemFluentBuilder {
    client.put_item()
          .set_table_name(Some( dafny_runtime::dafny_runtime_conversions::unicode_chars_false::dafny_string_to_string(dafny_value.TableName()) ))
 .set_item(Some( ::dafny_runtime::dafny_runtime_conversions::dafny_map_to_hashmap(&dafny_value.Item(),
    |k| dafny_runtime::dafny_runtime_conversions::unicode_chars_false::dafny_string_to_string(k),
    |v| crate::conversions::attribute_value::from_dafny(v.clone())
,
)
 ))
 .set_expected(match (*dafny_value.Expected()).as_ref() {
    crate::r#_Wrappers_Compile::Option::Some { value } =>
        Some(
            ::dafny_runtime::dafny_runtime_conversions::dafny_map_to_hashmap(value,
                |k| dafny_runtime::dafny_runtime_conversions::unicode_chars_false::dafny_string_to_string(k),
                |v| crate::conversions::expected_attribute_value::from_dafny(v.clone())
,
            )
        ),
    _ => None
}
)
 .set_return_values(match &**dafny_value.ReturnValues() {
    crate::r#_Wrappers_Compile::Option::Some { value } => Some(
        crate::conversions::return_value::from_dafny(value)
    ),
    _ => None,
}
)
 .set_return_consumed_capacity(match &**dafny_value.ReturnConsumedCapacity() {
    crate::r#_Wrappers_Compile::Option::Some { value } => Some(
        crate::conversions::return_consumed_capacity::from_dafny(value)
    ),
    _ => None,
}
)
 .set_return_item_collection_metrics(match &**dafny_value.ReturnItemCollectionMetrics() {
    crate::r#_Wrappers_Compile::Option::Some { value } => Some(
        crate::conversions::return_item_collection_metrics::from_dafny(value)
    ),
    _ => None,
}
)
 .set_conditional_operator(match &**dafny_value.ConditionalOperator() {
    crate::r#_Wrappers_Compile::Option::Some { value } => Some(
        crate::conversions::conditional_operator::from_dafny(value)
    ),
    _ => None,
}
)
 .set_condition_expression(crate::standard_library_conversions::ostring_from_dafny(dafny_value.ConditionExpression().clone()))
 .set_expression_attribute_names(match (*dafny_value.ExpressionAttributeNames()).as_ref() {
    crate::r#_Wrappers_Compile::Option::Some { value } =>
        Some(
            ::dafny_runtime::dafny_runtime_conversions::dafny_map_to_hashmap(value,
                |k| dafny_runtime::dafny_runtime_conversions::unicode_chars_false::dafny_string_to_string(k),
                |v| dafny_runtime::dafny_runtime_conversions::unicode_chars_false::dafny_string_to_string(v),
            )
        ),
    _ => None
}
)
 .set_expression_attribute_values(match (*dafny_value.ExpressionAttributeValues()).as_ref() {
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
}
