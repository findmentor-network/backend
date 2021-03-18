package errors

import "fmt"

type Error interface {
	Error() string
	Code() string
	Number() int
	Message() string
	InnerCause() string
	StackTrace() string
}

type errorImpl struct {
	definition ErrorDef
	innerCause string
	params     []interface{}
	stackFrame string
}

func (err *errorImpl) Error() string {
	innerCause := err.InnerCause()
	//params := err.Params()
	stackTrace := err.StackTrace()

	errorMessage := fmt.Sprintf("%s", err.Message())
	//errorMessage := fmt.Sprintf("%s - %s", err.Code(), err.Message())

	if innerCause != "" {
		errorMessage += fmt.Sprintf(" - %s", innerCause)
	}

	//if params != "" {
	//	errorMessage += fmt.Sprintf(" - %s", params)
	//}

	if stackTrace != "" {
		errorMessage += fmt.Sprintf(" - %s", stackTrace)
	}

	return errorMessage
}

func (err *errorImpl) Code() string {
	return err.definition.Code()
}
func (err *errorImpl) Number() int {
	return err.definition.number
}
func (err *errorImpl) Message() string {
	return fmt.Sprintf(err.definition.Template(), err.params...)
}

func (err *errorImpl) InnerCause() string {
	return err.innerCause
}

func (err *errorImpl) StackTrace() string {
	return fmt.Sprintf(err.stackFrame)
}

func (err *errorImpl) Params() string {
	return fmt.Sprintf("%v", err.params)
}
