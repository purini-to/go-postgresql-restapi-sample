package db

import (
	"fmt"

	"github.com/purini-to/go-postgresql-restapi-sample/domain/model"

	"github.com/purini-to/go-postgresql-restapi-sample/core/config"
	"github.com/purini-to/go-postgresql-restapi-sample/core/env"

	"github.com/jinzhu/gorm"
	// PostgreSQL driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB is database client.
type DB struct {
	*gorm.DB
}

// NewDB create db connection.
func NewDB(cf config.Config) (*DB, func(), error) {
	connection := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s %s",
		cf.GetString(`db.host`),
		cf.GetString(`db.port`),
		cf.GetString(`db.user`),
		cf.GetString(`db.name`),
		cf.GetString(`db.pass`),
		cf.GetString(`db.extension`),
	)
	db, err := gorm.Open("postgres", connection)
	if err != nil {
		return nil, nil, err
	}

	if !env.IsProduction() {
		db.LogMode(true)
	}

	db.AutoMigrate(&model.System{})

	return &DB{
			db,
		}, func() {
			db.Close()
		}, nil
}
