package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func Warn(data any) {
	log.Out = os.Stdout
	log.Warn(data)
}

func Info(data any) {
	log.Out = os.Stdout
	log.Info(data)
}