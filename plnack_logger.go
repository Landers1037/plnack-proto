/*
Project: plnack-proto plnack_logger.go
Created: 2021/10/25 by Landers
*/

package plnack_proto

import (
	"fmt"
	"io"
	"os"
	"sync"
)

// 可以选择是否激活日志
// 日志记录器指针在init中激活 根据全局变量或者环境变量启用

// PLNACK_LOG 控制是否打印日志
var PLNACK_LOG bool

type logger struct {
	// multi writer use io.MultiWriter()
	Writer io.Writer
}

var syonce sync.Once

func active() {
	syonce.Do(func() {
		if PLNACK_LOG || os.Getenv("PLNACK_LOG") != "" && os.Getenv("PLNACK_LOG") == "yes" {
			fmt.Printf("[plnack] logger enabled\n")
		} else {
			fmt.Printf("[plnack] logger disabled\n")
		}
	})
}

func (l *logger) error(f string, args ...interface{}) {
	active()
	if PLNACK_LOG || os.Getenv("PLNACK_LOG") != "" && os.Getenv("PLNACK_LOG") == "yes" {
		fmt.Fprintf(l.Writer, fmt.Sprintf("[ERROR] [plnack] %s", f), args...)
	}
}

func (l *logger) info(f string, args ...interface{}) {
	active()
	if PLNACK_LOG || os.Getenv("PLNACK_LOG") != "" && os.Getenv("PLNACK_LOG") == "yes" {
		fmt.Fprintf(l.Writer, fmt.Sprintf("[INFO] [plnack] %s", f), args...)
	}
}

func newLogger(w io.Writer) *logger {
	l := new(logger)
	l.Writer = w

	return l
}

func SetLoggerWriter(w io.Writer) {
	mux := sync.Mutex{}
	mux.Lock()
	pl.Writer = w

	defer mux.Unlock()
}

func EnableLog() {
	PLNACK_LOG = true
}
