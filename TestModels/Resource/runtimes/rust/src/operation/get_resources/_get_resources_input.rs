// Code generated by software.amazon.smithy.rust.codegen.smithy-rs. DO NOT EDIT.
#[allow(missing_docs)] // documentation missing in model
#[non_exhaustive]
#[derive(::std::clone::Clone, ::std::cmp::PartialEq, ::std::fmt::Debug)]
pub struct GetResourcesInput {
    #[allow(missing_docs)] // documentation missing in model
    pub(crate) value: ::std::option::Option<::std::string::String>,
}
impl GetResourcesInput {
    #[allow(missing_docs)] // documentation missing in model
    pub fn value(&self) -> &::std::option::Option<String> {
        &self.value
    }
}
impl GetResourcesInput {
    /// Creates a new builder-style object to manufacture [`GetResourcesInput`](crate::operation::operation::GetResourcesInput).
    pub fn builder() -> crate::operation::get_resources::builders::GetResourcesInputBuilder {
        crate::operation::get_resources::builders::GetResourcesInputBuilder::default()
    }
}

/// A builder for [`GetResourcesInput`](crate::operation::operation::GetResourcesInput).
#[non_exhaustive]
#[derive(
    ::std::clone::Clone, ::std::cmp::PartialEq, ::std::default::Default, ::std::fmt::Debug,
)]
pub struct GetResourcesInputBuilder {
    pub(crate) value: ::std::option::Option<::std::string::String>,
}
impl GetResourcesInputBuilder {
    #[allow(missing_docs)] // documentation missing in model
    pub fn value(mut self, input: impl ::std::convert::Into<::std::string::String>) -> Self {
        self.value = ::std::option::Option::Some(input.into());
        self
    }
    #[allow(missing_docs)] // documentation missing in model
    pub fn set_value(mut self, input: ::std::option::Option<::std::string::String>) -> Self {
        self.value = input;
        self
    }
    #[allow(missing_docs)] // documentation missing in model
    pub fn get_value(&self) -> &::std::option::Option<::std::string::String> {
        &self.value
    }
    /// Consumes the builder and constructs a [`GetResourcesInput`](crate::operation::operation::GetResourcesInput).
    pub fn build(
        self,
    ) -> ::std::result::Result<
        crate::operation::get_resources::GetResourcesInput,
        ::aws_smithy_types::error::operation::BuildError,
    > {
        ::std::result::Result::Ok(crate::operation::get_resources::GetResourcesInput {
            value: self.value,
        })
    }
}
