package drain

import "io"

//go:generate go run github.com/mssvpn/Xray-Core-1.5.4/common/errors/errorgen

type Drainer interface {
	AcknowledgeReceive(size int)
	Drain(reader io.Reader) error
}
