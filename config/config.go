package config

import (
	"strings"

	"github.com/spf13/viper"
	"gopkg.in/errgo.v2/errors"
)

type Config struct {
	SonosIP string
	StatePath string
}

func setDefaults() {
	viper.SetDefault("StatePath", "/var/local/sonos-standup")
}

func Load(filename *string) (*Config, error) {
	setDefaults()
	viper.SetConfigType("yaml")

	if filename != nil && strings.Compare(*filename, "") != 0 {
		viper.SetConfigFile(*filename)
	} else {
		viper.AddConfigPath("/etc/sonos-standup")
		viper.AddConfigPath("/usr/local/etc/sonos-standup")
		viper.AddConfigPath("./")
		viper.SetConfigName("config")
	}

	err := viper.ReadInConfig()
	if err != nil {
		return nil, errors.Because(err, nil, "Could not read config: ")
	}

	config := Config{}

	err = viper.UnmarshalExact(&config)
	if err != nil {
		return nil, errors.Because(err, nil, "Could not parse config")
	}

	return &config, nil
}