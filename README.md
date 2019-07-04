# how to use it
```go
// new a hook
nsqHook, err := NewNsqHook(logrus.Level, "nsq_addr", "nsq_topic")
// add hook
logrus.AddHook(nsqHook)
```