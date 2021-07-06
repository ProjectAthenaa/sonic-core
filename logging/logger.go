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
	buffer []byte
}

type LoggerConfig struct {
	AlwaysSubmit bool
	PrintLogs    bool
	Level        LogLevel
	name         string
}

func NewLogger(serviceName string, config ...LoggerConfig) *Log {
	if len(serviceName) == 0 {
		panic(nameCannotBeEmptyError)
	}
	if len(config) == 1 {
		config[0].name = serviceName
		return &Log{Config: config[0]}
	}

	return &Log{Config: LoggerConfig{name: serviceName}}
}

func (l *Log) Println(in ...interface{}) {

}

func (l *Log) Logln(in ...interface{}) {

}

func (l *Log) Log(in ...interface{}) {}
