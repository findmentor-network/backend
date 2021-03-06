package echoextention

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoLog "github.com/labstack/gommon/log"
	"strings"
)

func RegisterGlobalMiddlewares(e *echo.Echo, f GetStatusCodeFunc) {

	e.Use(middleware.BodyLimit("1M"))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://findmentor.network"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: Myskipper,
		Level:   -1,
	}))
	e.Use(HookGateLoggerWithConfig(GateLoggerConfig{
		IncludeRequestBodies:  true,
		IncludeResponseBodies: true,
		Skipper:               Myskipper,
	}))
	e.Use(RecoverWithConfig(RecoverConfig{
		Skipper:           Myskipper,
		StackSize:         4 << 10,
		DisableStackAll:   false,
		DisablePrintStack: false,
		LogLevel:          echoLog.INFO,
		statusCodeMapping: f(),
	}))
}

func Myskipper(context echo.Context) bool {
	if strings.HasPrefix(context.Path(), "/status") ||
		strings.HasPrefix(context.Path(), "/swagger") ||
		strings.HasPrefix(context.Path(), "/metrics") {
		return true
	}

	return false
}
