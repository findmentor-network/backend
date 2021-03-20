package echoextention

import (
	"fmt"
	"github.com/findmentor-network/backend/pkg/errors"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoLog "github.com/labstack/gommon/log"
	"runtime"
)

const UNKNOWN_STATUS_CODE = 500

var UndefinedError = errors.DefineError("Middleware", 14, "UndefinedError: %s")

type Response struct {
	//Code    int    `json:"code"`
	Message string `json:"errorDescription"`
}

var DefaultRecoverConfig = RecoverConfig{
	Skipper:           middleware.DefaultSkipper,
	StackSize:         4 << 10, // 4 KB
	DisableStackAll:   false,
	DisablePrintStack: false,
	LogLevel:          echoLog.INFO,
	statusCodeMapping: errors.StatusCodeList{},
}

type (
	Skipper           func(echo.Context) bool
	GetStatusCodeFunc func() errors.StatusCodeList
	RecoverConfig     struct {
		Skipper           Skipper
		StackSize         int  `yaml:"stack_size"`
		DisableStackAll   bool `yaml:"disable_stack_all"`
		DisablePrintStack bool `yaml:"disable_print_stack"`
		LogLevel          echoLog.Lvl
		statusCodeMapping errors.StatusCodeList
	}
)

func Recover() echo.MiddlewareFunc {
	return RecoverWithConfig(DefaultRecoverConfig)
}
func RecoverWithConfig(config RecoverConfig) echo.MiddlewareFunc {

	if config.Skipper == nil {
		config.Skipper = DefaultRecoverConfig.Skipper
	}
	if config.StackSize == 0 {
		config.StackSize = DefaultRecoverConfig.StackSize
	}
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			if config.Skipper(c) {
				return next(c)
			}

			defer func() {
				if r := recover(); r != nil {

					err, ok := r.(error)
					if !ok {
						err = fmt.Errorf("%v", r)
					}
					stack := make([]byte, config.StackSize)
					length := runtime.Stack(stack, !config.DisableStackAll)

					var correlationId = c.Response().Header().Get(echo.HeaderXRequestID)
					msg := fmt.Sprintf("[PANIC RECOVER] CorrelationId:%s, %v %s\n", correlationId, err, stack[:length])

					switch r.(type) {
					case errors.Error:
						e := err.(errors.Error)
						statusCode, errorCode := getStatusAndErrorCode(e, config.statusCodeMapping)
						//if echoLog.ERROR < config.LogLevel{
						c.Logger().Errorf(msg, err)
						//}
						jsonWithAbort(c, statusCode, errorCode, e.Message())

					default:
						u := errors.New(UndefinedError, err.(error).Error())
						statusCode, errorCode := getStatusAndErrorCode(u, config.statusCodeMapping)

						c.Logger().Errorf(msg, err)
						jsonWithAbort(c, statusCode, errorCode, u.Message())
					}
				}
			}()
			return next(c)
		}
	}

}

func getStatusAndErrorCode(e errors.Error, statusCodes errors.StatusCodeList) (statusCode, errorCode int) {
	if val, ok := statusCodes[e.Code()]; ok {
		return val.StatusCode, val.ErrorCode
	}
	return UNKNOWN_STATUS_CODE, 0
}

func jsonWithAbort(c echo.Context, statusCode, errorCode int, message string) {
	c.JSON(statusCode, Response{message})
}
