package logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

const (
	INFO = iota
	NOTICE
	WARN
	ERR
	CRIT
	EMERG
)

const (
	fileFlag = os.O_WRONLY | os.O_APPEND | os.O_CREATE
	fileMode = 0666
)

var levelsStr = map[int]string{
	INFO:   "[\x1b[1;42;37mINFO\x1b[0m]",
	NOTICE: "[\x1b[1;42;37mNOTICE\x1b[0m]",
	WARN:   "[\x1b[1;43;37mWARN\x1b[0m]",
	ERR:    "[\x1b[1;41;37mERR\x1b[0m]",
	CRIT:   "[\x1b[1;45;37mCRIT\x1b[0m]",
	EMERG:  "[\x1b[1;45;37mEMERG\x1b[0m]",
}

type InitArgs struct {
	Writer   io.Writer
	Filename string
	Prefix   string
	Flags    int
	LogLevel int
}

type Logger struct {
	Logger   *log.Logger
	LogLevel int
	writer   io.Writer
}

func New(args *InitArgs) (*Logger, error) {
	self := &Logger{
		LogLevel: args.LogLevel,
	}

	writer := args.Writer
	if nil == writer {
		file, err := os.OpenFile(args.Filename, fileFlag, fileMode)
		if nil != err {
			return nil, err
		}
		writer = file
	}

	self.writer = writer
	self.Logger = log.New(writer, args.Prefix, args.Flags)
	return self, nil
}

func (l *Logger) GetWriter() io.Writer {
	return l.writer
}

func (l *Logger) SetWriter(writer io.Writer) {
	l.writer = writer
	l.Logger.SetOutput(writer)
}

func (l *Logger) FileRotate() (*os.File, error) {
	file, ok := l.writer.(*os.File)
	if !ok {
		return nil, nil
	}
	fileNew, err := os.OpenFile(file.Name(), fileFlag, fileMode)
	if nil != err {
		return nil, err
	}
	l.SetWriter(fileNew)
	return file, nil
}

func (l *Logger) log(calldepth int, level int, v []any) {
	if level < l.LogLevel {
		return
	}
	levelStr, ok := levelsStr[level]
	if !ok {
		levelStr = fmt.Sprintf("[level %d]", level)
	}
	s := fmt.Sprintf("%s%s\n", levelStr, fmt.Sprint(v...))
	if nil != l.Logger {
		l.Logger.Output(calldepth, s)
	} else {
		fmt.Print(s)
	}
}

func (l *Logger) logf(calldepth int, level int, format string, v []any) {
	if level < l.LogLevel {
		return
	}
	levelStr, ok := levelsStr[level]
	if !ok {
		levelStr = fmt.Sprintf("[level %d]", level)
	}
	s := fmt.Sprintf("%s%s\n", levelStr, fmt.Sprintf(format, v...))
	if nil != l.Logger {
		l.Logger.Output(calldepth, s)
	} else {
		fmt.Print(s)
	}
}

func (l *Logger) Log(level int, v ...any) {
	l.log(3, level, v)
}

func (l *Logger) Logf(level int, format string, v ...any) {
	l.logf(3, level, format, v)
}

func (l *Logger) Info(v ...any) {
	l.log(3, INFO, v)
}

func (l *Logger) Infof(format string, v ...any) {
	l.logf(3, INFO, format, v)
}

func (l *Logger) Notice(v ...any) {
	l.log(3, NOTICE, v)
}

func (l *Logger) Noticef(format string, v ...any) {
	l.logf(3, NOTICE, format, v)
}

func (l *Logger) Warn(v ...any) {
	l.log(3, WARN, v)
}

func (l *Logger) Warnf(format string, v ...any) {
	l.logf(3, WARN, format, v)
}

func (l *Logger) Err(v ...any) {
	l.log(3, ERR, v)
}

func (l *Logger) Errf(format string, v ...any) {
	l.logf(3, ERR, format, v)
}

func (l *Logger) Crit(v ...any) {
	l.log(3, CRIT, v)
}

func (l *Logger) Critf(format string, v ...any) {
	l.logf(3, CRIT, format, v)
}

func (l *Logger) Emerg(v ...any) {
	l.log(3, EMERG, v)
}

func (l *Logger) Emergf(format string, v ...any) {
	l.logf(3, EMERG, format, v)
}
