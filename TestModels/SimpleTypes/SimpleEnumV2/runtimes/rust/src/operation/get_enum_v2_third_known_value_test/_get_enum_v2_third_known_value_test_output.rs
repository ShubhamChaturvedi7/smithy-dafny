// Code generated by software.amazon.smithy.rust.codegen.smithy-rs. DO NOT EDIT.

#[allow(missing_docs)] // documentation missing in model
#[non_exhaustive]
#[derive(::std::clone::Clone, ::std::cmp::PartialEq, ::std::fmt::Debug)]
pub struct GetEnumV2ThirdKnownValueTestOutput {
    #[allow(missing_docs)] // documentation missing in model
    pub value: ::std::option::Option<crate::types::simple_enum_v2_shape::SimpleEnumV2Shape>,
}
impl GetEnumV2ThirdKnownValueTestOutput {
    #[allow(missing_docs)] // documentation missing in model
    pub fn value(
        &self,
    ) -> ::std::option::Option<crate::types::simple_enum_v2_shape::SimpleEnumV2Shape> {
        self.value
    }
}
impl GetEnumV2ThirdKnownValueTestOutput {
    /// Creates a new builder-style object to manufacture [`GetEnumV2ThirdKnownValueTestOutput`](crate::operation::operation::GetEnumV2ThirdKnownValueTestOutput).
    pub fn builder(
    ) -> crate::operation::get_enum_v2_third_known_value_test::builders::GetEnumV2ThirdKnownValueTestOutputBuilder
    {
        crate::operation::get_enum_v2_third_known_value_test::builders::GetEnumV2ThirdKnownValueTestOutputBuilder::default()
    }
}

/// A builder for [`GetEnumV2ThirdKnownValueTestOutput`](crate::operation::operation::GetEnumV2ThirdKnownValueTestOutput).
#[non_exhaustive]
#[derive(
    ::std::clone::Clone, ::std::cmp::PartialEq, ::std::default::Default, ::std::fmt::Debug,
)]
pub struct GetEnumV2ThirdKnownValueTestOutputBuilder {
    pub(crate) value: ::std::option::Option<crate::types::simple_enum_v2_shape::SimpleEnumV2Shape>,
}
impl GetEnumV2ThirdKnownValueTestOutputBuilder {
    #[allow(missing_docs)] // documentation missing in model
    pub fn value(
        mut self,
        input: impl ::std::convert::Into<crate::types::simple_enum_v2_shape::SimpleEnumV2Shape>,
    ) -> Self {
        self.value = ::std::option::Option::Some(input.into());
        self
    }
    #[allow(missing_docs)] // documentation missing in model
    pub fn set_value(
        mut self,
        input: ::std::option::Option<crate::types::simple_enum_v2_shape::SimpleEnumV2Shape>,
    ) -> Self {
        self.value = input;
        self
    }
    #[allow(missing_docs)] // documentation missing in model
    pub fn get_value(
        &self,
    ) -> &::std::option::Option<crate::types::simple_enum_v2_shape::SimpleEnumV2Shape> {
        &self.value
    }
    /// Consumes the builder and constructs a [`GetEnumV2ThirdKnownValueTestOutput`](crate::operation::operation::GetEnumV2ThirdKnownValueTestOutput).
    pub fn build(
        self,
    ) -> ::std::result::Result<
        crate::operation::get_enum_v2_third_known_value_test::GetEnumV2ThirdKnownValueTestOutput,
        ::aws_smithy_types::error::operation::BuildError,
    > {
        ::std::result::Result::Ok(
            crate::operation::get_enum_v2_third_known_value_test::GetEnumV2ThirdKnownValueTestOutput {
                value: self.value,
            },
        )
    }
}
