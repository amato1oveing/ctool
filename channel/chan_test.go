package channel

import (
	"testing"
)

func TestNewChannel(t *testing.T) {
	channel := NewChannel(200)
	t.Log(channel.String())
	for i := 0; i < 1000; i++ {
		channel.Run(func() error {
			return nil
		})
	}
}
