package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Bot         BotConfig               `yaml:"bot"`
	Logger      map[string]LoggerConfig `yaml:"logger"`
	Sqlite      SqliteConfig            `yaml:"sqlite"`
	BloodReport BloodReportConfig       `yaml:"blood_report"`
}

type BloodReportConfig struct {
	BaseURL string `yaml:"base_url" env:"BLOOD_REPORT_URL"`
}

type BotConfig struct {
	Token  string        `yaml:"token" env:"TG_TOKEN"`
	Poller time.Duration `yaml:"poller"`
}

type SqliteConfig struct {
	Path string `yaml:"path"`
}

type LoggerConfig struct {
	Level            string   `yaml:"level"`
	Development      bool     `yaml:"development"`
	Encoding         string   `yaml:"encoding"`
	OutputPaths      []string `yaml:"outputPaths"`
	ErrorOutputPaths []string `yaml:"errorOutputPaths"`
}

// MustLoad loads the config from the given path or the CONFIG_PATH env variable
func MustLoad() Config {
	path := fetchConfigPath()
	if path == "" {
		panic("config path is empty: pass -config flag or set CONFIG_PATH env")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exist: " + path)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return cfg
}

// fetchConfigPath returns the path to the config file, either from the -config flag or the CONFIG_PATH env variable
func fetchConfigPath() string {
	var path string
	flag.StringVar(&path, "config", "", "path to config file")
	flag.Parse()

	if path == "" {
		path = os.Getenv("CONFIG_PATH")
	}

	return path
}
