package xlogger

import (
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

func TestStdLogger(t *testing.T) {
	Output("test")
	Outputf("a%sc", "b")
	Outputln("test\n")
	Output("test\n")
	Output("test")
}

func TestLogrus(t *testing.T) {
	l := logrus.New()
	l.SetFormatter(&CustomFormatter{ForceColor: true})
	l.Error("test")
	l.Warn("test")
}

func TestRotateFileHook(t *testing.T) {
	l := logrus.New()
	l.SetLevel(logrus.TraceLevel)
	l.AddHook(NewRotateFileHook(&RotateFileConfig{
		Filename:  "./xlogger/logs/file.log",
		Level:     logrus.TraceLevel,
		Formatter: &logrus.JSONFormatter{TimestampFormat: time.RFC3339},
	}))
	l.SetFormatter(&CustomFormatter{ForceColor: true})

	for {
		l.Errorf("test at %s", time.Now().Format(time.RFC3339))
		time.Sleep(time.Second * 5)
	}
}

func TestRotateLogHook(t *testing.T) {
	l := logrus.New()
	l.SetLevel(logrus.TraceLevel)
	l.AddHook(NewRotateLogHook(&RotateLogConfig{
		MaxAge:       15 * 24 * time.Hour,
		RotationTime: 24 * time.Hour,
		LocalTime:    true,
		Filepath:     "./xlogger/logs/",
		Filename:     "console",
		Level:        logrus.TraceLevel,
		Formatter:    &logrus.JSONFormatter{TimestampFormat: time.RFC3339},
	}))
	l.SetFormatter(&CustomFormatter{ForceColor: true})

	for {
		l.Errorf("test at %s", time.Now().Format(time.RFC3339))
		time.Sleep(time.Second * 5)
	}
}
