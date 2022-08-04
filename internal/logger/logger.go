package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

func DefaultLogger() *logrus.Entry {
	log := logrus.New()
	log.SetOutput(os.Stdout)
	log.SetFormatter(&logrus.JSONFormatter{})

	return logrus.NewEntry(log)
}
