/*
Project: plnack-proto plnack_verify.go
Created: 2021/11/10 by Landers
*/

// 内置的基于validate的结构体数据校验

package plnack_proto

import (
	"github.com/gookit/validate"
)

var (
	PlnackVerify = false
)

// Plnack数据的全局规则
func globalValidate(v *validate.Validation) {
	v.AddRule("app_name", "minLen", 4)
	v.AddRule("version", "minLen", 3)
}

// Check 公共的结构体验证配置
func (pd *PlnackData) Check() error {
	if PlnackVerify {
		v := validate.Struct(pd)
		globalValidate(v)
		if v.Validate() {
			// success
			return nil
		}
		pl.error("validate data failed: %s\n", v.Errors.Error())
		return v.Errors
	}

	return nil
}

func (pi *PlnackInData) Check() error {
	if PlnackVerify {
		v := validate.Struct(pi)
		globalValidate(v)
		if v.Validate() {
			// success
			return nil
		}
		pl.error("validate data failed: %s\n", v.Errors.Error())
		return v.Errors
	}

	return nil
}

func (po *PlnackOutData) Check() error {
	if PlnackVerify {
		v := validate.Struct(po)
		globalValidate(v)
		if v.Validate() {
			// success
			return nil
		}
		pl.error("validate data failed: %s\n", v.Errors.Error())
		return v.Errors
	}

	return nil
}

// VerifyPlnackIn 校验plnack结构体
func VerifyPlnackIn(pi *PlnackInData) error {
	v := validate.Struct(pi)
	globalValidate(v)
	if v.Validate() {
		// success
		return nil
	}

	return v.Errors
}

// VerifyPlnackOut 校验输出
func VerifyPlnackOut(po *PlnackOutData) error {
	v := validate.Struct(po)
	globalValidate(v)
	if v.Validate() {
		// success
		return nil
	}

	return v.Errors
}

// VerifyAny 校验任意结构体
// 需要满足validate
func VerifyAny(d interface{}) *validate.Validation {
	v := validate.Struct(d)
	return v
}

// EnableVerify 启动内置的校验
func EnableVerify() {
	PlnackVerify = true
}
