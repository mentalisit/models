package models

import "sync"

type BridgeTempMemory struct {
	Timestamp int64          `json:"timestamp"`
	RelayName string         `json:"relayName"`
	MessageDs []MessageDs    `json:"messageDs"`
	MessageTg []MessageTg    `json:"messageTg"`
	Wg        sync.WaitGroup `json:"wg"`
}
type MessageDs struct {
	MessageId string `json:"message_id"`
	ChatId    string `json:"chat_id"`
}
type MessageTg struct {
	MessageId string `json:"message_id"`
	ChatId    string `json:"chat_id"`
}
