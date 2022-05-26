package log

import (
	"custom-geth/internal/goeth/common"
	goethTypes "custom-geth/internal/goeth/core/types"
	"custom-geth/internal/log/types"
)

func MapBchLogToDbLog(input []*goethTypes.Log) (output []*types.Log) {
	for _, l := range input {
		log := &types.Log{
			Address:     l.Address,
			BlockNumber: l.BlockNumber,
			BlockHash:   l.BlockHash,
			Topic0:      l.Topics[0],
			Data:        common.BytesToHash(l.Data),
			TxHash:      l.TxHash,
			TxIndex:     l.TxIndex,
			LogIndex:    l.Index,
			Removed:     l.Removed,
		}

		if len(l.Topics) == 2 {
			log.Topic1 = &l.Topics[1]
		}

		if len(l.Topics) == 3 {
			log.Topic2 = &l.Topics[2]
		}

		if len(l.Topics) == 4 {
			log.Topic3 = &l.Topics[3]
		}

		output = append(output, log)
	}

	return
}
