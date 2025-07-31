package protocol_test

import (
	"encoding/json"
	"github.com/akitsuki-labs/gafka/pkg/protocol"
	"testing"
	"time"
)

func TestMessageSerialization(t *testing.T) {
	msg := protocol.Message{
		Topic:     "test-topic",
		Partition: 0,
		Offset:    42,
		Key:       "user-42",
		Value:     "hello",
		Timestamp: time.Now().Unix(),
		Size:      5,
	}

	data, err := json.Marshal(msg)
	if err != nil {
		t.Fatal(err)
	}

	var decoded protocol.Message
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatal(err)
	}

	if decoded.Value != msg.Value {
		t.Errorf("expected value %s, got %s", msg.Value, decoded.Value)
	}
}
