/*
Project: plnack-proto plnack_input.go
Created: 2021/10/25 by Landers
*/

package plnack_proto

import (
	"encoding/gob"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
)

// PlnackInData 标准的plnack输入结构体
type PlnackInData struct {
	Key     string      `json:"key" validate:"required"`
	Version string      `json:"version" validate:"required"`
	AppName string      `json:"app_name" validate:"required"`
	Data    interface{} `json:"data" validate:"required"`
}

// DecodeData 从标准输入输出读取数据
func DecodeData(r io.Reader) (PlnackInData, error) {
	var res PlnackInData
	d := gob.NewDecoder(r)
	e := d.Decode(&res)
	if e != nil {
		pl.error("failed to decode reader %v\n", r)
		return PlnackInData{}, e
	}
	pl.info("decode reader successful\n")
	return res, nil
}

// DecodeGinData 从gin 上下文读取数据
func DecodeGinData(c *gin.Context) (PlnackInData, error) {
	var res PlnackInData
	d := gob.NewDecoder(c.Request.Body)
	e := d.Decode(&res)
	if e != nil {
		pl.error("failed to decode gin body %v\n", c.Request.Body)
		return PlnackInData{}, e
	}
	pl.info("decode gin context successful\n")
	return res, nil
}

// DecodeJSONData 从非gob的json数据读取
// 直接接受字节流
func DecodeJSONData(j []byte) (PlnackInData, error) {
	var res PlnackInData
	e := json.Unmarshal(j, &res)
	if e != nil {
		pl.error("failed to decode json bytes %v\n", j)
		return PlnackInData{}, e
	}
	pl.info("decode json data successful %s\n", j)
	return res, nil
}