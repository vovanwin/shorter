package helper

import (
	"encoding/json"
	"github.com/vovanwin/shorter/internal/app/model"
	"os"
)

type Producer struct {
	file    *os.File
	encoder *json.Encoder
}

func NewProducer(fileName string) (*Producer, error) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return nil, err
	}
	return &Producer{
		file:    file,
		encoder: json.NewEncoder(file),
	}, nil
}

func (p *Producer) WriteEvent(event *[]model.URLLink) error {
	return p.encoder.Encode(&event)
}

func (p *Producer) Close() error {
	return p.file.Close()
}

type Consumer struct {
	file    *os.File
	decoder *json.Decoder
}

func NewConsumer(fileName string) (*Consumer, error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY|os.O_CREATE, 0777)
	if err != nil {
		return nil, err
	}
	return &Consumer{
		file:    file,
		decoder: json.NewDecoder(file),
	}, nil
}
func (c *Consumer) ReadEvent() ([]model.URLLink, error) {
	event := []model.URLLink{}
	if err := c.decoder.Decode(&event); err != nil {
		return nil, err
	}
	return event, nil
}
func (c *Consumer) Close() error {
	return c.file.Close()
}
