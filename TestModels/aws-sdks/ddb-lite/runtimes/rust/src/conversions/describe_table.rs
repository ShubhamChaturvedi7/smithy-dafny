// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
// Do not modify this file. This file is machine generated, and any changes to it will be overwritten.
pub mod _describe_table_request;

 pub mod _describe_table_response;
 #[allow(dead_code)]
pub fn to_dafny_error(
    value: &::aws_smithy_runtime_api::client::result::SdkError<
        aws_sdk_dynamodb::operation::describe_table::DescribeTableError,
        ::aws_smithy_runtime_api::client::orchestrator::HttpResponse,
    >,
) -> ::std::rc::Rc<crate::r#software::amazon::cryptography::services::dynamodb::internaldafny::types::Error> {
    match value {
      aws_sdk_dynamodb::error::SdkError::ServiceError(service_error) => match service_error.err() {
                aws_sdk_dynamodb::operation::describe_table::DescribeTableError::InternalServerError(e) =>
            crate::conversions::error::internal_server_error::to_dafny(e.clone()),
         aws_sdk_dynamodb::operation::describe_table::DescribeTableError::InvalidEndpointException(e) =>
            crate::conversions::error::invalid_endpoint_exception::to_dafny(e.clone()),
         aws_sdk_dynamodb::operation::describe_table::DescribeTableError::ResourceNotFoundException(e) =>
            crate::conversions::error::resource_not_found_exception::to_dafny(e.clone()),
        e => crate::conversions::error::to_opaque_error(e.to_string()),
      },
      _ => {
        crate::conversions::error::to_opaque_error(value.to_string())
      }
   }
}
