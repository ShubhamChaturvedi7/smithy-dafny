// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
// Do not modify this file. This file is machine generated, and any changes to it will be overwritten.
#[allow(dead_code)]
pub fn to_dafny(
    value: &aws_sdk_kms::operation::generate_data_key_without_plaintext::GenerateDataKeyWithoutPlaintextInput,
) -> ::std::rc::Rc<
    crate::r#software::amazon::cryptography::services::kms::internaldafny::types::GenerateDataKeyWithoutPlaintextRequest,
>{
    ::std::rc::Rc::new(crate::r#software::amazon::cryptography::services::kms::internaldafny::types::GenerateDataKeyWithoutPlaintextRequest::GenerateDataKeyWithoutPlaintextRequest {
        KeyId: crate::standard_library_conversions::ostring_to_dafny(&value.key_id) .Extract(),
 EncryptionContext:
::std::rc::Rc::new(match &value.encryption_context {
    Some(x) => crate::r#_Wrappers_Compile::Option::Some { value :
        ::dafny_runtime::dafny_runtime_conversions::hashmap_to_dafny_map(x,
            |k| dafny_runtime::dafny_runtime_conversions::unicode_chars_false::string_to_dafny_string(&k),
            |v| dafny_runtime::dafny_runtime_conversions::unicode_chars_false::string_to_dafny_string(&v),
        )
    },
    None => crate::r#_Wrappers_Compile::Option::None {}
})
,
 KeySpec: ::std::rc::Rc::new(match &value.key_spec {
    Some(x) => crate::_Wrappers_Compile::Option::Some { value: crate::conversions::data_key_spec::to_dafny(x.clone()) },
    None => crate::_Wrappers_Compile::Option::None { }
})
,
 NumberOfBytes: crate::standard_library_conversions::oint_to_dafny(value.number_of_bytes),
 GrantTokens: ::std::rc::Rc::new(match &value.grant_tokens {
    Some(x) => crate::r#_Wrappers_Compile::Option::Some { value :
        ::dafny_runtime::dafny_runtime_conversions::vec_to_dafny_sequence(x,
            |e| dafny_runtime::dafny_runtime_conversions::unicode_chars_false::string_to_dafny_string(&e),
        )
    },
    None => crate::r#_Wrappers_Compile::Option::None {}
})
,
 DryRun: crate::standard_library_conversions::obool_to_dafny(&value.dry_run),
    })
}
 #[allow(dead_code)]
pub fn from_dafny(
    dafny_value: ::std::rc::Rc<
        crate::r#software::amazon::cryptography::services::kms::internaldafny::types::GenerateDataKeyWithoutPlaintextRequest,
    >,
    client: aws_sdk_kms::Client,
) -> aws_sdk_kms::operation::generate_data_key_without_plaintext::builders::GenerateDataKeyWithoutPlaintextFluentBuilder {
    client.generate_data_key_without_plaintext()
          .set_key_id(Some( dafny_runtime::dafny_runtime_conversions::unicode_chars_false::dafny_string_to_string(dafny_value.KeyId()) ))
 .set_encryption_context(match (*dafny_value.EncryptionContext()).as_ref() {
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
 .set_key_spec(match &**dafny_value.KeySpec() {
    crate::r#_Wrappers_Compile::Option::Some { value } => Some(
        crate::conversions::data_key_spec::from_dafny(value)
    ),
    _ => None,
}
)
 .set_number_of_bytes(crate::standard_library_conversions::oint_from_dafny(dafny_value.NumberOfBytes().clone()))
 .set_grant_tokens(match (*dafny_value.GrantTokens()).as_ref() {
    crate::r#_Wrappers_Compile::Option::Some { value } =>
        Some(
            ::dafny_runtime::dafny_runtime_conversions::dafny_sequence_to_vec(value,
                |e| dafny_runtime::dafny_runtime_conversions::unicode_chars_false::dafny_string_to_string(e),
            )
        ),
    _ => None
}
)
 .set_dry_run(crate::standard_library_conversions::obool_from_dafny(dafny_value.DryRun().clone()))
}
