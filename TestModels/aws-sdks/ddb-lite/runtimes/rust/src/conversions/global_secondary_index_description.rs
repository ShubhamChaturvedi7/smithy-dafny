// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
// Do not modify this file. This file is machine generated, and any changes to it will be overwritten.
#[allow(dead_code)]
pub fn to_dafny(
    value: &aws_sdk_dynamodb::types::GlobalSecondaryIndexDescription,
) -> ::std::rc::Rc<crate::r#software::amazon::cryptography::services::dynamodb::internaldafny::types::GlobalSecondaryIndexDescription>{
  ::std::rc::Rc::new(
    crate::r#software::amazon::cryptography::services::dynamodb::internaldafny::types::GlobalSecondaryIndexDescription::GlobalSecondaryIndexDescription {
        IndexName: crate::standard_library_conversions::ostring_to_dafny(&value.index_name),
 KeySchema: ::std::rc::Rc::new(match &value.key_schema {
    Some(x) => crate::r#_Wrappers_Compile::Option::Some { value :
        ::dafny_runtime::dafny_runtime_conversions::vec_to_dafny_sequence(x,
            |e| crate::conversions::key_schema_element::to_dafny(&e)
,
        )
    },
    None => crate::r#_Wrappers_Compile::Option::None {}
})
,
 Projection: ::std::rc::Rc::new(match &value.projection {
    Some(x) => crate::_Wrappers_Compile::Option::Some { value: crate::conversions::projection::to_dafny(&x) },
    None => crate::_Wrappers_Compile::Option::None { }
})
,
 IndexStatus: ::std::rc::Rc::new(match &value.index_status {
    Some(x) => crate::_Wrappers_Compile::Option::Some { value: crate::conversions::index_status::to_dafny(x.clone()) },
    None => crate::_Wrappers_Compile::Option::None { }
})
,
 Backfilling: crate::standard_library_conversions::obool_to_dafny(&value.backfilling),
 ProvisionedThroughput: ::std::rc::Rc::new(match &value.provisioned_throughput {
    Some(x) => crate::_Wrappers_Compile::Option::Some { value: crate::conversions::provisioned_throughput_description::to_dafny(&x) },
    None => crate::_Wrappers_Compile::Option::None { }
})
,
 IndexSizeBytes: crate::standard_library_conversions::olong_to_dafny(&value.index_size_bytes),
 ItemCount: crate::standard_library_conversions::olong_to_dafny(&value.item_count),
 IndexArn: crate::standard_library_conversions::ostring_to_dafny(&value.index_arn),
    }
  )
} #[allow(dead_code)]
pub fn from_dafny(
    dafny_value: ::std::rc::Rc<
        crate::r#software::amazon::cryptography::services::dynamodb::internaldafny::types::GlobalSecondaryIndexDescription,
    >,
) -> aws_sdk_dynamodb::types::GlobalSecondaryIndexDescription {
    aws_sdk_dynamodb::types::GlobalSecondaryIndexDescription::builder()
          .set_index_name(crate::standard_library_conversions::ostring_from_dafny(dafny_value.IndexName().clone()))
 .set_key_schema(match (*dafny_value.KeySchema()).as_ref() {
    crate::r#_Wrappers_Compile::Option::Some { value } =>
        Some(
            ::dafny_runtime::dafny_runtime_conversions::dafny_sequence_to_vec(value,
                |e| crate::conversions::key_schema_element::from_dafny(e.clone())
,
            )
        ),
    _ => None
}
)
 .set_projection(match (*dafny_value.Projection()).as_ref() {
    crate::r#_Wrappers_Compile::Option::Some { value } =>
        Some(crate::conversions::projection::from_dafny(value.clone())),
    _ => None,
}
)
 .set_index_status(match &**dafny_value.IndexStatus() {
    crate::r#_Wrappers_Compile::Option::Some { value } => Some(
        crate::conversions::index_status::from_dafny(value)
    ),
    _ => None,
}
)
 .set_backfilling(crate::standard_library_conversions::obool_from_dafny(dafny_value.Backfilling().clone()))
 .set_provisioned_throughput(match (*dafny_value.ProvisionedThroughput()).as_ref() {
    crate::r#_Wrappers_Compile::Option::Some { value } =>
        Some(crate::conversions::provisioned_throughput_description::from_dafny(value.clone())),
    _ => None,
}
)
 .set_index_size_bytes(crate::standard_library_conversions::olong_from_dafny(dafny_value.IndexSizeBytes().clone()))
 .set_item_count(crate::standard_library_conversions::olong_from_dafny(dafny_value.ItemCount().clone()))
 .set_index_arn(crate::standard_library_conversions::ostring_from_dafny(dafny_value.IndexArn().clone()))
          .build()

}
