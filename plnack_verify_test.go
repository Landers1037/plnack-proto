/*
Project: plnack-proto plnack_verify_test.go
Created: 2021/11/14 by Landers
*/

package plnack_proto

import "testing"

// test for verify

func TestVerifyPlnackIn(t *testing.T) {
	pi := &PlnackInData{
		Key:     "",
		Version: "",
		AppName: "",
		Data:    nil,
	}

	err := VerifyPlnackIn(pi)
	t.Log(err)
}
