package tagged

import (
	"context"

	"github.com/mssvpn/Xray-Core-1.5.4/common/net"
)

type DialFunc func(ctx context.Context, dest net.Destination, tag string) (net.Conn, error)

var Dialer DialFunc
