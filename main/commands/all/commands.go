package all

import (
	"github.com/mssvpn/Xray-Core-1.5.4/main/commands/all/api"
	"github.com/mssvpn/Xray-Core-1.5.4/main/commands/all/tls"
	"github.com/mssvpn/Xray-Core-1.5.4/main/commands/base"
)

// go:generate go run github.com/mssvpn/Xray-Core-1.5.4/common/errors/errorgen

func init() {
	base.RootCommand.Commands = append(
		base.RootCommand.Commands,
		api.CmdAPI,
		// cmdConvert,
		tls.CmdTLS,
		cmdUUID,
	)
}
