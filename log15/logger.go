/*
 * Package goalog15 contains an adapter that makes it possible to configure goa so it uses log15
 * as logger backend.
 * Usage:
 *
 *     logger := goalog15.New()
 *     // Initialize logger handler using log15 package
 *     goa.Log = logger
 *     // ... Proceed with configuring and starting the goa service
 */
package goalog15

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"gopkg.in/inconshreveable/log15.v2"
)

// Logger is the log15 goa adapter logger.
type Logger struct {
	log15.Logger
}

// New wraps a log15 logger into a goa logger.
func New(logger log15.Logger) goa.Logger {
	return &Logger{Logger: logger}
}

// Info logs informational messages using log15.
func (l *Logger) Info(ctx context.Context, msg string, data ...goa.KV) {
	l.Logger.Info(msg, data215(data)...)
}

// Error logs error messages using log15.
func (l *Logger) Error(ctx context.Context, msg string, data ...goa.KV) {
	l.Logger.Error(msg, data215(data)...)
}

func data215(data []goa.KV) []interface{} {
	res := make([]interface{}, 2*len(data))
	for i, d := range data {
		res[i*2] = d.Key
		res[i*2+1] = d.Value
	}
	return res
}
