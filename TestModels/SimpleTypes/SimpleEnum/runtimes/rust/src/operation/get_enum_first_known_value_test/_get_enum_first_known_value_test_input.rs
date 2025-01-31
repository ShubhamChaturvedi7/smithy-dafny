// Code generated by software.amazon.smithy.rust.codegen.smithy-rs. DO NOT EDIT.

#[allow(missing_docs)] // documentation missing in model
#[non_exhaustive]
#[derive(::std::clone::Clone, ::std::cmp::PartialEq, ::std::fmt::Debug)]
pub struct GetEnumFirstKnownValueTestInput {
    #[allow(missing_docs)] // documentation missing in model
    pub value: ::std::option::Option<crate::types::simple_enum_shape::SimpleEnumShape>,
}
impl GetEnumFirstKnownValueTestInput {
    #[allow(missing_docs)] // documentation missing in model
    pub fn value(&self) -> ::std::option::Option<crate::types::simple_enum_shape::SimpleEnumShape> {
        self.value
    }
}
impl GetEnumFirstKnownValueTestInput {
    /// Creates a new builder-style object to manufacture [`GetEnumFirstKnownValueTestInput`](crate::operation::operation::GetEnumFirstKnownValueTestInput).
    pub fn builder(
    ) -> crate::operation::get_enum_first_known_value_test::builders::GetEnumFirstKnownValueTestInputBuilder
    {
        crate::operation::get_enum_first_known_value_test::builders::GetEnumFirstKnownValueTestInputBuilder::default()
    }
}

/// A builder for [`GetEnumFirstKnownValueTestInput`](crate::operation::operation::GetEnumFirstKnownValueTestInput).
#[non_exhaustive]
#[derive(
    ::std::clone::Clone, ::std::cmp::PartialEq, ::std::default::Default, ::std::fmt::Debug,
)]
pub struct GetEnumFirstKnownValueTestInputBuilder {
    pub(crate) value: ::std::option::Option<crate::types::simple_enum_shape::SimpleEnumShape>,
}
impl GetEnumFirstKnownValueTestInputBuilder {
    #[allow(missing_docs)] // documentation missing in model
    pub fn value(
        mut self,
        input: impl ::std::convert::Into<crate::types::simple_enum_shape::SimpleEnumShape>,
    ) -> Self {
        self.value = ::std::option::Option::Some(input.into());
        self
    }
    #[allow(missing_docs)] // documentation missing in model
    pub fn set_value(
        mut self,
        input: ::std::option::Option<crate::types::simple_enum_shape::SimpleEnumShape>,
    ) -> Self {
        self.value = input;
        self
    }
    #[allow(missing_docs)] // documentation missing in model
    pub fn get_value(
        &self,
    ) -> &::std::option::Option<crate::types::simple_enum_shape::SimpleEnumShape> {
        &self.value
    }
    /// Consumes the builder and constructs a [`GetEnumFirstKnownValueTestInput`](crate::operation::operation::GetEnumFirstKnownValueTestInput).
    pub fn build(
        self,
    ) -> ::std::result::Result<
        crate::operation::get_enum_first_known_value_test::GetEnumFirstKnownValueTestInput,
        ::aws_smithy_types::error::operation::BuildError,
    > {
        ::std::result::Result::Ok(
            crate::operation::get_enum_first_known_value_test::GetEnumFirstKnownValueTestInput {
                value: self.value,
            },
        )
    }
}
