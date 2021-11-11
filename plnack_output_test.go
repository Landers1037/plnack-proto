/*
Project: plnack-proto plnack_output_test.go
Created: 2021/11/3 by Landers
*/

package plnack_proto

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"
)

var dataOut = PlnackOutData{
	KeyVerify: false,
	Version:   "1.0",
	AppName:   "plnack",
	Data:      map[string]string{"name": "landers"},
	Time:      time.Now(),
}

// 输出gob测试
func TestPlnackOut(t *testing.T) {
	w, e := os.OpenFile("test.out", os.O_CREATE|os.O_TRUNC, 0644)
	if e != nil {
		t.Error(e.Error())
	}

	e = EncodeData(w, dataOut)
	if e != nil {
		t.Error(e.Error())
	}
}

// 输出gin字节流测试
func TestGinOutPut(t *testing.T) {
	PLNACK_LOG = true
	go func(data PlnackOutData) {
		startGinWriter(data)
	}(dataOut)

	res, e := http.Post("http://localhost:8080/test", "application/octet-stream", strings.NewReader(""))
	if e != nil {
		t.Error(e)
	}

	//b, e := ioutil.ReadAll(res.Body)
	//if e != nil {
	//	t.Error(e)
	//}
	//
	//t.Logf("%+v\n", b)
	// fmt to plnack data
	planckD, e := DecodeData(res.Body)
	t.Logf("response: \n%+v\n", planckD)
	defer res.Body.Close()
}

func startGinWriter(data PlnackOutData) {
	fmt.Println(data)
	g := gin.Default()
	g.POST("/test", func(c *gin.Context) {
		e := EncodeGinData(c, data)
		if e != nil {
			fmt.Printf("%v\n", e)
			return
		}

		return
	})

	g.Run(":8080")
}
