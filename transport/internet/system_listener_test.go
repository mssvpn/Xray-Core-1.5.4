package internet_test

import (
	"context"
	"net"
	"testing"

	"github.com/mssvpn/Xray-Core-1.5.4/common"
	"github.com/mssvpn/Xray-Core-1.5.4/transport/internet"
)

func TestRegisterListenerController(t *testing.T) {
	var gotFd uintptr

	common.Must(internet.RegisterListenerController(func(network string, addr string, fd uintptr) error {
		gotFd = fd
		return nil
	}))

	conn, err := internet.ListenSystemPacket(context.Background(), &net.UDPAddr{
		IP: net.IPv4zero,
	}, nil)
	common.Must(err)
	common.Must(conn.Close())

	if gotFd == 0 {
		t.Error("expected none-zero fd, but actually 0")
	}
}
