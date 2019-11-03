package lorca

type Logger interface {
	Debug(args ...interface{})
}

type loggerFunc func(args ...interface{})

func (l loggerFunc) Debug(args ...interface{}) {
	l(args...)
}
