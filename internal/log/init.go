package log

import "custom-geth/internal/db"

func init() {
	if service == nil {
		dbConfig := &db.Config{
			AppName:            "custom-geth",
			Driver:             "postgres",
			DsnMaster:          "host=localhost port=5432 user=postgres password=password dbname=defaultdb_dev sslmode=disable",
			ConnectionLifetime: 10,
			MaxIdleConnections: 100,
			MaxOpenConnections: 100,
		}

		dbManager, err := db.NewManager(dbConfig)
		if err != nil {
			panic("Failed to init LogService. " + err.Error())
		}

		service = &Service{
			dbManager: dbManager,
		}
	}
}
