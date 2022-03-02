package logger

type Logger interface {
	Log(level Level, kvs ...interface{}) error
}


