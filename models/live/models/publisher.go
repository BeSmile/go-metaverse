package models

import (
	"fmt"
	"go-metaverse/models/live/constants"
	"sync"
)

type Event struct {
	name string
	data interface{}
}

type Observer interface {
	notify(Event)
}

type Publisher struct {
	observers map[constants.MsgType][]chan Message
	mu        sync.RWMutex
	handle    func(p *Publisher, message string)
}

func NewPublisher() *Publisher {
	pb := &Publisher{}
	pb.observers = make(map[constants.MsgType][]chan Message)
	return pb
}

func (p *Publisher) AddMessageHandle(handle func(p *Publisher, message string)) {
	p.handle = handle
}

func (p *Publisher) Subscribe(topic constants.MsgType, ch chan Message) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.observers[topic] = append(p.observers[topic], ch)
}

//func (p *Publisher) RemoveObserver(o Observer) {
//	p.mu.Lock()
//	defer p.mu.Unlock()
//	if p.observers != nil {
//		delete(p.observers, o)
//	}
//}

func (p *Publisher) Publish(topic constants.MsgType, msg Message) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	for _, ch := range p.observers[topic] {
		//go o.notify(e)
		ch <- msg
	}
}

func (p *Publisher) AddEventListener(topic constants.MsgType, consumer func(name string, msg Message)) {
	if p.observers[topic] != nil {
		for i, ch := range p.observers[topic] {
			for data := range ch {
				consumer(fmt.Sprintf("channal:%d", i), data)
			}
		}
	}
}

type LoggingObserver struct{}

func (l *LoggingObserver) notify(e Event) {
	fmt.Printf("Logged event %s: %v\\n", e.name, e.data)
}

func (p *Publisher) MessageHandle(message string) {
	if p.handle != nil {
		p.handle(p, message)
	}
}
