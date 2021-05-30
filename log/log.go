package log

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var (
	errorLog = log.New(os.Stdout, "\033[31m[error]\033[0m ", log.LstdFlags|log.Lshortfile)
	infoLog  = log.New(os.Stdout, "\033[34m[info ]\033[0m ", log.LstdFlags|log.Lshortfile)
	loggers  = []*log.Logger{errorLog, infoLog}
	mu       sync.Mutex
)

// log的一些方法
var (
	Error  = errorLog.Print
	Errorf = errorLog.Printf
	Info   = infoLog.Print
	Infof  = infoLog.Printf
)

//配置log的分级
const (
	InfoLevel = iota
	ErrorLevel
	Disable
)

func SetLevel(level int)  {
	mu.Lock()
	defer mu.Unlock()

	for _,logger := range loggers {
		logger.SetOutput(os.Stdout)
	}

	if ErrorLevel < level {
		errorLog.SetOutput(ioutil.Discard)
	}

	if InfoLevel < level {
		errorLog.SetOutput(ioutil.Discard)
	}
}


