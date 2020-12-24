package log

import (
	"github.com/bitcubix/go-rest-api-boilerplate/pkg/fs"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

//Logger implements structured and leveled logging
type Logger interface {
	Debug(args ...interface{})
	Debugln(args ...interface{})
	Debugf(msg string, args ...interface{})
	Info(msg string)
	Infoln(...interface{})
	Infof(msg string, args ...interface{})
	Warn(msg string)
	Warnln(...interface{})
	Warnf(msg string, args ...interface{})
	Error(msg string)
	Errorf(msg string, args ...interface{})
	Fatalf(msg string, args ...interface{})
	Print(args ...interface{})
	Println(...interface{})
	Printf(msg string, args ...interface{})
	Trace(args ...interface{})
	Traceln(...interface{})
	Tracef(msg string, args ...interface{})
	Verbose() bool
	WithFields(m map[string]interface{}) Logger
	WithPrefix(prefix string) Logger
}

//Fields own declaration of logrus.Fields
type Fields logrus.Fields

//New returns a implementation of Logger with logrus
func New(w io.Writer, level, dir string) Logger {
	if w == nil {
		w = os.Stderr
	}

	lr := logrus.New()
	lr.Out = w

	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		lvl = logrus.WarnLevel
		lr.Warnf("failed to parse log-level '%s', set default to 'warning'", level)
	}

	lr.SetLevel(lvl)
	lr.SetFormatter(getFormatter(false))

	if dir != "" {
		_ = fs.EnsureDir(dir)
		fileHook, err := NewLogrusFileHook(dir+"/go-rest-api-boilerplate.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
		if err == nil {
			lr.Hooks.Add(fileHook)
		} else {
			lr.Warnf("Failed to open logfile, using standard out: %v", err)
		}
	}

	return &logrusLogger{
		Entry: logrus.NewEntry(lr),
	}
}

type logrusLogger struct {
	*logrus.Entry
}

func (l *logrusLogger) WithFields(fields map[string]interface{}) Logger {
	annotatedEntry := l.Entry.WithFields(fields)
	return &logrusLogger{
		Entry: annotatedEntry,
	}
}

func (l *logrusLogger) Error(msg string) {
	l.Errorf(msg)
}

func (l *logrusLogger) Info(msg string) {
	l.Infof(msg)
}

func (l *logrusLogger) Print(args ...interface{}) {
	l.Debug(args...)
}

func (l *logrusLogger) Warn(msg string) {
	l.Warnf(msg)
}

func (l *logrusLogger) Verbose() bool {
	return l.Entry.Logger.GetLevel().String() == "debug"
}

func (l *logrusLogger) WithPrefix(prefix string) Logger {
	return l.WithFields(Fields{"prefix": prefix})
}

func getFormatter(disableColors bool) *textFormatter {
	return &textFormatter{
		DisableColors:    disableColors,
		ForceFormatting:  true,
		ForceColors:      true,
		DisableTimestamp: false,
		FullTimestamp:    true,
		DisableSorting:   true,
		TimestampFormat:  "2006-01-02 15:04:05.000000",
		SpacePadding:     45,
	}
}
