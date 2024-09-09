// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
// Do not modify this file. This file is machine generated, and any changes to it will be overwritten.
use aws_smithy_types::error::operation::BuildError;

#[derive(::std::clone::Clone, ::std::fmt::Debug)]
pub struct Client {
    pub(crate) dafny_client: ::dafny_runtime::Object<dyn crate::r#simple::types::enumv2::internaldafny::types::ISimpleTypesEnumV2Client>
}

impl Client {
    /// Creates a new client from the service [`Config`](crate::Config).
    #[track_caller]
    pub fn from_conf(
        conf: crate::types::simple_enum_v2_config::SimpleEnumV2Config,
    ) -> Result<Self, BuildError> {
        let inner =
            crate::simple::types::enumv2::internaldafny::_default::SimpleEnumV2(
                &crate::conversions::simple_enum_v2_config::_simple_enum_v2_config::to_dafny(conf),
            );
        if matches!(
            inner.as_ref(),
            crate::_Wrappers_Compile::Result::Failure { .. }
        ) {
            // TODO: convert error - the potential types are not modeled!
            return Err(BuildError::other(
                ::aws_smithy_types::error::metadata::ErrorMetadata::builder()
                    .message("Invalid client config")
                    .build(),
            ));
        }
        Ok(Self {
            dafny_client: ::dafny_runtime::upcast_object()(inner.Extract())
        })
    }
}

mod get_enum_v2;

mod get_enum_v2_first_known_value_test;

mod get_enum_v2_second_known_value_test;

mod get_enum_v2_third_known_value_test;
