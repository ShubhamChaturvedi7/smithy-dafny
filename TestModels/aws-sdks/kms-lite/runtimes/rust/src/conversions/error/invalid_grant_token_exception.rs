// Code generated by software.amazon.smithy.rust.codegen.smithy-rs. DO NOT EDIT.

#[allow(dead_code)]
pub fn to_dafny(
    value: aws_sdk_kms::types::error::InvalidGrantTokenException,
) -> ::std::rc::Rc<crate::r#software::amazon::cryptography::services::kms::internaldafny::types::Error>{
  ::std::rc::Rc::new(
    crate::r#software::amazon::cryptography::services::kms::internaldafny::types::Error::InvalidGrantTokenException {
      message: crate::standard_library_conversions::ostring_to_dafny(&value.message)
    }
  )
}
