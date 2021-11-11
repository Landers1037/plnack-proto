/*
Project: plnack-proto doc.go
Created: 2021/11/11 by Landers

plnack-proto for go-service and other json-api app
使用方式：
go get github.com/landers1037/plnack-proto
编译选项：
使用两种json解析库json-iterator 和 encoding/json
指定-tag json以使用encoding/json
默认使用json-iterator

格式化输入：
PlnackInData 标准的plnack输入结构体
type PlnackInData struct {
	Key     string      `json:"key" validate:"required"`
	Version string      `json:"version" validate:"required"`
	AppName string      `json:"app_name" validate:"required"`
	Data    interface{} `json:"data" validate:"required"`
}

对于标准的plnack应用 必须具备通信密钥Key plnack通信版本 服务名称AppName 通信数据

- json-go
DecodeData(r io.Reader) (PlnackInData, error)
w为标准的io输入

DecodeGinData(c *gin.Context) (PlnackInData, error)
使用gin的上下文读取标准的gin.context

DecodeJSONData(j []byte) (PlnackInData, error)
读取外部的json字节流

自定义输入结构体类型
DecodeAny(model interface{}, r io.Reader) error
传入自定义的结构体指针
示例：
type T1 struct{
	Name string
	Age  int
}
f, e := os.OpenFile("test.gob", os.O_RDONLY, 0644)
if e != nil {
	t.Error(e)
}
DecodeAny(&td, f)

格式化输出：
type PlnackOutData struct {
	KeyVerify bool        `json:"key_verify" validate:"required"`
	Version   string      `json:"version" validate:"required"`
	AppName   string      `json:"app_name" validate:"required"`
	Data      interface{} `json:"data" validate:"required"`
	Time      time.Time   `json:"time" validate:"required"`
}

标准的输出数据必须包含 密钥验证结果KeyVerify 通信版本Version 服务名称AppName 数据Data 返回时间Time

EncodeData(w io.Writer, data interface{}) error
向标准输出中写入数据

EncodeGinData(c *gin.Context, data interface{}) error
向gin的上下文响应中写入

EncodeJSONData(data PlnackOutData) ([]byte, error)
响应为json字节流输出
*/

package plnack_proto
