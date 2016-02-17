/*
Package goalogrus contains an adapter that makes it possible to configure goa so it uses logrus
as logger backend.
Usage:

    logger := logrus.New()
    // Initialize logger handler using logrus package
    goa.Log = goalogrus.New(logger)
    // ... Proceed with configuring and starting the goa service
*/
package goalogrus

import (
	"github.com/Sirupsen/logrus"
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
)

// Logger is the logrus goa adapter logger.
type Logger struct {
	*logrus.Logger
}

// New wraps a logrus logger into a goa logger.
func New(logger *logrus.Logger) goa.Logger {
	return &Logger{Logger: logger}
}

// Info logs informational messages using logrus.
func (l *Logger) Info(ctx context.Context, msg string, data ...goa.KV) {
	l.Logger.WithFields(data2rus(data)).Info(msg)
}

// Error logs error messages using logrus.
func (l *Logger) Error(ctx context.Context, msg string, data ...goa.KV) {
	l.Logger.WithFields(data2rus(data)).Error(msg)
}

func data2rus(data []goa.KV) logrus.Fields {
	res := make(logrus.Fields, len(data))
	for _, d := range data {
		res[d.Key] = d.Value
	}
	return res
}
