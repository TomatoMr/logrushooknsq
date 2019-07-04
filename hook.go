package logrushooknsq

import (
	"encoding/json"
	"github.com/nsqio/go-nsq"
	"github.com/sirupsen/logrus"
)

type Nsq struct {
	addr        string
	topic       string
	nsqProducer *nsq.Producer
}

type NsqHook struct {
	levels []logrus.Level
	nsq    *Nsq
}

func NewNsqHook(level logrus.Level, addr, topic string) (*NsqHook, error) {
	levels := []logrus.Level{}
	for _, l := range []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	} {
		if l <= level {
			levels = append(levels, l)
		}
	}

	nsqProducer, err := nsq.NewProducer(addr, nsq.NewConfig())
	if err != nil {
		return nil, err
	}

	return &NsqHook{
		levels: levels,
		nsq: &Nsq{
			addr:        addr,
			topic:       topic,
			nsqProducer: nsqProducer,
		},
	}, nil
}

func (hook *NsqHook) Fire(entry *logrus.Entry) error {
	msg := entry.Data
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	if err := hook.nsq.nsqProducer.Publish(hook.nsq.topic, data); err != nil {
		return err
	}
	return nil
}

func (hook *NsqHook) Levels() []logrus.Level {
	return hook.levels
}
