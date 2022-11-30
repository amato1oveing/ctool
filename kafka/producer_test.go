package kafka

import (
	"context"
	"testing"
)

func TestNewProducer(t *testing.T) {
	producer := NewProducer(&ProducerOption{
		Addrs:      []string{""},
		Topic:      "",
		MaxProcs:   100,
		ChannelMax: 100,
	})
	if err := producer.Start(context.TODO()); err != nil {
		t.Log(err)
		return
	}

	for i := 0; i < 1000; i++ {
		producer.WriteOne([]byte("test"))
	}
}
