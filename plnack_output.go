/*
Project: plnack-proto plnack_output.go
Created: 2021/10/25 by Landers
*/

package plnack_proto

import (
	"encoding/gob"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"time"
)

// PlnackOutData 标准的plnack返回值
type PlnackOutData struct {
	KeyVerify bool        `json:"key_verify" validate:"required"`
	Version   string      `json:"version" validate:"required"`
	AppName   string      `json:"app_name" validate:"required"`
	Data      interface{} `json:"data" validate:"required"`
	Time      time.Time   `json:"time" validate:"required"`
}

func EncodeData(w io.Writer, data interface{}) error {
	enc := gob.NewEncoder(w)
	e := enc.Encode(data)
	if e != nil {
		pl.error("failed to encode writer %v\n", e)
		return e
	}
	pl.info("encode writer successful %+v\n", data)
	return nil
}

func EncodeGinData(c *gin.Context, data interface{}) error {
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

func EncodeJSONData(data PlnackOutData) ([]byte, error) {
	b, err := json.Marshal(data)
	if err != nil {
		pl.error("failed to marshal data %v\n", data)
		return nil, err
	}
	pl.info("encode json successful %+v\n", data)
	return b, nil
}
