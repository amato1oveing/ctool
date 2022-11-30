package kafka

import (
	"context"
	"testing"
)

func TestNewConsumer(t *testing.T) {
	consumer := NewConsumer(&ConsumerOption{
		Addrs:  []string{""},
		Topics: []string{""},
		Group:  "",
		Offset: "latest",
	})
	if err := consumer.Start(context.TODO()); err != nil {
		panic(err)
	}

	for {
		readOne, err := consumer.ReadOne()
		if err != nil {
			t.Log(err)
			continue
		}
		t.Log(string(readOne))
	}
}
