package log

import (
	"custom-geth/internal/db"
	goethTypes "custom-geth/internal/goeth/core/types"
	logger "custom-geth/internal/goeth/log"
	"github.com/Masterminds/squirrel"
)

var service *Service

func DoLog(block *goethTypes.Block, logs []*goethTypes.Log) {
	logger.Warn("Received Logs", "blockNumber", block.NumberU64(), "logs", logs, ",")
}

type Service struct {
	dbManager *db.Manager
}

func (s Service) StoreLogs(bchLogs []*goethTypes.Log) (err error) {
	if len(bchLogs) == 0 {
		return
	}

	dbLogs := MapBchLogToDbLog(bchLogs)

	conn, err := s.dbManager.GetWriteConnection()
	if err != nil {
		return
	}

	query := squirrel.
		Insert(dbLogs[0].TableName()).
		Columns(dbLogs[0].InsertionFields()...).
		PlaceholderFormat(squirrel.Dollar).
		RunWith(conn)

	for _, log := range dbLogs {
		query = query.Values(log.PrepareInsertionFields()...)
	}

	_, err = query.Exec()

	return
}

func GetService() *Service {
	if service == nil {
		panic("LogService is not initialized")
	}

	return service
}
