/*
Project: plnack-proto plnack_input_test.go
Created: 2021/11/3 by Landers
*/

package plnack_proto

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

const (
	INPUT1 = `
	{
		"key": "testkey",
		"version": "1.0",
		"app_name": "plnack",
		"data": "{hello}"
	}
`
)

var dataIn = PlnackInData{
	Key:     "key",
	Version: "1.0",
	AppName: "plnack",
	Data:    "hello",
}

// 用于处理plnack输入的测试
func TestPlnackInJson(t *testing.T) {
	PLNACK_LOG = true
	data, e := DecodeJSONData([]byte(INPUT1))
	if e != nil {
		t.Error(e.Error())
	}
	testSetLoggerWriter()
	t.Logf("%+v", data)
}

// read gob from io
func TestPlnackInReader(t *testing.T) {
	EnableLog()
	EnableVerify()

	w, e := os.OpenFile("test.out", os.O_CREATE|os.O_TRUNC, 0644)
	if e != nil {
		t.Error(e.Error())
	}

	e = EncodeAny(w, PlnackInData{
		Key:     "hello",
		Version: "1.0",
		AppName: "plnack",
		Data:    "{}",
	})
	if e != nil {
		t.Error(e.Error())
	}

	defer os.Remove("test.out")

	file, e := os.Open("test.out")
	if e != nil {
		t.Error(e.Error())
	}

	data, e := DecodeData(file)
	if e != nil {
		t.Error(e.Error())
	}

	t.Logf("%+v", data)
}

// read from gin
func TestPlnackGinReader(t *testing.T) {
	PLNACK_LOG = true
	go func() {
		startGinTest()
	}()

	var bf bytes.Buffer
	enc := gob.NewEncoder(&bf)
	e := enc.Encode(dataIn)
	if e != nil {
		t.Error(e)
	}
	_, e = http.Post("http://localhost:8080/test", "application/octet-stream", &bf)
	if e != nil {
		t.Error(e)
	}
}

func startGinTest() {
	g := gin.Default()
	g.POST("/test", func(c *gin.Context) {
		plnData, e := DecodeGinData(c)
		if e != nil {
			fmt.Printf("%v\n", e)
		}

		fmt.Printf("received: \n%+v\n", plnData)
	})

	g.Run(":8080")
}

func testSetLoggerWriter() {
	logFile, _ := os.Create("test.log")
	SetLoggerWriter(io.MultiWriter(os.Stdout, logFile))
	pl.error("test error\n")
}

// decode any struct
type T1 struct {
	Name string
	Age  int
}

func TestDecodeAny(t *testing.T) {
	var td T1
	w, e := os.OpenFile("test.gob", os.O_CREATE|os.O_TRUNC, 0644)
	if e != nil {
		t.Error(e.Error())
	}

	e = EncodeAny(w, T1{
		Name: "test",
		Age:  1,
	})

	w.Close()
	if e != nil {
		t.Error(e.Error())
	}

	f, e := os.OpenFile("test.gob", os.O_RDONLY, 0644)
	if e != nil {
		t.Error(e)
	}

	DecodeAny(&td, f)
	f.Close()
	defer os.Remove("test.gob")
	t.Log(td)
}

// test any json
func TestDecodeAnyJSON(t *testing.T) {
	j := `{"name": "jiejie"}`
	var td T1
	DecodeAnyJSON(&td, []byte(j))

	t.Log(td)
}
