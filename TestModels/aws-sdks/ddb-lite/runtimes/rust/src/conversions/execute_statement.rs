// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
// Do not modify this file. This file is machine generated, and any changes to it will be overwritten.
pub mod _execute_statement_request;

 pub mod _execute_statement_response;
 #[allow(dead_code)]
pub fn to_dafny_error(
    value: &::aws_smithy_runtime_api::client::result::SdkError<
        aws_sdk_dynamodb::operation::execute_statement::ExecuteStatementError,
        ::aws_smithy_runtime_api::client::orchestrator::HttpResponse,
    >,
) -> ::std::rc::Rc<crate::r#software::amazon::cryptography::services::dynamodb::internaldafny::types::Error> {
    match value {
      aws_sdk_dynamodb::error::SdkError::ServiceError(service_error) => match service_error.err() {
                aws_sdk_dynamodb::operation::execute_statement::ExecuteStatementError::InternalServerError(e) =>
            crate::conversions::error::internal_server_error::to_dafny(e.clone()),
         aws_sdk_dynamodb::operation::execute_statement::ExecuteStatementError::ResourceNotFoundException(e) =>
            crate::conversions::error::resource_not_found_exception::to_dafny(e.clone()),
         aws_sdk_dynamodb::operation::execute_statement::ExecuteStatementError::DuplicateItemException(e) =>
            crate::conversions::error::duplicate_item_exception::to_dafny(e.clone()),
         aws_sdk_dynamodb::operation::execute_statement::ExecuteStatementError::RequestLimitExceeded(e) =>
            crate::conversions::error::request_limit_exceeded::to_dafny(e.clone()),
         aws_sdk_dynamodb::operation::execute_statement::ExecuteStatementError::TransactionConflictException(e) =>
            crate::conversions::error::transaction_conflict_exception::to_dafny(e.clone()),
         aws_sdk_dynamodb::operation::execute_statement::ExecuteStatementError::ConditionalCheckFailedException(e) =>
            crate::conversions::error::conditional_check_failed_exception::to_dafny(e.clone()),
         aws_sdk_dynamodb::operation::execute_statement::ExecuteStatementError::ItemCollectionSizeLimitExceededException(e) =>
            crate::conversions::error::item_collection_size_limit_exceeded_exception::to_dafny(e.clone()),
         aws_sdk_dynamodb::operation::execute_statement::ExecuteStatementError::ProvisionedThroughputExceededException(e) =>
            crate::conversions::error::provisioned_throughput_exceeded_exception::to_dafny(e.clone()),
        e => crate::conversions::error::to_opaque_error(e.to_string()),
      },
      _ => {
        crate::conversions::error::to_opaque_error(value.to_string())
      }
   }
}
