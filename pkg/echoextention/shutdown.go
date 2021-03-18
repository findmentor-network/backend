package echoextention

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Shutdown shuts down the given HTTP server gracefully when receiving an os.Interrupt or syscall.SIGTERM signal.
// It will wait for the specified timeout to stop hanging HTTP handlers.
func Shutdown(instance *echo.Echo, timeout time.Duration) {
	stop := make(chan os.Signal, 1)

	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	log.Infof("shutting down server with %s timeout", timeout)

	if err := instance.Shutdown(ctx); err != nil {
		log.Errorf("error while shutting down server: %v", err)
	} else {
		log.Info("server was shut down gracefully")
	}
}
