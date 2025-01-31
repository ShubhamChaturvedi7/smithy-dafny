// Code generated by software.amazon.smithy.rust.codegen.smithy-rs. DO NOT EDIT.
#[allow(missing_docs)] // documentation missing in model
#[non_exhaustive]
#[derive(::std::clone::Clone, ::std::cmp::PartialEq, ::std::fmt::Debug)]
pub struct GetKnownValueUnionInput {
    #[allow(missing_docs)] // documentation missing in model
    pub value: ::std::option::Option<crate::types::_known_value_union::KnownValueUnion>,
}
impl GetKnownValueUnionInput {
    #[allow(missing_docs)] // documentation missing in model
    pub fn value(
        &self,
    ) -> &::std::option::Option<crate::types::_known_value_union::KnownValueUnion> {
        &self.value
    }
}
impl GetKnownValueUnionInput {
    /// Creates a new builder-style object to manufacture [`GetKnownValueUnionInput`](crate::operation::get_known_value_union::GetKnownValueUnionInput).
    pub fn builder(
    ) -> crate::operation::get_known_value_union::builders::GetKnownValueUnionInputBuilder {
        crate::operation::get_known_value_union::builders::GetKnownValueUnionInputBuilder::default()
    }
}

/// A builder for [`GetKnownValueUnionInput`](crate::operation::get_known_value_union::GetKnownValueUnionInput).
#[non_exhaustive]
#[derive(
    ::std::clone::Clone, ::std::cmp::PartialEq, ::std::default::Default, ::std::fmt::Debug,
)]
pub struct GetKnownValueUnionInputBuilder {
    pub(crate) value: ::std::option::Option<crate::types::_known_value_union::KnownValueUnion>,
}
impl GetKnownValueUnionInputBuilder {
    #[allow(missing_docs)] // documentation missing in model
    pub fn value(mut self, input: crate::types::_known_value_union::KnownValueUnion) -> Self {
        self.value = ::std::option::Option::Some(input);
        self
    }
    #[allow(missing_docs)] // documentation missing in model
    pub fn set_value(
        mut self,
        input: ::std::option::Option<crate::types::_known_value_union::KnownValueUnion>,
    ) -> Self {
        self.value = input;
        self
    }
    #[allow(missing_docs)] // documentation missing in model
    pub fn get_value(
        &self,
    ) -> &::std::option::Option<crate::types::_known_value_union::KnownValueUnion> {
        &self.value
    }
    /// Consumes the builder and constructs a [`GetKnownValueUnionInput`](crate::operation::get_known_value_union::GetKnownValueUnionInput).
    pub fn build(
        self,
    ) -> ::std::result::Result<
        crate::operation::get_known_value_union::GetKnownValueUnionInput,
        ::aws_smithy_types::error::operation::BuildError,
    > {
        ::std::result::Result::Ok(
            crate::operation::get_known_value_union::GetKnownValueUnionInput { value: self.value },
        )
    }
}
