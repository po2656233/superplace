package extendConnector

import (
	"net"
	"sync"
	"testing"

	clog "superplace/logger"
)

func TestNewTCPConnector(t *testing.T) {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	tcp := NewTCP(":9071")
	tcp.OnConnect(func(conn net.Conn) {
		clog.Infof("new net.Conn = %s", conn.RemoteAddr())
	})

	tcp.Start()

	wg.Wait()
}
