package config_test

import (
	"fmt"
	"github.com/kaiaverkvist/go-tomlconfig/pkg/config"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"reflect"
	"testing"
	"time"
)

type Config struct {
	Database    Database
	HttpOptions HttpOptions
}

type Database struct {
	Hostname string
	Database string
	Username string
	Password string
	Port     int
}

type HttpOptions struct {
	Host string
	Port int

	UseSSL      bool
	SSLKeyFile  string
	SSLCertFile string

	AllowedHosts []string
}

var defaultConfig = Config{
	HttpOptions: HttpOptions{
		Host:        "127.0.0.1",
		Port:        8081,
		UseSSL:      false,
		SSLKeyFile:  "file.key",
		SSLCertFile: "server.crt",

		AllowedHosts: []string{"localhost", "127.0.0.1"},
	},
	Database: Database{
		Hostname: "hostname",
		Database: "database-test",
		Username: "username",
		Password: "password2",
		Port:     5432,
	},
}

func TestLoadOrCreateConfig(t *testing.T) {
	filePath := fmt.Sprintf("test_%d_config.toml", time.Now().Unix())
	err := config.LoadOrCreateConfig(filePath, defaultConfig)

	var conf *Config
	conf = config.GetConfig().(*Config)

	assert.Nil(t, err, "expected nil err in config load/create")
	assert.NotNil(t, conf, "expected non nil config")
	assert.Equal(t, reflect.TypeOf(&defaultConfig).String(), reflect.TypeOf(conf).String(), "expected same config type between default config and loaded config")
	assert.Equal(t, conf.Database.Database, "database-test")
	assert.Equal(t, conf.Database.Password, "password2")

	log.Println("Removing test config file", filePath)
	_ = os.Remove(filePath)
}
