//go:build !json
// +build !json

/*
Project: plnack-proto plnack_gob.go
Created: 2021/10/25 by Landers
*/

package plnack_proto

import (
	"encoding/gob"
	"os"

	jsoniter "github.com/json-iterator/go"
)

var pl *logger
var json = jsoniter.ConfigCompatibleWithStandardLibrary

// 注册全局的默认gob类型
func init() {
	pl = newLogger(os.Stdout)
	gob.Register(map[string]string{})
	gob.Register(map[string]map[string]interface{}{})
}

// RegStruct 注册自定义的结构体
func RegStruct(s interface{}) {
	gob.Register(s)
}

func RegStructs(ss ...interface{}) {
	for s := range ss {
		gob.Register(s)
	}
}
