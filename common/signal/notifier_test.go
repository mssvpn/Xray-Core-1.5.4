package signal_test

import (
	"testing"

	. "github.com/mssvpn/Xray-Core-1.5.4/common/signal"
)

func TestNotifierSignal(t *testing.T) {
	n := NewNotifier()

	w := n.Wait()
	n.Signal()

	select {
	case <-w:
	default:
		t.Fail()
	}
}
