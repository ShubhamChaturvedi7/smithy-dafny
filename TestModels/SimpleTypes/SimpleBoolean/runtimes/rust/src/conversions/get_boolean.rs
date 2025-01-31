// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
// Do not modify this file. This file is machine generated, and any changes to it will be overwritten.
use std::any::Any;

#[allow(dead_code)]
pub fn to_dafny_error(
    value: crate::operation::get_boolean::GetBooleanError,
) -> ::std::rc::Rc<crate::r#simple::types::boolean::internaldafny::types::Error> {
    match value {
    crate::operation::get_boolean::GetBooleanError::Unhandled(unhandled) =>
      ::std::rc::Rc::new(crate::r#simple::types::boolean::internaldafny::types::Error::Opaque { obj: ::dafny_runtime::upcast_object()(::dafny_runtime::object::new(unhandled)) })
  }
}

#[allow(dead_code)]
pub fn from_dafny_error(
    dafny_value: ::std::rc::Rc<
        crate::r#simple::types::boolean::internaldafny::types::Error,
    >,
) -> crate::operation::get_boolean::GetBooleanError {
    // TODO: Losing information here, but we have to figure out how to wrap an arbitrary Dafny value as std::error::Error
    if matches!(&dafny_value.as_ref(), crate::r#simple::types::boolean::internaldafny::types::Error::CollectionOfErrors { .. }) {
    let error_message = "TODO: can't get message yet";
    crate::operation::get_boolean::GetBooleanError::generic(::aws_smithy_types::error::metadata::ErrorMetadata::builder().message(error_message).build())
  } else {
    crate::operation::get_boolean::GetBooleanError::generic(::aws_smithy_types::error::metadata::ErrorMetadata::builder().message("Opaque error").build())
  }
}

pub mod _get_boolean_input;

pub mod _get_boolean_output;
