package kafka

import (
	"context"
	"testing"
)

func TestNewProducer(t *testing.T) {
	producer := NewProducer()
	producer.Addrs = []string{""}
	producer.Topic = "topic"
	producer.MaxProcs = 500
	if err := producer.Start(context.TODO()); err != nil {
		t.Log(err)
		return
	}

	for i := 0; i < 1000; i++ {
		producer.WriteOne([]byte("test"))
	}
}
