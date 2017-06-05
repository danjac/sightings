package config

import (
	"fmt"
	"github.com/danjac/sightings/repo"
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

const ApiVersion = "v1"

type ApiConfig struct {
	Version string
	Path    string
	Port    string
}

func NewApiConfig(port string) *ApiConfig {
	return &ApiConfig{
		Version: ApiVersion,
		Path:    "/api/" + ApiVersion,
		Port:    port,
	}
}

type Config struct {
	Api  *ApiConfig
	Repo repo.Repo
}

func (cfg *Config) Close() error {
	return cfg.Repo.Close()
}

func Configure() (*Config, error) {

	viper.AutomaticEnv()

	cfg := &Config{}

	connection := fmt.Sprintf("dbname=%s user=%s password=%s host=%s sslmode=%s",
		viper.Get("db_name"),
		viper.Get("db_user"),
		viper.Get("db_password"),
		viper.Get("db_host"),
		viper.Get("db_sslmode"),
	)

	db, err := repo.Connect(connection)
	if err != nil {
		return nil, err
	}

	cfg.Repo = repo.New(db)

	cfg.Api = NewApiConfig(viper.GetString("port"))

	return cfg, nil

}
