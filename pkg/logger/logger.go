package logger

import (
	"log"
	"os"
)

type Logger struct {
	infoL *log.Logger
	warnL *log.Logger
	errL  *log.Logger
}

var instance *Logger

func GetLogger() Logger {
	if instance == nil {
		flags := log.Ltime | log.Lshortfile
		instance = &Logger{
			infoL: log.New(os.Stdout, "#INFO#", flags),
			warnL: log.New(os.Stdout, "[WARN]", flags),
			errL:  log.New(os.Stderr, "!ERROR", flags),
		}
		instance.Info("Logger created")
	}
	return *instance
}

func (l *Logger) Info(v ...interface{}) {
	l.infoL.Println(v...)
}

func (l *Logger) Infof(s string, v ...interface{}) {
	l.infoL.Printf(s, v...)
}

func (l *Logger) Warn(v ...interface{}) {
	l.warnL.Println(v...)
}

func (l *Logger) Warnf(s string, v ...interface{}) {
	l.warnL.Printf(s, v...)
}

func (l *Logger) Err(v ...interface{}) {
	l.errL.Println(v...)
}

func (l *Logger) Errf(s string, v ...interface{}) {
	l.errL.Printf(s, v...)
}

func (l *Logger) Fatal(v ...interface{}) {
	l.errL.Fatal(v...)
}

func (l *Logger) Fatalf(s string, v ...interface{}) {
	l.errL.Fatalf(s, v...)
}
