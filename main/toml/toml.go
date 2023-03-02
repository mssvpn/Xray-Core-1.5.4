package toml

import (
	"io"

	"github.com/mssvpn/Xray-Core-1.5.4/common"
	"github.com/mssvpn/Xray-Core-1.5.4/common/cmdarg"
	"github.com/mssvpn/Xray-Core-1.5.4/core"
	"github.com/mssvpn/Xray-Core-1.5.4/infra/conf"
	"github.com/mssvpn/Xray-Core-1.5.4/infra/conf/serial"
	"github.com/mssvpn/Xray-Core-1.5.4/main/confloader"
)

func init() {
	common.Must(core.RegisterConfigLoader(&core.ConfigFormat{
		Name:      "TOML",
		Extension: []string{"toml"},
		Loader: func(input interface{}) (*core.Config, error) {
			switch v := input.(type) {
			case cmdarg.Arg:
				cf := &conf.Config{}
				for i, arg := range v {
					newError("Reading config: ", arg).AtInfo().WriteToLog()
					r, err := confloader.LoadConfig(arg)
					if err != nil {
						return nil, newError("failed to read config: ", arg).Base(err)
					}
					c, err := serial.DecodeTOMLConfig(r)
					if err != nil {
						return nil, newError("failed to decode config: ", arg).Base(err)
					}
					if i == 0 {
						// This ensure even if the muti-json parser do not support a setting,
						// It is still respected automatically for the first configure file
						*cf = *c
						continue
					}
					cf.Override(c, arg)
				}
				return cf.Build()
			case io.Reader:
				return serial.LoadTOMLConfig(v)
			default:
				return nil, newError("unknow type")
			}
		},
	}))
}
