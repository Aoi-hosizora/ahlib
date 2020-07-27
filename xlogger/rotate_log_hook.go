package xlogger

import (
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"io"
	"path"
	"time"
)

// RotateLogHook's config
type RotateLogConfig struct {
	MaxAge       time.Duration
	RotationTime time.Duration
	LocalTime    bool

	Filepath     string
	Filename     string // without ext
	ForceNewFile bool
	Level        logrus.Level
	Formatter    logrus.Formatter
}

// Write log into files (split logs to files automatically)
type RotateLogHook struct {
	config    *RotateLogConfig
	logWriter io.Writer
}

// noinspection GoUnusedExportedFunction
func NewRotateLogHook(config *RotateLogConfig) logrus.Hook {
	fileName := path.Join(config.Filepath, config.Filename)

	options := []rotatelogs.Option{
		rotatelogs.WithLinkName(fileName),
		rotatelogs.WithMaxAge(config.MaxAge),
		rotatelogs.WithRotationTime(config.RotationTime),
	}
	if config.LocalTime {
		options = append(options, rotatelogs.WithClock(rotatelogs.Local))
	} else {
		options = append(options, rotatelogs.WithClock(rotatelogs.UTC))
	}
	if config.ForceNewFile {
		options = append(options, rotatelogs.ForceNewFile())
	}

	writer, err := rotatelogs.New(fileName+".%Y%m%d.log", options...)
	if err != nil {
		panic(err)
	}

	return &RotateLogHook{
		config:    config,
		logWriter: writer,
	}
}

func (r *RotateLogHook) Levels() []logrus.Level {
	return logrus.AllLevels[:r.config.Level+1]
}

func (r *RotateLogHook) Fire(entry *logrus.Entry) error {
	b, err := r.config.Formatter.Format(entry)
	if err != nil {
		return err
	}
	_, _ = r.logWriter.Write(b) // lock
	return nil
}
