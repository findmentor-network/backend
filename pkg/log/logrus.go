package log

import (
	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

// Logrus : implement Logger
type Logrus struct {
	*logrus.Logger
}

// Logger ...
var Logger *logrus.Logger
var llogrus *Logrus

// GetEchoLogger for e.Logger
func SetupLogger() Logrus {

	if Logger == nil {
		//formatter := &logrus.TextFormatter{
		//	FullTimestamp: false,
		//	ForceColors:   true,
		//}
		jsonFormatter := &logrus.JSONFormatter{
			PrettyPrint: true,
		}
		runtimeFormatter := &runtime.Formatter{
			ChildFormatter: jsonFormatter,
			Line:           true,
			Package:        true,
			File:           false,
			BaseNameOnly:   false,
		}
		Logger = logrus.New()
		Logger.SetFormatter(runtimeFormatter)
		Logger.SetOutput(os.Stdout)
		llogrus = &Logrus{Logger}
	}
	return *llogrus
}

// Level returns logger level
func (l Logrus) Level() log.Lvl {
	switch l.Logger.Level {
	case logrus.DebugLevel:
		return log.DEBUG
	case logrus.WarnLevel:
		return log.WARN
	case logrus.ErrorLevel:
		return log.ERROR
	case logrus.InfoLevel:
		return log.INFO
	default:
		l.Panic("Invalid level")
	}

	return log.OFF
}

// SetHeader is a stub to satisfy interface
// It's controlled by Logger
func (l Logrus) SetHeader(_ string) {}

// SetPrefix It's controlled by Logger
func (l Logrus) SetPrefix(s string) {}

// Prefix It's controlled by Logger
func (l Logrus) Prefix() string {
	return ""
}

// SetLevel set level to logger from given log.Lvl
func (l Logrus) SetLevel(lvl log.Lvl) {
	switch lvl {
	case log.DEBUG:
		Logger.SetLevel(logrus.DebugLevel)
	case log.WARN:
		Logger.SetLevel(logrus.WarnLevel)
	case log.ERROR:
		Logger.SetLevel(logrus.ErrorLevel)
	case log.INFO:
		Logger.SetLevel(logrus.InfoLevel)
	default:
		l.Panic("Invalid level")
	}
}

// Output logger output func
func (l Logrus) Output() io.Writer {
	return l.Out
}

// SetOutput change output, default os.Stdout
func (l Logrus) SetOutput(w io.Writer) {
	Logger.SetOutput(w)
}

// Printj print json log
func (l Logrus) Printj(j log.JSON) {
	Logger.WithFields(logrus.Fields(j)).Print()
}

// Debugj debug json log
func (l Logrus) Debugj(j log.JSON) {
	Logger.WithFields(logrus.Fields(j)).Debug()
}

// Infoj info json log
func (l Logrus) Infoj(j log.JSON) {
	Logger.WithFields(logrus.Fields(j)).Info()
}

// Warnj warning json log
func (l Logrus) Warnj(j log.JSON) {
	Logger.WithFields(logrus.Fields(j)).Warn()
}

// Errorj error json log
func (l Logrus) Errorj(j log.JSON) {
	Logger.WithFields(logrus.Fields(j)).Error()
}

// Fatalj fatal json log
func (l Logrus) Fatalj(j log.JSON) {
	Logger.WithFields(logrus.Fields(j)).Fatal()
}

// Panicj panic json log
func (l Logrus) Panicj(j log.JSON) {
	Logger.WithFields(logrus.Fields(j)).Panic()
}

// Print string log
func (l Logrus) Print(i ...interface{}) {
	Logger.Print(i[0].(string))
}

// Debug string log
func (l Logrus) Debug(i ...interface{}) {
	Logger.Debug(i[0].(string))
}

// Info string log
func (l Logrus) Info(i ...interface{}) {
	Logger.Info(i[0].(string))
}

// Warn string log
func (l Logrus) Warn(i ...interface{}) {
	Logger.Warn(i[0].(string))
}

// Error string log
func (l Logrus) Error(i ...interface{}) {
	Logger.Error(i[0].(string))
}

// Fatal string log
func (l Logrus) Fatal(i ...interface{}) {
	Logger.Fatal(i[0].(string))
}

// Panic string log
func (l Logrus) Panic(i ...interface{}) {
	Logger.Panic(i[0].(string))
}
