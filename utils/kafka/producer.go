package kafka

import (
	"encoding/json"
	"errors"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"strings"
	"time"
)

type Data struct {
	Log    string    `json:"log"`
	Stream string    `json:"stream"`
	Time   time.Time `json:"time"`
	Space  string    `json:"space"`
	Tag    string    `json:"tag"`
	Site   string    `json:"site,omitempty"`
}

func CreateProducer(kafkaAddrs []string, kafkaGroup string) *kafka.Producer {
	c, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":  strings.Join(kafkaAddrs, ","),
		"group.id":           kafkaGroup,
		"session.timeout.ms": 6000,
	})
	if err != nil {
		log.Fatal(err)
	}
	return c
}

type DataProducer struct {
	IsOpen   bool
	address  []string
	group    string
	topic    string
	producer *kafka.Producer
}

func (pd *DataProducer) AddMessage(message Data) error {
	bytes, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return pd.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &pd.topic, Partition: kafka.PartitionAny},
		Value:          bytes,
		Headers:        []kafka.Header{},
	}, nil)
}

func (pd *DataProducer) Init(kafkaAddrs []string, kafkaGroup string, topic string) {
	pd.address = kafkaAddrs
	pd.group = kafkaGroup
	pd.topic = topic
	pd.Open()
}

func (pd *DataProducer) Open() error {
	if pd.IsOpen == true {
		return errors.New("Unable to open log consumer, its already open.")
	}
	if pd.address == nil || len(pd.address) == 0 {
		return errors.New("invalid address")
	}
	if pd.group == "" {
		return errors.New("invalid group")
	}
	pd.producer = CreateProducer(pd.address, pd.group)
	pd.IsOpen = true
	return nil
}

func (pd *DataProducer) Close() {
	if pd.IsOpen == false {
		return
	}

	pd.producer.Close()
	pd.IsOpen = false
}
