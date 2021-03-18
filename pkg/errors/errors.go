package errors

import (
	"fmt"
	"runtime"
)

func DefineError(prefix string, number int, template string) ErrorDef {
	return ErrorDef{prefix: prefix, number: number, template: template}
}

func New(definition ErrorDef, params ...interface{}) Error {
	return &errorImpl{definition: definition, params: params}
}

func NewWithCause(definition ErrorDef, innerCause error, params ...interface{}) Error {
	return &errorImpl{definition: definition, innerCause: innerCause.Error(), params: params}
}

func Panic(definition ErrorDef, params ...interface{}) Error {
	panic(&errorImpl{definition: definition, params: params, stackFrame: getStackFrameForPanic()})
}

func PanicWithCause(definition ErrorDef, innerCause error, params ...interface{}) Error {
	panic(&errorImpl{definition: definition, innerCause: innerCause.Error(), params: params, stackFrame: getStackFrameForPanic()})
}

func GetStackFrame(beginningLine, lineCount int) (stackFrame string) {
	pcList := make([]uintptr, lineCount)

	runtime.Callers(beginningLine, pcList)
	for i, pc := range pcList {
		f := runtime.FuncForPC(pc)

		if f != nil {
			file, line := f.FileLine(pc)
			stackFrame += fmt.Sprintf("Level:%d, File:%s, Line:%d, Function:%s\n", i, file, line, f.Name())
		} else {
			stackFrame += fmt.Sprintf("Can not infer stackFrame for level:%d\n", i)
		}
	}

	return
}

func getStackFrameForPanic() string {
	return GetStackFrame(4, 5)
}
