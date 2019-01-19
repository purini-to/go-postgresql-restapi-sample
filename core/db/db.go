package db

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/purini-to/go-postgresql-restapi-sample/core/config"

	"github.com/jinzhu/gorm"
	// PostgreSQL driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

// DB is database client.
type DB struct {
	*gorm.DB
}

// ProvideDB provide db connection.
func ProvideDB(cf *config.Config) (*DB, func(), error) {
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		viper.GetString(`db.user`),
		viper.GetString(`db.pass`),
		viper.GetString(`db.host`),
		viper.GetString(`db.port`),
		viper.GetString(`db.name`),
	)
	val := url.Values{}
	val.Add("parseTime", strconv.Itoa(viper.GetInt(`db.parseTime`)))
	val.Add("loc", viper.GetString(`db.loc`))
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())

	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		return nil, nil, err
	}

	if !cf.IsProduction() {
		db.LogMode(true)
	}

	return &DB{
			db,
		}, func() {
			db.Close()
		}, nil
}
