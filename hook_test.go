package logrushooknsq

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestNewNsqHook(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)
	nsqHook, err := NewNsqHook(logrus.DebugLevel, "127.0.0.1:4150", "nsqhook")
	if err != nil {
		logrus.Error(err)
		return
	}
	logrus.Infof("%+v", nsqHook)
	logrus.AddHook(nsqHook)
	logrus.Debug("my nsq hook")
}
