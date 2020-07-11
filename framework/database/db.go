package database

import (
	"encoder/domain"
	"github.com/jinzhu/gorm"
	"log"
)

type Database struct {
	Db             *gorm.DB
	Dsn            string
	DsnTest        string
	Dbtype         string
	DbtypeTest     string
	Debug          bool
	AutomMigrateDb bool
	Env            string
}

func NewDb() *Database {
	return &Database{}
}

func NewDbTest() *gorm.DB {
	dbInstance := NewDb()
	dbInstance.Env = "Test"
	dbInstance.DbtypeTest = "sqlite3"
	dbInstance.DsnTest = ":memory:"
	dbInstance.AutomMigrateDb = true
	dbInstance.Debug = true

	connection, err := dbInstance.Connect()

	if err != nil {
		log.Fatalf("Test db error: %v", err)
	}

	return connection
}

func (d *Database) Connect() (*gorm.DB, error) {
	var err error

	if d.Env != "test" {
		d.Db, err = gorm.Open(d.Dbtype, d.Dsn)
	} else {
		d.Db, err = gorm.Open(d.DbtypeTest, d.DsnTest)
	}

	if err != nil {
		return nil, err
	}

	if d.Debug {
		d.Db.LogMode(true)
	}

	if d.AutomMigrateDb {
		d.Db.AutoMigrate(&domain.Video{}, &domain.Job{})
	}

	return d.Db, nil
}
