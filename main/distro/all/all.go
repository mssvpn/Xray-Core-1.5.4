package all

import (
	// The following are necessary as they register handlers in their init functions.

	// Mandatory features. Can't remove unless there are replacements.
	_ "github.com/mssvpn/Xray-Core-1.5.4/app/dispatcher"
	_ "github.com/mssvpn/Xray-Core-1.5.4/app/proxyman/inbound"
	_ "github.com/mssvpn/Xray-Core-1.5.4/app/proxyman/outbound"

	// Default commander and all its services. This is an optional feature.
	_ "github.com/mssvpn/Xray-Core-1.5.4/app/commander"
	_ "github.com/mssvpn/Xray-Core-1.5.4/app/log/command"
	_ "github.com/mssvpn/Xray-Core-1.5.4/app/proxyman/command"
	_ "github.com/mssvpn/Xray-Core-1.5.4/app/stats/command"

	// Developer preview services
	_ "github.com/mssvpn/Xray-Core-1.5.4/app/observatory/command"

	// Other optional features.
	_ "github.com/mssvpn/Xray-Core-1.5.4/app/dns"
	_ "github.com/mssvpn/Xray-Core-1.5.4/app/dns/fakedns"
	_ "github.com/mssvpn/Xray-Core-1.5.4/app/log"
	_ "github.com/mssvpn/Xray-Core-1.5.4/app/policy"
	_ "github.com/mssvpn/Xray-Core-1.5.4/app/reverse"
	_ "github.com/mssvpn/Xray-Core-1.5.4/app/router"
	_ "github.com/mssvpn/Xray-Core-1.5.4/app/stats"

	// Fix dependency cycle caused by core import in internet package
	_ "github.com/mssvpn/Xray-Core-1.5.4/transport/internet/tagged/taggedimpl"

	// Developer preview features
	_ "github.com/mssvpn/Xray-Core-1.5.4/app/observatory"

	// Inbound and outbound proxies.
	_ "github.com/mssvpn/Xray-Core-1.5.4/proxy/blackhole"
	_ "github.com/mssvpn/Xray-Core-1.5.4/proxy/dns"
	_ "github.com/mssvpn/Xray-Core-1.5.4/proxy/dokodemo"
	_ "github.com/mssvpn/Xray-Core-1.5.4/proxy/freedom"
	_ "github.com/mssvpn/Xray-Core-1.5.4/proxy/http"
	_ "github.com/mssvpn/Xray-Core-1.5.4/proxy/loopback"
	_ "github.com/mssvpn/Xray-Core-1.5.4/proxy/mtproto"
	_ "github.com/mssvpn/Xray-Core-1.5.4/proxy/shadowsocks"
	_ "github.com/mssvpn/Xray-Core-1.5.4/proxy/socks"
	_ "github.com/mssvpn/Xray-Core-1.5.4/proxy/trojan"
	_ "github.com/mssvpn/Xray-Core-1.5.4/proxy/vless/inbound"
	_ "github.com/mssvpn/Xray-Core-1.5.4/proxy/vless/outbound"
	_ "github.com/mssvpn/Xray-Core-1.5.4/proxy/vmess/inbound"
	_ "github.com/mssvpn/Xray-Core-1.5.4/proxy/vmess/outbound"

	// Transports
	_ "github.com/mssvpn/Xray-Core-1.5.4/transport/internet/domainsocket"
	_ "github.com/mssvpn/Xray-Core-1.5.4/transport/internet/grpc"
	_ "github.com/mssvpn/Xray-Core-1.5.4/transport/internet/http"
	_ "github.com/mssvpn/Xray-Core-1.5.4/transport/internet/kcp"
	_ "github.com/mssvpn/Xray-Core-1.5.4/transport/internet/quic"
	_ "github.com/mssvpn/Xray-Core-1.5.4/transport/internet/tcp"
	_ "github.com/mssvpn/Xray-Core-1.5.4/transport/internet/tls"
	_ "github.com/mssvpn/Xray-Core-1.5.4/transport/internet/udp"
	_ "github.com/mssvpn/Xray-Core-1.5.4/transport/internet/websocket"
	_ "github.com/mssvpn/Xray-Core-1.5.4/transport/internet/xtls"

	// Transport headers
	_ "github.com/mssvpn/Xray-Core-1.5.4/transport/internet/headers/http"
	_ "github.com/mssvpn/Xray-Core-1.5.4/transport/internet/headers/noop"
	_ "github.com/mssvpn/Xray-Core-1.5.4/transport/internet/headers/srtp"
	_ "github.com/mssvpn/Xray-Core-1.5.4/transport/internet/headers/tls"
	_ "github.com/mssvpn/Xray-Core-1.5.4/transport/internet/headers/utp"
	_ "github.com/mssvpn/Xray-Core-1.5.4/transport/internet/headers/wechat"
	_ "github.com/mssvpn/Xray-Core-1.5.4/transport/internet/headers/wireguard"

	// JSON & TOML & YAML
	_ "github.com/mssvpn/Xray-Core-1.5.4/main/json"
	_ "github.com/mssvpn/Xray-Core-1.5.4/main/toml"
	_ "github.com/mssvpn/Xray-Core-1.5.4/main/yaml"

	// Load config from file or http(s)
	_ "github.com/mssvpn/Xray-Core-1.5.4/main/confloader/external"

	// Commands
	_ "github.com/mssvpn/Xray-Core-1.5.4/main/commands/all"
)
