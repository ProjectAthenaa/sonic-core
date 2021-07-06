package logging

type LogLevel int

const (
	Info LogLevel = iota + 1
	Warning
	Error
)

type Logger interface {
	Println(in ...interface{})
	Logln(in ...interface{})
	Log(in ...interface{})
}

type Log struct {
	Config LoggerConfig
}

type LoggerConfig struct {
	AlwaysSubmit bool
	PrintLogs    bool
	Level        LogLevel
}

func NewLogger(config ...LoggerConfig) Log {
	if len(config) == 1 {
		return Log{Config: config[0]}
	}

	return Log{}
}

func (l *Log) Println(in ...interface{}) {

}

func (l *Log) Logln(in ...interface{}) {

}

func (l *Log) Log(in ...interface{}) {}
