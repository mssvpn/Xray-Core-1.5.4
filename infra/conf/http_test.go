package conf_test

import (
	"testing"

	. "github.com/mssvpn/Xray-Core-1.5.4/infra/conf"
	"github.com/mssvpn/Xray-Core-1.5.4/proxy/http"
)

func TestHTTPServerConfig(t *testing.T) {
	creator := func() Buildable {
		return new(HTTPServerConfig)
	}

	runMultiTestCase(t, []TestCase{
		{
			Input: `{
				"timeout": 10,
				"accounts": [
					{
						"user": "my-username",
						"pass": "my-password"
					}
				],
				"allowTransparent": true,
				"userLevel": 1
			}`,
			Parser: loadJSON(creator),
			Output: &http.ServerConfig{
				Accounts: map[string]string{
					"my-username": "my-password",
				},
				AllowTransparent: true,
				UserLevel:        1,
				Timeout:          10,
			},
		},
	})
}
