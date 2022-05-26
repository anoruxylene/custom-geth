package core

import (
	"custom-geth/internal/goethereum/core/types"
	logger "custom-geth/internal/goethereum/log"
)

func DoLog(block *types.Block, logs []*types.Log) {
	logger.Warn("Received Logs", "blockNumber", block.NumberU64(), "logs", logs, ",")
}
