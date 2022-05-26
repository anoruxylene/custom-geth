package core

import (
	"custom-geth/internal/goeth/core/types"
	logger "custom-geth/internal/goeth/log"
)

func DoLog(block *types.Block, logs []*types.Log) {
	logger.Warn("Received Logs", "blockNumber", block.NumberU64(), "logs", logs, ",")
}
