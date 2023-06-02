package live

import (
	"go-metaverse/models/live/bilibili"
	"go-metaverse/models/live/constants"
	"go-metaverse/models/live/douyu"
	"go-metaverse/models/live/models"
)

type ClientInterface interface {
	SendData(bytesData []byte) error
	JoinRoom(roomId int) error
	Watch()
	HeartBeat()
	Connection() error
	GetPublisher() *models.Publisher
}

type ClientAdapter struct {
	client ClientInterface
}

func NewClient(domain string, port string, platform string) *ClientAdapter {
	var client ClientInterface
	switch platform {
	case constants.DouYu:
		client = douyu.NewClient(domain, port)
		break
	case constants.BiLiBiLi:
		client = bilibili.NewClient(domain, port)
		break

	}
	adapter := &ClientAdapter{
		client,
	}
	return adapter
}

func (c *ClientAdapter) AddEventListener(topic constants.MsgType, consumer func(name string, msg models.Message)) {
	ch := make(chan models.Message, 1)
	c.client.GetPublisher().Subscribe(topic, ch)
	c.client.GetPublisher().AddEventListener(topic, consumer)
}

func (c *ClientAdapter) SendData(bytesData []byte) error {
	err := c.client.SendData(bytesData)
	return err
}

func (c *ClientAdapter) HeartBeat() {
	go c.client.HeartBeat()
}

func (c *ClientAdapter) JoinRoom(roomId int) error {
	return c.client.JoinRoom(roomId)
}

func (c *ClientAdapter) Connection() error {
	err := c.client.Connection()
	return err
}

func (c *ClientAdapter) Watch() {
	c.client.Watch()
}
