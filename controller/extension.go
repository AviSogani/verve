package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"net/http"
	"time"
)

var kafkaWriter *kafka.Writer

func sendToKafka(count int) error {
	payload, err := json.Marshal(map[string]int{"count": count})
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	msg := kafka.Message{
		Key:   []byte(fmt.Sprintf("TS::%d", time.Now().Unix())),
		Value: payload,
	}

	return kafkaWriter.WriteMessages(context.Background(), msg)
}

func sendRequestExtension(endpoint string) {
	req.mu.Lock()

	// count of unique requests in the current minute
	count := len(req.requestIdMap)

	req.mu.Unlock()

	// Create a JSON payload with the unique request count
	payload, err := json.Marshal(map[string]int{"count": count})
	if err != nil {
		log.Printf("Failed to marshal JSON: %v", err)
		return
	}

	resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		log.Printf("Failed to send POST request: %v", err)
		return
	}
	defer resp.Body.Close()

	log.Printf("Sent POST request to %s, got response status: %s", endpoint, resp.Status)
}

// LogRequestExtension logs count of unique received ids into kafka (extension #3)
func LogRequestExtension() {
	req.mu.Lock()

	// Send count to Kafka
	err := sendToKafka(len(req.requestIdMap))
	if err != nil {
		log.Printf("Failed to send to Kafka: %v", err)
	}

	// re-initialise it for the next minute after logging
	req.requestIdMap = make(map[int]int)

	req.mu.Unlock()
}

func ExtensionInit() {
	// Initialize Kafka writer
	kafkaWriter = &kafka.Writer{
		Addr:       kafka.TCP("127.0.0.1:9092"),
		Topic:      "unique_request_count",
		BatchSize:  100,
		Async:      true,
		Completion: nil,
	}
}
