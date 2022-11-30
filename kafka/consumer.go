package kafka

import (
	"context"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"strings"
)

// Consumer
// 一些使用说明:
// sarame.OffsetNewest int64 = -1
// sarame.OffsetOldest int64 = -2
type Consumer struct {
	Addrs   []string //如果定义了group,则addrs是zookeeper的地址(2181)，否则的话是kafka的地址(9092)
	Topics  []string
	Group   string
	Offset  string
	Message chan []byte //从这个管道中读取数据
	*kafka.Consumer
}

// NewConsumer new consumer
func NewConsumer(option *ConsumerOption) *Consumer {
	if option == nil {
		option = NewConsumerOption()
	}
	return &Consumer{
		Addrs:  option.Addrs,
		Topics: option.Topics,
		Group:  option.Group,
		Offset: option.Offset,
	}
}

// ConsumerOption consumer option
type ConsumerOption struct {
	Addrs  []string `json:"addrs" yaml:"addrs"`
	Topics []string `json:"topics" yaml:"topics"`
	Group  string   `json:"group" yaml:"group"`
	Offset string   `json:"offset" yaml:"offset"`
}

func NewConsumerOption() *ConsumerOption {
	return &ConsumerOption{
		Offset: "latest",
	}
}

//ReadOne read one message
func (c *Consumer) ReadOne() (msg []byte, err error) {
	msg = <-c.Message
	return msg, nil
}

//ReadN read n message
func (c *Consumer) ReadN(n int) (msg [][]byte, err error) {
	for i := 0; i < n; i++ {
		msg = append(msg, <-c.Message)
	}
	return msg, nil
}

func (c *Consumer) readFromTopics(ctx context.Context) error {
LOOP:
	for {
		select {
		case <-ctx.Done():
			break LOOP
		default:
			msg, err := c.Consumer.ReadMessage(-1)
			if err != nil {
				fmt.Printf("Consumer error: %v (%v)\n", err, msg)
				break LOOP
			}
			c.Message <- msg.Value
		}
	}
	return nil
}

// Start consumer
func (c *Consumer) Start(ctx context.Context) error {
	var err error
	c.Message = make(chan []byte, 1000)

	if c.Consumer, err = kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": strings.Join(c.Addrs, ","),
		"group.id":          c.Group,
		"auto.offset.reset": c.Offset,
	}); err != nil {
		return err
	}
	c.Consumer.SubscribeTopics(c.Topics, nil)
	go c.readFromTopics(ctx)
	return nil
}

// Close consumer
func (c *Consumer) Close() error {
	close(c.Message)
	return c.Consumer.Close()
}
