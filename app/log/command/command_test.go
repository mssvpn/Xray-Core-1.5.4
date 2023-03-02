package command_test

import (
	"context"
	"testing"

	"github.com/mssvpn/Xray-Core-1.5.4/app/dispatcher"
	"github.com/mssvpn/Xray-Core-1.5.4/app/log"
	. "github.com/mssvpn/Xray-Core-1.5.4/app/log/command"
	"github.com/mssvpn/Xray-Core-1.5.4/app/proxyman"
	_ "github.com/mssvpn/Xray-Core-1.5.4/app/proxyman/inbound"
	_ "github.com/mssvpn/Xray-Core-1.5.4/app/proxyman/outbound"
	"github.com/mssvpn/Xray-Core-1.5.4/common"
	"github.com/mssvpn/Xray-Core-1.5.4/common/serial"
	"github.com/mssvpn/Xray-Core-1.5.4/core"
)

func TestLoggerRestart(t *testing.T) {
	v, err := core.New(&core.Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&log.Config{}),
			serial.ToTypedMessage(&dispatcher.Config{}),
			serial.ToTypedMessage(&proxyman.InboundConfig{}),
			serial.ToTypedMessage(&proxyman.OutboundConfig{}),
		},
	})
	common.Must(err)
	common.Must(v.Start())

	server := &LoggerServer{
		V: v,
	}
	common.Must2(server.RestartLogger(context.Background(), &RestartLoggerRequest{}))
}
