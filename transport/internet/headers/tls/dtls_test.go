package tls_test

import (
	"context"
	"testing"

	"github.com/mssvpn/Xray-Core-1.5.4/common"
	"github.com/mssvpn/Xray-Core-1.5.4/common/buf"
	. "github.com/mssvpn/Xray-Core-1.5.4/transport/internet/headers/tls"
)

func TestDTLSWrite(t *testing.T) {
	content := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'}
	dtlsRaw, err := New(context.Background(), &PacketConfig{})
	common.Must(err)

	dtls := dtlsRaw.(*DTLS)

	payload := buf.New()
	dtls.Serialize(payload.Extend(dtls.Size()))
	payload.Write(content)

	if payload.Len() != int32(len(content))+dtls.Size() {
		t.Error("payload len: ", payload.Len(), " want ", int32(len(content))+dtls.Size())
	}
}
