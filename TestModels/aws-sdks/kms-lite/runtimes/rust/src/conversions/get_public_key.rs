// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
// Do not modify this file. This file is machine generated, and any changes to it will be overwritten.
pub mod _get_public_key_request;

 pub mod _get_public_key_response;
 #[allow(dead_code)]
pub fn to_dafny_error(
    value: &::aws_smithy_runtime_api::client::result::SdkError<
        aws_sdk_kms::operation::get_public_key::GetPublicKeyError,
        ::aws_smithy_runtime_api::client::orchestrator::HttpResponse,
    >,
) -> ::std::rc::Rc<crate::r#software::amazon::cryptography::services::kms::internaldafny::types::Error> {
    match value {
      aws_sdk_kms::error::SdkError::ServiceError(service_error) => match service_error.err() {
                aws_sdk_kms::operation::get_public_key::GetPublicKeyError::NotFoundException(e) =>
            crate::conversions::error::not_found_exception::to_dafny(e.clone()),
         aws_sdk_kms::operation::get_public_key::GetPublicKeyError::DisabledException(e) =>
            crate::conversions::error::disabled_exception::to_dafny(e.clone()),
         aws_sdk_kms::operation::get_public_key::GetPublicKeyError::InvalidArnException(e) =>
            crate::conversions::error::invalid_arn_exception::to_dafny(e.clone()),
         aws_sdk_kms::operation::get_public_key::GetPublicKeyError::UnsupportedOperationException(e) =>
            crate::conversions::error::unsupported_operation_exception::to_dafny(e.clone()),
         aws_sdk_kms::operation::get_public_key::GetPublicKeyError::KmsInternalException(e) =>
            crate::conversions::error::kms_internal_exception::to_dafny(e.clone()),
         aws_sdk_kms::operation::get_public_key::GetPublicKeyError::KmsInvalidStateException(e) =>
            crate::conversions::error::kms_invalid_state_exception::to_dafny(e.clone()),
         aws_sdk_kms::operation::get_public_key::GetPublicKeyError::InvalidKeyUsageException(e) =>
            crate::conversions::error::invalid_key_usage_exception::to_dafny(e.clone()),
         aws_sdk_kms::operation::get_public_key::GetPublicKeyError::KeyUnavailableException(e) =>
            crate::conversions::error::key_unavailable_exception::to_dafny(e.clone()),
         aws_sdk_kms::operation::get_public_key::GetPublicKeyError::DependencyTimeoutException(e) =>
            crate::conversions::error::dependency_timeout_exception::to_dafny(e.clone()),
         aws_sdk_kms::operation::get_public_key::GetPublicKeyError::InvalidGrantTokenException(e) =>
            crate::conversions::error::invalid_grant_token_exception::to_dafny(e.clone()),
        e => crate::conversions::error::to_opaque_error(e.to_string()),
      },
      _ => {
        crate::conversions::error::to_opaque_error(value.to_string())
      }
   }
}
