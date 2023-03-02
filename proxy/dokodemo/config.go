package dokodemo

import (
	"github.com/mssvpn/Xray-Core-1.5.4/common/net"
)

// GetPredefinedAddress returns the defined address from proto config. Null if address is not valid.
func (v *Config) GetPredefinedAddress() net.Address {
	addr := v.Address.AsAddress()
	if addr == nil {
		return nil
	}
	return addr
}
