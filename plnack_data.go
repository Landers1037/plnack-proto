/*
Project: plnack-proto plnack_data.go
Created: 2021/11/16 by Landers
*/

package plnack_proto

import (
	"encoding/gob"
	"io"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	PtypeServer  = "server"
	PtypeClient  = "client"
	PtypeDefault = ""
)

// PlnackData 双向通信的服务协议
type PlnackData struct {
	Key       string      `json:"key" validate:"required"`
	Type      string      `json:"type" validate:"in:server, client,''"`
	Version   string      `json:"version" validate:"required"`
	AppName   string      `json:"app_name" validate:"required"`
	Data      interface{} `json:"data" validate:"required"`
	KeyVerify bool        `json:"key_verify"`
	Time      time.Time   `json:"time" validate:"required"`
}

func EncodePlnack(w io.Writer, data PlnackData) error {
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

func DecodePlnack(r io.Reader) (PlnackData, error) {
	var res PlnackData
	d := gob.NewDecoder(r)
	e := d.Decode(&res)
	_ = res.Check()
	if e != nil {
		pl.error("failed to decode reader %v\n", r)
		return PlnackData{}, e
	}
	pl.info("decode reader successful\n")
	return res, nil
}

func EncodePlnackJson(data PlnackData) ([]byte, error) {
	_ = data.Check()
	b, err := json.Marshal(data)
	if err != nil {
		pl.error("failed to marshal data %v\n", data)
		return nil, err
	}
	pl.info("encode json successful %+v\n", data)
	return b, nil
}

func DecodePlnackJson(j []byte) (PlnackData, error) {
	var res PlnackData
	e := json.Unmarshal(j, &res)
	_ = res.Check()
	if e != nil {
		pl.error("failed to decode json bytes %v\n", j)
		return PlnackData{}, e
	}
	pl.info("decode json data successful %s\n", j)
	return res, nil
}

// 对gin上下文的封装

// EncodeGin 直接将数据定向输出到gin的上下文
func EncodeGin(c *gin.Context, data PlnackData) error {
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

// DecodeGin 直接从gin的body中解码数据
func DecodeGin(c *gin.Context) (PlnackData, error) {
	var res PlnackData
	d := gob.NewDecoder(c.Request.Body)
	e := d.Decode(&res)
	_ = res.Check()
	if e != nil {
		pl.error("failed to decode gin body %v\n", c.Request.Body)
		return PlnackData{}, e
	}
	pl.info("decode gin context successful\n")
	return res, nil
}