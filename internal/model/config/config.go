package config

import (
	"fmt"

	toml "github.com/pelletier/go-toml/v2"
)

type Config struct {
	Logger Logger `toml:"logger"`
	Server Server `toml:"server"`
}

type Logger struct {
	Debug bool `toml:"debug"`
}

type Server struct {
	Address string `toml:"address"`
}

func DefaultCfg() *Config {
	cfg := new(Config)
	cfg.Logger = Logger{
		Debug: true,
	}

	cfg.Server = Server{
		Address: ":8080",
	}
	return cfg
}

var ErrCfgInvalid = fmt.Errorf("invalid config")

func cfgErr(msg string, args ...interface{}) error {
	return fmt.Errorf("%w: %s", ErrCfgInvalid, fmt.Sprintf(msg, args...))
}

func Validate(cfg *Config) error {
	if cfg == nil {
		return cfgErr("nil cfg")
	}

	return nil
}

func UnmarshalToml(data []byte) (cfg *Config, err error) {
	cfg = new(Config)
	err = toml.Unmarshal(data, cfg)
	return
}

func MarshalToml(cfg *Config) ([]byte, error) {
	return toml.Marshal(cfg)
}
