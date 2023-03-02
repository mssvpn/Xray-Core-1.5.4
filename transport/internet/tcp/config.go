package tcp

import (
	"github.com/mssvpn/Xray-Core-1.5.4/common"
	"github.com/mssvpn/Xray-Core-1.5.4/transport/internet"
)

const protocolName = "tcp"

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
