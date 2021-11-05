/*
Project: plnack-proto plnack_input_test.go
Created: 2021/11/3 by Landers
*/

package plnack_proto

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"testing"
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
	t.Logf("%+v", data)
}

// read gob from io
func TestPlnackInReader(t *testing.T) {
	PLNACK_LOG = true
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