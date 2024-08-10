package superConnector

import (
	"net"
	"sync"
	"testing"

	clog "github.com/po2656233/superplace/logger"
)

func TestNewKcpConnector(t *testing.T) {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	kcp := NewKCP(":9071")
	kcp.OnConnect(func(conn net.Conn) {
		clog.Infof("new net.Conn = %s", conn.RemoteAddr())
	})

	kcp.Start()

	wg.Wait()
}
