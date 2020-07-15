package xlogger

import (
	"github.com/sirupsen/logrus"
	"testing"
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
