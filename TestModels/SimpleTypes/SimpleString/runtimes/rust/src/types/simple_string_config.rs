// Code generated by software.amazon.smithy.rust.codegen.smithy-rs. DO NOT EDIT.

#[derive(::std::clone::Clone, ::std::fmt::Debug)]
pub struct SimpleStringConfig {}

impl SimpleStringConfig {
    pub fn builder() -> SimpleStringConfigBuilder {
        SimpleStringConfigBuilder::new()
    }
}

#[derive(::std::clone::Clone, ::std::fmt::Debug)]
pub struct SimpleStringConfigBuilder {}

impl SimpleStringConfigBuilder {
    /// Creates a new `SimpleStringConfigBuilder`.
    pub(crate) fn new() -> Self {
        Self {}
    }
    pub fn build(
        self,
    ) -> ::std::result::Result<SimpleStringConfig, ::aws_smithy_types::error::operation::BuildError>
    {
        ::std::result::Result::Ok(SimpleStringConfig {})
    }
}
