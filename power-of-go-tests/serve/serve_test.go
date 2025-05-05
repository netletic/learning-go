package serve_test

import (
	"net"
	"testing"
	"time"

	"github.com/netletic/serve"
)

func randomLocalAddr(t *testing.T) string {
	t.Helper()
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()
	return l.Addr().String()
}

func waitForServer(t *testing.T, addr string) {
	t.Helper()
	timeout := time.NewTimer(100 * time.Millisecond)
	defer timeout.Stop()
	_, err := net.Dial("tcp", addr)
	for err != nil {
		select {
		case <-timeout.C:
			t.Fatal("timed out")
		default:
			t.Log("retrying")
			time.Sleep(time.Millisecond)
			_, err = net.Dial("tcp", addr)
		}
	}
}

func TestListenAsync_ListensOnGivenAddr(t *testing.T) {
	t.Parallel()
	addr := randomLocalAddr(t)
	serve.ListenAsync(addr)
	waitForServer(t, addr)
}
