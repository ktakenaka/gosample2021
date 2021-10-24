package logger

import (
	"net/http"

	"github.com/rollbar/rollbar-go"
)

const (
	// CRIT is the critial severity level.
	CRIT = "critical"
	// ERR is the error severity level.
	ERR = "error"
	// WARN is the warning severity level.
	WARN = "warning"
	// INFO is the info severity level.
	INFO = "info"
	// DEBUG is the debug severity level.
	DEBUG = "debug"
)

type Config struct {
	Token       string
	Environment string
	IsStdLogger bool
}

type Logger interface {
	RequestErrorWithPerson(userID, level string, r *http.Request, err error)
	RequestErrorWithExtrasAndPerson(userID, level string, r *http.Request, err error, extras map[string]interface{})
	Close()
}

type logger struct{}

// TODO: filter params
func New(cfg *Config) Logger {
	rollbar.SetToken(cfg.Token)
	rollbar.SetEnvironment(cfg.Environment)
	return logger{}
}

func (l logger) RequestErrorWithPerson(
	userID string,
	level string,
	r *http.Request,
	err error,
) {
	// TODO: check whether it's goroutine safe
	rollbar.SetPerson(userID, "", "")
	rollbar.RequestError(level, r, err)
}

func (l logger) RequestErrorWithExtrasAndPerson(
	userID string,
	level string,
	r *http.Request,
	err error,
	extras map[string]interface{},
) {
	// TODO: check whether it's goroutine safe
	rollbar.SetPerson(userID, "", "")
	rollbar.RequestErrorWithExtras(level, r, err, extras)
}

func (l logger) Close() {
	rollbar.Close()
}

// Normally you do not want to use these methods
// because you should log request along with error
// func (l logger) Debug(interfaces ...interface{}) {
// 	rollbar.Debug(interfaces...)
// }

// func (l logger) Info(interfaces ...interface{}) {
// 	rollbar.Info(interfaces...)
// }

// func (l logger) Warn(interfaces ...interface{}) {
// 	rollbar.Warning(interfaces...)
// }

// func (l logger) Error(interfaces ...interface{}) {
// 	rollbar.Error(interfaces...)
// }

// func (l logger) Critical(interfaces ...interface{}) {
// 	rollbar.Critical(interfaces...)
// }
