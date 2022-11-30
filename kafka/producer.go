package kafka

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/amato1oveing/ctool/channel"
	"log"
)

//Producer kafka producer
type Producer struct {
	Addrs    []string         `json:"addrs" yaml:"addrs"`
	Topic    string           `json:"topic" yaml:"topic"`
	MaxProcs int              `json:"max_procs" yaml:"max_procs"` //最大并发写协程, 由于并发写入topic,写入顺序不可控,想要严格数序的话,maxThreads = 1即可
	Message  chan []byte      `json:"-" yaml:"-"`                 //将数据写入这个管道中
	channel  *channel.Channel //并发写topic的协程控制
}

//NewProducer new producer
func NewProducer() *Producer {
	return new(Producer)
}

//WriteOne write one message
func (p *Producer) WriteOne(msg []byte) (int, error) {
	p.Message <- msg
	return len(msg), nil
}

//WriteN write n message
func (p *Producer) WriteN(msg [][]byte) error {
	for _, m := range msg {
		p.Message <- m
	}
	return nil
}

//Close producer
func (p *Producer) Close() error {
	close(p.Message)
	return p.channel.Close()
}

func (p *Producer) Start(ctx context.Context) error {
	p.Message = make(chan []byte, p.MaxProcs)
	p.channel = channel.NewChannel(p.MaxProcs)
	go p.writeToTopic(ctx)
	return nil
}

func (p *Producer) writeToTopic(ctx context.Context) error {

	config := sarama.NewConfig()
	config.ClientID = "TransportProducer"
	config.Producer.Return.Successes = true
	if err := config.Validate(); err != nil {
		return err
	}

	producer, err := sarama.NewSyncProducer(p.Addrs, config)
	if err != nil {
		return err
	}
	defer producer.Close()

LOOP:
	for {
		select {
		case <-ctx.Done():
			break LOOP
		case message := <-p.Message:
			p.channel.Add()
			go func(message []byte) {
				msg := &sarama.ProducerMessage{
					Topic: p.Topic,
					Value: sarama.ByteEncoder(message),
				}
				if partition, offset, err := producer.SendMessage(msg); err != nil {
					log.Printf("<write to kafka error,partition=%v,offset=%v> %v, %v", partition, offset, err, string(message))
				}
				p.channel.Done()
			}(message)
		}
	}
	return nil
}
