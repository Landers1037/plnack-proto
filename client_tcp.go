/*
Project: plnack-proto client_tcp.go
Created: 2021/11/16 by Landers
*/

package plnack_proto

import (
	"net"
)

type PlnackClient struct {
	conn    *net.TCPConn
	Address string
	Port    int
}

// NewPlnackClient 封装的tcp客户端
func NewPlnackClient(address string, port int) *PlnackClient {
	conn, err := net.DialTCP(TCP, nil, &net.TCPAddr{
		IP:   net.ParseIP(address),
		Port: port,
	})

	if err != nil {
		return nil
	}

	return &PlnackClient{
		conn:    conn,
		Address: address,
		Port:    port,
	}
}

func (p *PlnackClient) Send(data PlnackData) error {
	return EncodePlnack(p.conn, data)
}