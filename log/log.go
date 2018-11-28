package log

import (
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
)

func getLogger(pfx string, enabled bool) *logrus.Entry {
	// Use the standardLogger's format
	logger := logrus.New()
	formatter := logrus.StandardLogger().Formatter
	logger.SetFormatter(formatter)
	if enabled {
		logger.SetOutput(os.Stdout)
	} else {
		logger.SetOutput(ioutil.Discard)
	}
	return logger.WithFields(logrus.Fields{
		"from": "hkapi-debug",
	})
}

var (
	// Debug generates debug lines of output with a "DEBUG" prefix.
	// By default the lines are written to /dev/null.
	Debug = getLogger("DEBUG", false)

	// Info generates debug lines of output with a "INFO" prefix.
	// By default the lines are written to stdout.
	Info = logrus.StandardLogger().WithFields(logrus.Fields{"from": "hkapi"})
)

// Logger is a wrapper for logrus.Logger and provides
// methods to enable and disable logging.
type Logger struct {
	*logrus.Logger
}

// Disable sets the logging output to /dev/null.
func (l *Logger) Disable() {
	l.SetOutput(ioutil.Discard)
}

// Enable sets the logging output to stdout.
func (l *Logger) Enable() {
	l.SetOutput(os.Stdout)
}
