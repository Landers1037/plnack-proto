/*
Project: plnack-proto server_tcp.go
Created: 2021/11/16 by Landers
*/

package plnack_proto

import (
	"net"
)

const (
	TCP = "tcp4"
)

type PlnackServer struct {
	listener *net.TCPListener
	conn     *net.TCPConn
	Address  string
	Port     int
}

// NewPlnackServer 封装的tcp服务端
func NewPlnackServer(address string, port int) *PlnackServer {
	l, err := net.ListenTCP(TCP, &net.TCPAddr{
		IP:   net.ParseIP(address),
		Port: port,
	})

	if err != nil {
		return nil
	}
	return &PlnackServer{
		listener: l,
		Address:  address,
		Port:     port,
	}
}

func (p *PlnackServer) Run() error {
	conn, err := p.listener.AcceptTCP()
	p.conn = conn

	return err
}

func (p *PlnackServer) Close() error {
	return p.listener.Close()
}

func (p *PlnackServer) RemoteIP() net.Addr {
	return p.conn.RemoteAddr()
}

func (p *PlnackServer) Receive() (PlnackData, error) {
	return DecodePlnack(p.conn)
}