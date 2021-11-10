/*
Project: plnack-proto plnack_logger_test.go
Created: 2021/11/6 by Landers
*/

package plnack_proto

import (
	"io"
	"os"
	"testing"
)

// test for logger
func TestLogger(t *testing.T) {
	EnableLog()
	pl := newLogger(os.Stdout)

	pl.error("error test\n")
	pl.info("info test\n")

	// use file
	logFile, _ := os.Create("test.log")
	pl2 := newLogger(io.MultiWriter(os.Stdout, logFile))
	pl2.error("error test\n")
	pl2.info("info test\n")

	// clean
	logFile.Close()
	os.Remove("test.log")
}
