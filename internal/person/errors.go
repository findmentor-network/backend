package person

import "github.com/findmentor-network/backend/pkg/errors"

const prefix = "Person"

var (
	DatabaseError = errors.DefineError(prefix, 100, "fething data, error :%s")
	NotFoundError = errors.DefineError(prefix, 101, "Person not found")
)

var StatusCodes errors.StatusCodeList = map[string]errors.StatusCode{

	DatabaseError.Code(): {Error: DatabaseError, StatusCode: 500, ErrorCode: 100},
	NotFoundError.Code(): {Error: NotFoundError, StatusCode: 404, ErrorCode: 101},
}
