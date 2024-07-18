// Code generated by smithy-go-codegen DO NOT EDIT.

package simpleconstraintstypes

import (
	"fmt"
)

type CollectionOfErrors struct {
    SimpleConstraintsBaseException
    ListOfErrors []error
    Message string
}

func (e CollectionOfErrors) Error() string {
	return fmt.Sprintf("message: %s\n err %v", e.Message, e.ListOfErrors)
}

type OpaqueError struct {
   SimpleConstraintsBaseException
   ErrObject interface{}
}

func (e OpaqueError) Error() string {
   return fmt.Sprintf("message: %v", e.ErrObject )
}
