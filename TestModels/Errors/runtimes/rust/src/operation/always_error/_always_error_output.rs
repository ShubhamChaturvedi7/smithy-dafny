// Code generated by software.amazon.smithy.rust.codegen.smithy-rs. DO NOT EDIT.
#[allow(missing_docs)] // documentation missing in model
#[non_exhaustive]
#[derive(::std::clone::Clone, ::std::cmp::PartialEq, ::std::fmt::Debug)]
pub struct AlwaysErrorOutput {
    #[allow(missing_docs)] // documentation missing in model
    pub value: ::std::option::Option<::std::string::String>,
}
impl AlwaysErrorOutput {
    #[allow(missing_docs)] // documentation missing in model
    pub fn value(&self) -> ::std::option::Option<&str> {
        self.value.as_deref()
    }
}
impl AlwaysErrorOutput {
    /// Creates a new builder-style object to manufacture [`AlwaysErrorOutput`](crate::operation::operation::AlwaysErrorOutput).
    pub fn builder() -> crate::operation::always_error::builders::AlwaysErrorOutputBuilder {
        crate::operation::always_error::builders::AlwaysErrorOutputBuilder::default()
    }
}

/// A builder for [`AlwaysErrorOutput`](crate::operation::operation::AlwaysErrorOutput).
#[non_exhaustive]
#[derive(
    ::std::clone::Clone, ::std::cmp::PartialEq, ::std::default::Default, ::std::fmt::Debug,
)]
pub struct AlwaysErrorOutputBuilder {
    pub(crate) value: ::std::option::Option<::std::string::String>,
}
impl AlwaysErrorOutputBuilder {
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
    /// Consumes the builder and constructs a [`AlwaysErrorOutput`](crate::operation::operation::AlwaysErrorOutput).
    pub fn build(
        self,
    ) -> ::std::result::Result<
        crate::operation::always_error::AlwaysErrorOutput,
        crate::types::error::Error,
    > {
        ::std::result::Result::Ok(crate::operation::always_error::AlwaysErrorOutput {
            value: self.value,
        })
    }
}
