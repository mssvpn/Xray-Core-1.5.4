package extension

import (
	"context"

	"github.com/golang/protobuf/proto"

	"github.com/mssvpn/Xray-Core-1.5.4/features"
)

type Observatory interface {
	features.Feature

	GetObservation(ctx context.Context) (proto.Message, error)
}

func ObservatoryType() interface{} {
	return (*Observatory)(nil)
}
