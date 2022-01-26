package nodes

import (
	"dev/server"
	"net"
)

type MLclient struct {
	clientId string
	father *server.MLserver
	clientConn *net.Conn
}

func NewMLclient (id string, con *net.Conn, ser *server.MLserver) *MLclient {
	return &MLclient{
		clientId: id,
		father: ser,
		clientConn: con,
	}
}