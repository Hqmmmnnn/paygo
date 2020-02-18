package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	dsn  string
	port string
}

func InitConfig(cfgPath string) (*Config, error) {
	viper.SetConfigFile(cfgPath)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	dbUser := viper.GetString("postgres.auth")
	dbPwd := viper.GetString("postgres.password")
	dbHost := viper.GetString("postgres.host")
	dbPort := viper.GetInt("postgres.port")
	dbName := viper.GetString("postgres.name")

	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%v/%s?sslmode=disable", dbUser, dbPwd, dbHost, dbPort, dbName)
	port := viper.GetString("port")

	return &Config{dsn: dsn, port: port}, nil
}

func (cfg *Config) GetDsn() string {
	return cfg.dsn
}

func (cfg *Config) GetPort() string {
	return cfg.port
}
