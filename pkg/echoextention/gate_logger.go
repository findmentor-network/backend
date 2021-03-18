package echoextention

import (
	"bufio"
	"bytes"
	"encoding/json"
	"github.com/findmentor-network/backend/pkg/log"
	"github.com/labstack/echo/v4/middleware"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type GateLoggerConfig struct {
	IncludeRequestBodies  bool
	IncludeResponseBodies bool
	Skipper               Skipper
}

var DefaultConfig = GateLoggerConfig{
	IncludeRequestBodies:  false,
	IncludeResponseBodies: false,
	Skipper:               middleware.DefaultSkipper,
}

func logrusMiddlewareHandler(c echo.Context, next echo.HandlerFunc, config GateLoggerConfig) error {

	if config.Skipper != nil && config.Skipper(c) {
		return next(c)
	}

	start := time.Now()

	// Request
	req := c.Request()
	var reqBody []byte
	if config.IncludeRequestBodies {
		if req.Body != nil { // Read
			reqBody, _ = ioutil.ReadAll(req.Body)
		}
		req.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody)) // Reset
	}

	// Response
	res := c.Response()
	resBody := new(bytes.Buffer)
	if config.IncludeResponseBodies {
		mw := io.MultiWriter(res.Writer, resBody)
		writer := &bodyDumpResponseWriter{Writer: mw, ResponseWriter: c.Response().Writer}
		res.Writer = writer
	}

	var err error
	if err = next(c); err != nil {
		c.Error(err)
	}
	stop := time.Now()
	fieldsMap := map[string]interface{}{
		"time_rfc3339":  time.Now().UTC().Format(time.RFC3339),
		"remote_ip":     c.RealIP(),
		"host":          req.Host,
		"uri":           req.RequestURI,
		"method":        req.Method,
		"path":          getPath(req),
		"referer":       req.Referer(),
		"user_agent":    req.UserAgent(),
		"status":        res.Status,
		"latency":       strconv.FormatInt(stop.Sub(start).Nanoseconds()/1000, 10),
		"latency_human": stop.Sub(start).String(),
		"bytes_in":      getBytesIn(req),
		"bytes_out":     strconv.FormatInt(res.Size, 10),
		"request_id":    getRequestID(req, res),
		"error":         err,
	}

	if config.IncludeRequestBodies {

		//fieldsMap["request_body"] = string(reqBody)
		var r interface{}
		json.Unmarshal(reqBody, &r)
		fieldsMap["request_body"] = r
	}

	if config.IncludeResponseBodies {
		var s interface{}
		json.Unmarshal(resBody.Bytes(), &s)
		fieldsMap["response_body"] = s
	}
	log.Logger.WithFields(fieldsMap).Info("Handled request")

	return nil
}

func getBytesIn(req *http.Request) string {
	bytesIn := req.Header.Get(echo.HeaderContentLength)
	if bytesIn == "" {
		bytesIn = "0"
	}
	return bytesIn
}

func getPath(req *http.Request) string {
	p := req.URL.Path
	if p == "" {
		p = "/"
	}
	return p
}

func getRequestID(req *http.Request, res *echo.Response) string {
	var id = req.Header.Get(echo.HeaderXRequestID)
	if id == "" {
		id = res.Header().Get(echo.HeaderXRequestID)
	}
	return id
}

func HookGateLoggerWithConfig(config GateLoggerConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return logrusMiddlewareHandler(c, next, config)
		}
	}
}
func HookGateLogger() echo.MiddlewareFunc {
	return HookGateLoggerWithConfig(DefaultConfig)
}

type bodyDumpResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w *bodyDumpResponseWriter) WriteHeader(code int) {
	w.ResponseWriter.WriteHeader(code)
}

func (w *bodyDumpResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func (w *bodyDumpResponseWriter) Flush() {
	w.ResponseWriter.(http.Flusher).Flush()
}

func (w *bodyDumpResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return w.ResponseWriter.(http.Hijacker).Hijack()
}

func (w *bodyDumpResponseWriter) CloseNotify() <-chan bool {
	return w.ResponseWriter.(http.CloseNotifier).CloseNotify()
}
