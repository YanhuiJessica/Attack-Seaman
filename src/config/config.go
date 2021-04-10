package config

import (
	"io/ioutil"

	"github.com/gotify/configor"
	"gopkg.in/yaml.v2"
)

// Configuration is stuff that can be configured externally per env variables or config file (config.yml).
type Configuration struct {
	Server struct {
		ListenAddr      string `default:""`
		Port            int    `default:"80"`
		ResponseHeaders map[string]string
	}
	Database struct {
		Dbname     string `default:""`
		Connection string `default:""`
		Tbname     string `default:""`
	}
}

// Get returns the configuration extracted from env variables or config file.
func Get() *Configuration {
	conf := new(Configuration)
	err := configor.New(&configor.Config{EnvironmentPrefix: "TenMinutesApi"}).Load(conf, "config.yml")
	if err != nil {
		panic(err)
	}
	return conf
}

func Save(config *Configuration) error {
	var js []byte
	var err error
	js, err = yaml.Marshal(&config)
	if err != nil {
		return nil
	}

	err = ioutil.WriteFile("config.yml", js, 0600)
	return err
}
