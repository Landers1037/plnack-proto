//go:build !json
// +build !json

/*
Project: plnack-proto plnack_output.go
Created: 2021/10/25 by Landers
*/

package plnack_proto

import (
	"encoding/gob"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// PlnackOutData 标准的plnack返回值
type PlnackOutData struct {
	KeyVerify bool        `json:"key_verify" validate:"required"`
	Version   string      `json:"version" validate:"required"`
	AppName   string      `json:"app_name" validate:"required"`
	Data      interface{} `json:"data" validate:"required"`
	Time      time.Time   `json:"time" validate:"required"`
}

// EncodeData 向标准IO 输出写入
func EncodeData(w io.Writer, data PlnackOutData) error {
	_ = data.Check()
	enc := gob.NewEncoder(w)
	e := enc.Encode(data)
	if e != nil {
		pl.error("failed to encode writer %v\n", e)
		return e
	}
	pl.info("encode writer successful %+v\n", data)
	return nil
}

// EncodeHTTP 向http标准输出写入
func EncodeHTTP(r http.ResponseWriter, data PlnackOutData) error {
	_ = data.Check()
	enc := gob.NewEncoder(r)
	e := enc.Encode(data)
	if e != nil {
		pl.error("failed to encode http writer %v\n", e)
		return e
	}
	pl.info("encode http writer successful %+v\n", data)
	return nil
}

// EncodeGinData 向gin上下文写入
func EncodeGinData(c *gin.Context, data PlnackOutData) error {
	_ = data.Check()
	c.Header("Content-Type", "application/octet-stream")
	c.Header("App", "Plnack")
	enc := gob.NewEncoder(c.Writer)
	e := enc.Encode(data)
	if e != nil {
		pl.error("failed to encode gin writer %v\n", e)
		return e
	}
	pl.info("encode gin writer successful %+v\n", data)
	return nil
}

func EncodeSocket(conn *net.TCPConn, data PlnackOutData) error {
	_ = data.Check()
	enc := gob.NewEncoder(conn)
	e := enc.Encode(data)
	if e != nil {
		pl.error("failed to encode socket writer %v\n", e)
		return e
	}
	pl.info("encode socket writer successful %+v\n", data)
	return nil
}

// EncodeJSONData 输出为JSON数据
func EncodeJSONData(data PlnackOutData) ([]byte, error) {
	_ = data.Check()
	b, err := json.Marshal(data)
	if err != nil {
		pl.error("failed to marshal data %v\n", data)
		return nil, err
	}
	pl.info("encode json successful %+v\n", data)
	return b, nil
}

// EncodeAny 接受任意数据类型
func EncodeAny(w io.Writer, data interface{}) error {
	enc := gob.NewEncoder(w)
	e := enc.Encode(data)
	if e != nil {
		pl.error("failed to encode writer %v\n", e)
		return e
	}
	pl.info("encode writer successful %+v\n", data)
	return nil
}

func EncodeAnyJSON(data interface{}) ([]byte, error) {
	b, err := json.Marshal(data)
	if err != nil {
		pl.error("failed to marshal data %v\n", data)
		return nil, err
	}
	pl.info("encode json successful %+v\n", data)
	return b, nil
}