package configs

import (
	"github.com/andreis3/foodtosave-case/internal/util"
	"github.com/spf13/viper"
)

type Conf struct {
	DBDriver            string `mapstructure:"DB_DRIVER"`
	DBHost              string `mapstructure:"DB_HOST"`
	DBPort              string `mapstructure:"DB_PORT"`
	DBUser              string `mapstructure:"DB_USER"`
	DBPassword          string `mapstructure:"DB_PASSWORD"`
	DBName              string `mapstructure:"DB_NAME"`
	ServerPort          string `mapstructure:"SERVER_PORT"`
	MaxConnections      string `mapstructure:"MAX_CONNECTIONS"`
	MinConnections      string `mapstructure:"MIN_CONNECTIONS"`
	MaxConnLifetime     string `mapstructure:"MAX_CONN_LIFETIME"`
	MaxConnIdleTime     string `mapstructure:"MAX_CONN_IDLE_TIME"`
	Topic               string `mapstructure:"TOPIC"`
	TopicDiff           string `mapstructure:"TOPIC_DIFF"`
	Broker              string `mapstructure:"BROKER"`
	Group               string `mapstructure:"GROUP"`
	GroupID             string `mapstructure:"GROUP_ID"`
	NumberBatchMessages string `mapstructure:"NUMBER_BATCH_MESSAGES"`
	NumberWorkers       string `mapstructure:"NUMBER_WORKERS"`
}

func LoadConfig(path string) (*Conf, *util.Error) {
	var cfg *Conf
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		cfg = &Conf{
			DBHost:              viper.GetString("DB_HOST"),
			DBPort:              viper.GetString("DB_PORT"),
			DBUser:              viper.GetString("DB_USER"),
			DBPassword:          viper.GetString("DB_PASSWORD"),
			DBName:              viper.GetString("DB_NAME"),
			ServerPort:          viper.GetString("SERVER_PORT"),
			MaxConnections:      viper.GetString("MAX_CONNECTIONS"),
			MinConnections:      viper.GetString("MIN_CONNECTIONS"),
			MaxConnLifetime:     viper.GetString("MAX_CONN_LIFETIME"),
			MaxConnIdleTime:     viper.GetString("MAX_CONN_IDLE_TIME"),
			Topic:               viper.GetString("TOPIC"),
			TopicDiff:           viper.GetString("TOPIC_DIFF"),
			Broker:              viper.GetString("BROKER"),
			Group:               viper.GetString("GROUP"),
			GroupID:             viper.GetString("GROUP_ID"),
			NumberBatchMessages: viper.GetString("NUMBER_BATCH_MESSAGES"),
			NumberWorkers:       viper.GetString("NUMBER_WORKERS"),
		}
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, &util.Error{
			Code:     "CFG-0001",
			Origin:   "LoadConfig",
			LogError: []string{err.Error()},
		}
	}
	return cfg, nil
}
