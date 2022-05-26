package db

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"go.elastic.co/apm/module/apmsql"
)

var (
	MasterConnectionNotDefinedError = errors.New("master db connection not defined")
	SlaveConnectionNotDefinedError  = errors.New("slave db connection not defined")
	NotConnectedError               = errors.New("manager not connected to db")
	ManagerAlreadyDefinedError      = errors.New("connection manager already defined. Use GetManager")
	ManagerNotDefinedError          = errors.New("connection manager not defined. Use NewManager")

	managerInstance *Manager = nil
)

type Manager struct {
	masterConnection *sqlx.DB
	slaveConnection  *sqlx.DB
	isConnected      bool
	migrationPath    string
}

func (manager Manager) IsConnected() bool {
	return manager.isConnected
}

func (manager *Manager) CloseConnection() error {
	if !manager.IsConnected() {
		return NotConnectedError
	}

	err := manager.masterConnection.Close()

	if err != nil {
		return err
	}

	err = manager.slaveConnection.Close()

	if err != nil {
		return err
	}

	manager.isConnected = true

	return nil
}

func (manager Manager) GetWriteConnection() (*sqlx.DB, error) {
	if !manager.IsConnected() {
		return new(sqlx.DB), NotConnectedError
	}

	if manager.masterConnection == nil {
		return new(sqlx.DB), MasterConnectionNotDefinedError
	}

	return manager.masterConnection, nil
}

func (manager Manager) GetReadConnection() (*sqlx.DB, error) {
	if !manager.IsConnected() {
		return new(sqlx.DB), NotConnectedError
	}

	if manager.slaveConnection == nil {
		return new(sqlx.DB), SlaveConnectionNotDefinedError
	}

	return manager.slaveConnection, nil
}

func (manager Manager) GetMasterStats() (*sql.DBStats, error) {
	if !manager.IsConnected() {
		return new(sql.DBStats), NotConnectedError
	}

	mStats := manager.masterConnection.DB.Stats()

	return &mStats, nil
}

func (manager Manager) GetSlaveStats() (*sql.DBStats, error) {
	if !manager.IsConnected() {
		return new(sql.DBStats), NotConnectedError
	}

	sStats := manager.slaveConnection.DB.Stats()

	return &sStats, nil
}

func (manager Manager) GetMigrationsPath() string {
	return manager.migrationPath
}

func GetManager() (*Manager, error) {
	if managerInstance == nil {
		return new(Manager), ManagerNotDefinedError
	}

	if !managerInstance.IsConnected() {
		return managerInstance, NotConnectedError
	}

	return managerInstance, nil
}

func NewManager(config *Config) (*Manager, error) {
	if managerInstance != nil {
		return managerInstance, ManagerAlreadyDefinedError
	}

	driver := apmsql.DriverPrefix + config.Driver
	apmDriver := apmsql.Wrap(&pq.Driver{})

	sql.Register(driver, apmDriver)

	masterDB, err := initDB(driver, config)
	if err != nil {
		return new(Manager), err
	}

	slaveDB, err := initDB(driver, config)
	if err != nil {
		return new(Manager), err
	}

	managerInstance = &Manager{
		masterConnection: masterDB,
		slaveConnection:  slaveDB,
		isConnected:      true,
	}

	return managerInstance, nil
}

func initDB(driver string, config *Config) (*sqlx.DB, error) {
	db, err := sqlx.Open(driver, config.DsnMaster)

	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(config.ConnectionLifetime)
	db.SetMaxOpenConns(config.MaxOpenConnections)
	db.SetMaxIdleConns(config.MaxIdleConnections)

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
