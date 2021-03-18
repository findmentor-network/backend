package errors

import "fmt"

type ErrorDef struct {
	number   int
	prefix   string
	template string
}

func (definition *ErrorDef) Code() string {
	return fmt.Sprintf("%0s-%06d", definition.prefix, definition.number)
}

func (definition *ErrorDef) Template() string {
	return definition.template
}

func (definition *ErrorDef) Equal(err error) bool {
	switch err.(type) {
	case Error:
		var ourError = err.(Error)

		return ourError.Code() == definition.Code()
	default:
		return false
	}
}
