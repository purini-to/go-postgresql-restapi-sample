package db

import (
	"fmt"

	"github.com/purini-to/go-postgresql-restapi-sample/context"

	"go.uber.org/zap"

	"github.com/jinzhu/gorm"
	// PostgreSQL driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

type logger struct {
	lg *zap.Logger
}

// ProvideDB provide db connection.
func ProvideDB(cf *viper.Viper, lg *zap.Logger) (*gorm.DB, func(), error) {
	args := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s %s",
		cf.GetString("db.host"),
		cf.GetString("db.user"),
		cf.GetString("db.password"),
		cf.GetString("db.dbname"),
		cf.GetString("db.extension"),
	)
	db, err := gorm.Open("postgres", args)
	if err != nil {
		return nil, nil, err
	}

	if !context.IsProduction() {
		db.LogMode(true)
	}
	return db, func() {
		db.Close()
	}, nil
}
