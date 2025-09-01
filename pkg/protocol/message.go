package protocol

type Message struct {
	Topic     string `json:"topic"`
	Partition int32  `json:"partition"`
	Offset    int64  `json:"offset"`
	Key       string `json:"key,omitempty"`
	Value     string `json:"value"` // TODO: Phase2에서 []byte로 변경
	Timestamp int64  `json:"timestamp"`
	Size      int32  `json:"size"`
}
