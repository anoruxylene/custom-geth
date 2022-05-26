package types

import (
	"custom-geth/internal/goeth/common"
)

type Log struct {
	Address     common.Address `json:"address"`
	Topic0      common.Hash    `json:"topic0"`
	Topic1      *common.Hash   `json:"topic1"`
	Topic2      *common.Hash   `json:"topic2"`
	Topic3      *common.Hash   `json:"topic3"`
	Data        common.Hash    `json:"data"`
	BlockNumber uint64         `json:"blockNumber"`
	BlockHash   common.Hash    `json:"blockHash"`
	TxHash      common.Hash    `json:"transactionHash"`
	TxIndex     uint           `json:"transactionIndex"`
	LogIndex    uint           `json:"logIndex"`
	Removed     bool           `json:"removed"`
}

func (l Log) TableName() string {
	return "logs"
}

func (l Log) InsertionFields() []string {
	return []string{
		"topic0",
		"topic1",
		"topic2",
		"topic3",
		"address",
		"data",
		"block_number",
		"block_hash",
		"tx_hash",
		"tx_index",
		"index",
		"removed",
	}
}

func (l *Log) PrepareInsertionFields() []interface{} {
	fields := []interface{}{
		l.Topic0.String(), // Мы должны держать порядок топиков во имя костыля
		nil,               //l.Topic1,
		nil,               //l.Topic2,
		nil,               //l.Topic3,
		l.Address.String(),
		l.Data.String(),
		l.BlockNumber,
		l.BlockHash.String(),
		l.TxHash.String(),
		l.TxIndex,
		l.LogIndex,
		l.Removed,
	}

	// Костыль для заполнения значений не нуловых топиков
	for i, hash := range []*common.Hash{l.Topic1, l.Topic2, l.Topic3} {
		if hash != nil {
			fields[i+1] = hash.String()
		}
	}

	return fields
}
