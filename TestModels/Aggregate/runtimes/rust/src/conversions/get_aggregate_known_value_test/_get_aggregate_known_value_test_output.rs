// Code generated by software.amazon.smithy.rust.codegen.smithy-rs. DO NOT EDIT.
#[allow(dead_code)]

pub fn to_dafny(
    value: crate::operation::get_aggregate_known_value_test::GetAggregateKnownValueTestOutput,
) -> ::std::rc::Rc<
    crate::r#simple::aggregate::internaldafny::types::GetAggregateOutput,
> {
    ::std::rc::Rc::new(crate::r#simple::aggregate::internaldafny::types::GetAggregateOutput::GetAggregateOutput {
        simpleStringList:
        ::std::rc::Rc::new(
            match value.simple_string_list() {
                Some(val) =>
                    crate::r#_Wrappers_Compile::Option::Some {
                        value : ::dafny_runtime::Sequence::from_array(
                            &val.iter().map(|x|
                                dafny_runtime::dafny_runtime_conversions::unicode_chars_false::string_to_dafny_string(&x))
                                .collect()
                        )
                    },
                None => crate::r#_Wrappers_Compile::Option::None{}
            }
        ),

        structureList:
        ::std::rc::Rc::new(
            match value.structure_list() {
                Some(val) =>
                    crate::r#_Wrappers_Compile::Option::Some {
                        value : ::dafny_runtime::Sequence::from_array(
                            &val.iter().map(|x|
                                crate::conversions::structure_list_element::to_dafny(x.clone()))
                                .collect()
                        )
                    },
                None => crate::r#_Wrappers_Compile::Option::None{}
            }
        ),

        simpleStringMap: ::std::rc::Rc::new(match value.simple_string_map() {
            Some(x) => crate::r#_Wrappers_Compile::Option::Some { value :
                ::dafny_runtime::dafny_runtime_conversions::hashmap_to_dafny_map(x,
                    |k| dafny_runtime::dafny_runtime_conversions::unicode_chars_false::string_to_dafny_string(&k),
                    |v| dafny_runtime::dafny_runtime_conversions::unicode_chars_false::string_to_dafny_string(&v),
                )
            },
            None => crate::r#_Wrappers_Compile::Option::None {}
        }),
        simpleIntegerMap: ::std::rc::Rc::new(match value.simple_integer_map() {
            Some(x) => crate::r#_Wrappers_Compile::Option::Some { value :
                ::dafny_runtime::dafny_runtime_conversions::hashmap_to_dafny_map(x,
                    |k| dafny_runtime::dafny_runtime_conversions::unicode_chars_false::string_to_dafny_string(&k),
                    |v| *v,
                )
            },
            None => crate::r#_Wrappers_Compile::Option::None {}
        }),
        nestedStructure: crate::conversions::nested_structure::option_to_dafny(value.nested_structure())
    })
}

// _get_aggregate_known_value_test_output.rs

#[allow(dead_code)]
pub fn from_dafny(
    dafny_value: ::std::rc::Rc<
        crate::r#simple::aggregate::internaldafny::types::GetAggregateOutput,
    >,
) -> crate::operation::get_aggregate_known_value_test::GetAggregateKnownValueTestOutput {
    match &*dafny_value {
        crate::r#simple::aggregate::internaldafny::types::GetAggregateOutput::GetAggregateOutput {
            simpleStringList,
            structureList,
            simpleStringMap,
            simpleIntegerMap,
            nestedStructure,
        } =>
        crate::operation::get_aggregate_known_value_test::GetAggregateKnownValueTestOutput {
            simple_string_list : match  (&**simpleStringList).as_ref() {
                crate::r#_Wrappers_Compile::Option::Some { value } =>
                    Some(
                        ::dafny_runtime::dafny_runtime_conversions::dafny_sequence_to_vec(&value, |x|
                            dafny_runtime::dafny_runtime_conversions::unicode_chars_false::dafny_string_to_string(&x))
                    ),
                _ => None
            },
            structure_list : match (&**structureList).as_ref() {
                crate::r#_Wrappers_Compile::Option::Some { value } =>
                    Some(
                        ::dafny_runtime::dafny_runtime_conversions::dafny_sequence_to_vec(&value, |x|
                            crate::conversions::structure_list_element::plain_from_dafny(&x))
                    ),
                _ => None
            },

            simple_string_map : match (&**simpleStringMap).as_ref() {
                crate::r#_Wrappers_Compile::Option::Some { value } =>
                    Some(
                        ::dafny_runtime::dafny_runtime_conversions::dafny_map_to_hashmap(&value,
                            |k| dafny_runtime::dafny_runtime_conversions::unicode_chars_false::dafny_string_to_string(&k),
                            |v| dafny_runtime::dafny_runtime_conversions::unicode_chars_false::dafny_string_to_string(&v),
                        )
                    ),
                _ => None
            },

            simple_integer_map : match (&**simpleIntegerMap).as_ref() {
                crate::r#_Wrappers_Compile::Option::Some { value } =>
                    Some(
                        ::dafny_runtime::dafny_runtime_conversions::dafny_map_to_hashmap(&value,
                            |k| dafny_runtime::dafny_runtime_conversions::unicode_chars_false::dafny_string_to_string(&k),
                            |v| *v,
                        )
                    ),
                _ => None
            },

            nested_structure: crate::conversions::nested_structure::option_from_dafny(nestedStructure.clone())
        }
    }
}
