package errors

type StatusCode struct {
	Error      ErrorDef
	StatusCode int
	ErrorCode  int
}

type StatusCodeList map[string]StatusCode
