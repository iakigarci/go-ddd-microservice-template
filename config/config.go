package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type LogLevel string

const (
	Info  LogLevel = "info"
	Debug LogLevel = "debug"
	Trace LogLevel = "trace"
	None  LogLevel = "none"
)

type Password string

func (p Password) MarshalText() ([]byte, error) {
	return []byte("*************"), nil
}

type Config struct {
	App      AppConfig      `mapstructure:"app"`
	HTTP     HTTPConfig     `mapstructure:"http"`
	Postgres DatabaseConfig `mapstructure:"postgres"`
	Logging  LogConfig      `mapstructure:"logging"`
}

type AppConfig struct {
	Name        string `mapstructure:"name"`
	Environment string `mapstructure:"environment"`
	Version     string `mapstructure:"version"`
}

type HTTPConfig struct {
	Host           string   `mapstructure:"host"`
	Port           int      `mapstructure:"port"`
	AllowedOrigins []string `mapstructure:"allowed_origins"`
	Timeout        int      `mapstructure:"timeout"`
}

type DatabaseConfig struct {
	Driver   string   `mapstructure:"driver"`
	Host     string   `mapstructure:"host"`
	Port     int      `mapstructure:"port"`
	Name     string   `mapstructure:"name"`
	User     string   `mapstructure:"user"`
	Password Password `mapstructure:"password"`
	SSLMode  string   `mapstructure:"ssl_mode"`
	PoolMax  int      `mapstructure:"pool_max"`
}

type LogConfig struct {
	Level  LogLevel `mapstructure:"level"`
	Format string   `mapstructure:"format"`
}

func readConfig[E any](configFilePath string) (*E, error) {
	vp := viper.New()
	vp.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	vp.AutomaticEnv()
	vp.SetConfigFile(configFilePath)

	var config E
	err := vp.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal config: %w", err)
	}

	return &config, nil
}

func LoadConfig[E any]() (*E, error) {
	var err error
	var conf *E

	conf, err = readConfig[E]("/etc/myorbik/config.yml")
	if err != nil {
		return nil, err
	}

	return conf, nil
}
