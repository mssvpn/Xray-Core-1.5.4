package task

import "github.com/mssvpn/Xray-Core-1.5.4/common"

// Close returns a func() that closes v.
func Close(v interface{}) func() error {
	return func() error {
		return common.Close(v)
	}
}
