package config

import (
	"fmt"
	"github.com/danjac/sightings/store"
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("db_name", "sightings")
	viper.SetDefault("db_user", "postgres")
	viper.SetDefault("db_password", "postgres")
	viper.SetDefault("db_host", "127.0.0.1")
	viper.SetDefault("db_sslmode", "disable")
	viper.SetDefault("port", "4000")
}

type AppConfig struct {
	Store store.Store
	Port  string
}

func (cfg *AppConfig) Close() error {
	return cfg.Store.Close()
}

func New() (*AppConfig, error) {

	viper.AutomaticEnv()

	cfg := &AppConfig{}

	connection := fmt.Sprintf("dbname=%s user=%s password=%s host=%s sslmode=%s",
		viper.Get("db_name"),
		viper.Get("db_user"),
		viper.Get("db_password"),
		viper.Get("db_host"),
		viper.Get("db_sslmode"),
	)

	db, err := store.Connect(connection)
	if err != nil {
		return nil, err
	}

	cfg.Store = store.New(db)

	cfg.Port = viper.GetString("port")

	return cfg, nil

}
