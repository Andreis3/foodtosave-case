package configs

import (
	"github.com/spf13/viper"
)

type Conf struct {
	ServerPort              string `mapstructure:"SERVER_PORT"`
	PostgresHost            string `mapstructure:"POSTGRES_HOST"`
	PostgresPort            string `mapstructure:"POSTGRES_PORT"`
	PostgresUser            string `mapstructure:"POSTGRES_USER"`
	PostgresPassword        string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresDBName          string `mapstructure:"POSTGRES_DB_NAME"`
	PostgresMaxConnections  string `mapstructure:"POSTGRES_MAX_CONNECTIONS"`
	PostgresMinConnections  string `mapstructure:"POSTGRES_MIN_CONNECTIONS"`
	PostgresMaxConnLifetime string `mapstructure:"POSTGRES_MAX_CONN_LIFETIME"`
	PostgresMaxConnIdleTime string `mapstructure:"POSTGRES_MAX_CONN_IDLE_TIME"`
	RedisHost               string `mapstructure:"REDIS_HOST"`
	RedisPort               string `mapstructure:"REDIS_PORT"`
	RedisPassword           string `mapstructure:"REDIS_PASSWORD"`
	RedisDB                 string `mapstructure:"REDIS_DB"`
}

func LoadConfig(path string) (*Conf, error) {
	var cfg *Conf
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		cfg = &Conf{
			ServerPort:              viper.GetString("SERVER_PORT"),
			PostgresHost:            viper.GetString("POSTGRES_HOST"),
			PostgresPort:            viper.GetString("POSTGRES_PORT"),
			PostgresUser:            viper.GetString("POSTGRES_USER"),
			PostgresPassword:        viper.GetString("POSTGRES_PASSWORD"),
			PostgresDBName:          viper.GetString("POSTGRES_DB_NAME"),
			PostgresMaxConnections:  viper.GetString("POSTGRES_MAX_CONNECTIONS"),
			PostgresMinConnections:  viper.GetString("POSTGRES_MIN_CONNECTIONS"),
			PostgresMaxConnLifetime: viper.GetString("POSTGRES_MAX_CONN_LIFETIME"),
			PostgresMaxConnIdleTime: viper.GetString("POSTGRES_MAX_CONN_IDLE_TIME"),
			RedisHost:               viper.GetString("REDIS_HOST"),
			RedisPort:               viper.GetString("REDIS_PORT"),
			RedisPassword:           viper.GetString("REDIS_PASSWORD"),
			RedisDB:                 viper.GetString("REDIS_DB"),
		}
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
