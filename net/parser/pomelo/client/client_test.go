package pomeloClient

import (
	"fmt"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	client := New(
		WithRequestTimeout(1 * time.Second),
	)
	client.TagName = "dog egg"
	client.ConnectToWS("127.0.0.1:10010", "")

	defer client.Disconnect()

	time.Sleep(30 * time.Second)

}

func BenchmarkClient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		client := New(
			WithRequestTimeout(1 * time.Second),
		)
		client.TagName = fmt.Sprintf("c-%d", i)
		client.ConnectToWS("172.16.124.137:21000", "")

		client.Disconnect()
	}
}
