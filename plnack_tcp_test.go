/*
Project: plnack-proto plnack_tcp_test.go
Created: 2021/11/16 by Landers
*/

package plnack_proto

import (
	"fmt"
	"net"
	"testing"
	"time"
)

// 基于socket连接的测试
func TestPlnackTCP(t *testing.T) {
	EnableLog()
	EnableVerify()
	go newTcpConn()
	tick := time.Tick(time.Second * 2)
	for now := range tick {
		fmt.Println(now)
		send()
	}
}

func newTcpConn() {
	fmt.Println("new tcp conn")
	netAddr := net.TCPAddr{IP: net.ParseIP("127.0.0.1"), Port: 9000}
	lis, _  := net.ListenTCP("tcp4", &netAddr)
	for {
		conn, err := lis.AcceptTCP()
		if err != nil {
			fmt.Println("create tcp conn failed", err)
		}
		go func() {
			data, err := DecodePlnack(conn)
			if err != nil {
				fmt.Println("decode failed", err)
			} else {
				fmt.Println("data received", data)
			}
		}()
	}

}

func send() {
	conn, err := net.DialTCP("tcp4", nil, &net.TCPAddr{IP: net.ParseIP("127.0.0.1"), Port: 9000})
	if err != nil {
		fmt.Println(err)
	}

	err = EncodePlnack(conn, PlnackData{
		Key:       "plnack_key",
		Type:      PtypeClient,
		Version:   "1.0.0",
		AppName:   "plnack_client",
		Data:      "{}",
		KeyVerify: false,
		Time:      time.Now(),
	})

	if err != nil {
		fmt.Println("send data failed", err)
	} else {
		fmt.Println("data sent")
	}
}