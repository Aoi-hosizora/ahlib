package xlogger

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
)

type RotateFileConfig struct {
	Level     logrus.Level
	Formatter logrus.Formatter
	Filename  string

	// MaxSize is the maximum size in megabytes of the log file before it gets
	// rotated. It defaults to 100 megabytes.
	MaxSize int

	// MaxAge is the maximum number of days to retain old log files based on the
	// timestamp encoded in their filename. The default is not to remove old log
	// files based on age.
	MaxAge int

	// MaxBackups is the maximum number of old log files to retain. The default
	// is to retain all old log files (though MaxAge may still cause them to get
	// deleted.)
	MaxBackups int

	// LocalTime determines if the time used for formatting the timestamps in
	// backup files is the computer's local time. The default is to use UTC
	// time.
	LocalTime bool

	// Compress determines if the rotated log files should be compressed
	// using gzip. The default is not to perform compression.
	Compress bool
}

type RotateFileHook struct {
	config    *RotateFileConfig
	logWriter io.Writer
}

func NewRotateFileHook(config *RotateFileConfig) (logrus.Hook, error) {
	return &RotateFileHook{
		config: config,
		logWriter: &lumberjack.Logger{
			Filename:   config.Filename,
			MaxSize:    config.MaxSize,
			MaxBackups: config.MaxBackups,
			MaxAge:     config.MaxAge,
			LocalTime:  config.LocalTime,
			Compress:   config.Compress,
		},
	}, nil
}

func (r *RotateFileHook) Levels() []logrus.Level {
	return logrus.AllLevels[:r.config.Level+1]
}

func (r *RotateFileHook) Fire(entry *logrus.Entry) (err error) {
	b, err := r.config.Formatter.Format(entry)
	if err != nil {
		return err
	}
	_, _ = r.logWriter.Write(b)
	return nil
}
