// Package log implements a minimalistic logging package.
package log

import "log"

// Callback describes a basic log callback function.
type Callback func(format string, args ...interface{})

// Logger descripts basic logger interface.
type Logger interface {
	Printf(format string, args ...interface{})
}

// New instantiates the a new basic logger.
func New(name string) Logger {
	logPrefix := "[" + name + "] "
	flags := log.Ldate | log.Ltime | log.Lmicroseconds | log.Lmsgprefix
	return log.New(log.Writer(), logPrefix, flags)
}

// NewCallback creates a new basic log callback.
func NewCallback(name string) Callback {
	return New(name).Printf
}
