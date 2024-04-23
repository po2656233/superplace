package cherryNats

import (
	"time"

	face "github/po2656233/superplace/facade"
)

var (
	instance = &Conn{
		running: false,
	}
)

func SetInstance(conn *Conn) {
	instance = conn
}

func NewFromConfig(config face.ProfileJSON) *Conn {
	conn := New()
	conn.address = config.GetString("address")
	conn.maxReconnects = config.GetInt("max_reconnects")
	conn.reconnectDelay = config.GetDuration("reconnect_delay", 1) * time.Second
	conn.requestTimeout = config.GetDuration("request_timeout", 1) * time.Second
	conn.user = config.GetString("user")
	conn.password = config.GetString("password")

	if conn.address == "" {
		panic("address is empty!")
	}

	return conn
}

func Get() *Conn {
	return instance
}
