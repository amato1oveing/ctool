package channel

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Channel struct {
	mux     *sync.Mutex
	channel chan interface{} //使用chan堵塞控制队列的最大数量
	cnt     int64            //通过队列数量计数
}

func NewChannel(max int) *Channel {
	channel := &Channel{
		mux:     new(sync.Mutex),
		channel: make(chan interface{}, max),
		cnt:     0,
	}
	return channel
}

func (ch *Channel) Close() error {
	for i := 0; i < 10; i++ {
		if len(ch.channel) > 0 {
			time.Sleep(100 * time.Millisecond)
			continue
		} else {
			close(ch.channel)
			return nil
		}
	}
	return fmt.Errorf("closed ch failed! ch is not empty, len is %d", ch.Len())
}

//计数
func (ch *Channel) count() {
	ch.mux.Lock()
	ch.cnt = ch.cnt + 1
	ch.mux.Unlock()
}

//往管道中写数据
func (ch *Channel) put(v interface{}) error {
	ch.channel <- v
	ch.count()
	return nil
}

//从管道中读数据
func (ch *Channel) get() (interface{}, bool) {
	v, ok := <-ch.channel
	return v, ok
}

// Add 往管道中放入一个标记，记录活跃数值
func (ch *Channel) Add() {
	ch.put(struct{}{})
}

// Done 从管道中取出一个标记，减少活跃数值
func (ch *Channel) Done() {
	ch.get()
}

// Total 返回总数
func (ch *Channel) Total() int64 { return ch.cnt }

// Cap 返回管道容量
func (ch *Channel) Cap() int { return cap(ch.channel) }

// Len 返回管道中的数据量
func (ch *Channel) Len() int { return len(ch.channel) }

// Idle 返回空闲数值
func (ch *Channel) Idle() int { return ch.Cap() - ch.Len() }

func (ch *Channel) String() string {
	return fmt.Sprintf("<Max:%d,Total:%d,Idle:%d,Len:%d>", ch.Cap(), ch.Total(), ch.Idle(), ch.Len())
}

func (ch *Channel) Run(fun func() error) {
	ch.Add()
	go func() {
		defer ch.Done()
		defer func() {
			if err := recover(); err != nil {
				log.Printf("run func panic: %v", err)
			}
		}()
		if err := fun(); err != nil {
			log.Printf("run func error: %v", err)
		}
	}()
}
