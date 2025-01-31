// Code generated by software.amazon.smithy.rust.codegen.smithy-rs. DO NOT EDIT.
pub use crate::operation::get_aggregate_known_value_test::_get_aggregate_known_value_test_output::GetAggregateKnownValueTestOutputBuilder;

pub use crate::operation::get_aggregate_known_value_test::_get_aggregate_known_value_test_input::GetAggregateKnownValueTestInputBuilder;

impl GetAggregateKnownValueTestInputBuilder {
    /// Sends a request with this input using the given client.
    pub async fn send_with(
        self,
        client: &crate::Client,
    ) -> ::std::result::Result<
        crate::operation::get_aggregate_known_value_test::GetAggregateKnownValueTestOutput,
        crate::operation::get_aggregate_known_value_test::GetAggregateKnownValueTestError,
    > {
        let mut fluent_builder = client.get_aggregate_known_value_test();
        fluent_builder.inner = self;
        fluent_builder.send().await
    }
}

/// Fluent builder constructing a request to `GetAggregate`.
///
#[derive(::std::clone::Clone, ::std::fmt::Debug)]
pub struct GetAggregateKnownValueTestFluentBuilder {
    client: crate::Client,
    inner: crate::operation::get_aggregate_known_value_test::builders::GetAggregateKnownValueTestInputBuilder,
}

impl GetAggregateKnownValueTestFluentBuilder {
    /// Creates a new `GetAggregateKnownValueTestFluentBuilder`.
    pub(crate) fn new(client: crate::Client) -> Self {
        Self {
            client,
            inner: ::std::default::Default::default(),
        }
    }
    /// Access the GetAggregateKnownValueTest as a reference.
    pub fn as_input(&self) -> &crate::operation::get_aggregate_known_value_test::builders::GetAggregateKnownValueTestInputBuilder{
        &self.inner
    }
    /// Sends the request and returns the response.
    pub async fn send(
        self,
    ) -> ::std::result::Result<
        crate::operation::get_aggregate_known_value_test::GetAggregateKnownValueTestOutput,
        crate::operation::get_aggregate_known_value_test::GetAggregateKnownValueTestError,
    > {
        let input = self
            .inner
            .build()
            // Using unhandled since GetAggregateKnownValueTest doesn't declare any validation,
            // and smithy-rs seems to not generate a ValidationError case unless there is
            // (but isn't that a backwards compatibility problem for output structures?)
            // Vanilla smithy-rs uses SdkError::construction_failure,
            // but we aren't using SdkError.
            .map_err(crate::operation::get_aggregate_known_value_test::GetAggregateKnownValueTestError::unhandled)?;
        crate::operation::get_aggregate_known_value_test::GetAggregateKnownValueTest::send(
            &self.client,
            input,
        )
        .await
    }

    ///
    /// Appends an item to `simpleStringList`.
    ///
    /// To override the contents of this collection use [`set_simple_string_list`](Self::set_simple_string_list).
    ///
    #[allow(missing_docs)] // documentation missing in model
    pub fn simple_string_list(
        mut self,
        input: impl ::std::convert::Into<::std::string::String>,
    ) -> Self {
        self.inner = self.inner.simple_string_list(input.into());
        self
    }
    #[allow(missing_docs)] // documentation missing in model
    pub fn set_simple_string_list(
        mut self,
        input: ::std::option::Option<::std::vec::Vec<::std::string::String>>,
    ) -> Self {
        self.inner = self.inner.set_simple_string_list(input);
        self
    }
    #[allow(missing_docs)] // documentation missing in model
    pub fn get_simple_string_list(
        &self,
    ) -> &::std::option::Option<::std::vec::Vec<::std::string::String>> {
        self.inner.get_simple_string_list()
    }
    ///
    /// Appends an item to `structureList`.
    ///
    /// To override the contents of this collection use [`set_structure_list`](Self::set_structure_list).
    ///
    #[allow(missing_docs)] // documentation missing in model
    pub fn structure_list(mut self, input: crate::types::StructureListElement) -> Self {
        self.inner = self.inner.structure_list(input);
        self
    }
    #[allow(missing_docs)] // documentation missing in model
    pub fn set_structure_list(
        mut self,
        input: ::std::option::Option<::std::vec::Vec<crate::types::StructureListElement>>,
    ) -> Self {
        self.inner = self.inner.set_structure_list(input);
        self
    }
    #[allow(missing_docs)] // documentation missing in model
    pub fn get_structure_list(
        &self,
    ) -> &::std::option::Option<::std::vec::Vec<crate::types::StructureListElement>> {
        self.inner.get_structure_list()
    }
    ///
    /// Adds a key-value pair to `simpleStringMap`.
    ///
    /// To override the contents of this collection use [`set_simple_string_map`](Self::set_simple_string_map).
    ///
    #[allow(missing_docs)] // documentation missing in model
    pub fn simple_string_map(
        mut self,
        k: impl ::std::convert::Into<::std::string::String>,
        v: impl ::std::convert::Into<::std::string::String>,
    ) -> Self {
        self.inner = self.inner.simple_string_map(k.into(), v.into());
        self
    }
    #[allow(missing_docs)] // documentation missing in model
    pub fn set_simple_string_map(
        mut self,
        input: ::std::option::Option<
            ::std::collections::HashMap<::std::string::String, ::std::string::String>,
        >,
    ) -> Self {
        self.inner = self.inner.set_simple_string_map(input);
        self
    }
    #[allow(missing_docs)] // documentation missing in model
    pub fn get_simple_string_map(
        &self,
    ) -> &::std::option::Option<
        ::std::collections::HashMap<::std::string::String, ::std::string::String>,
    > {
        self.inner.get_simple_string_map()
    }
    ///
    /// Adds a key-value pair to `simpleIntegerMap`.
    ///
    /// To override the contents of this collection use [`set_simple_integer_map`](Self::set_simple_integer_map).
    ///
    #[allow(missing_docs)] // documentation missing in model
    pub fn simple_integer_map(
        mut self,
        k: impl ::std::convert::Into<::std::string::String>,
        v: i32,
    ) -> Self {
        self.inner = self.inner.simple_integer_map(k.into(), v);
        self
    }
    #[allow(missing_docs)] // documentation missing in model
    pub fn set_simple_integer_map(
        mut self,
        input: ::std::option::Option<::std::collections::HashMap<::std::string::String, i32>>,
    ) -> Self {
        self.inner = self.inner.set_simple_integer_map(input);
        self
    }
    #[allow(missing_docs)] // documentation missing in model
    pub fn get_simple_integer_map(
        &self,
    ) -> &::std::option::Option<::std::collections::HashMap<::std::string::String, i32>> {
        self.inner.get_simple_integer_map()
    }
    #[allow(missing_docs)] // documentation missing in model
    pub fn nested_structure(mut self, input: crate::types::NestedStructure) -> Self {
        self.inner = self.inner.nested_structure(input);
        self
    }
    #[allow(missing_docs)] // documentation missing in model
    pub fn set_nested_structure(
        mut self,
        input: ::std::option::Option<crate::types::NestedStructure>,
    ) -> Self {
        self.inner = self.inner.set_nested_structure(input);
        self
    }
    #[allow(missing_docs)] // documentation missing in model
    pub fn get_nested_structure(&self) -> &::std::option::Option<crate::types::NestedStructure> {
        self.inner.get_nested_structure()
    }
}
