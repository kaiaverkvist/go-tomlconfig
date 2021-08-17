package main

import (
	"github.com/kaiaverkvist/go-tomlconfig/pkg/config"
	"log"
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
		Database: "database",
		Username: "username",
		Password: "password",
		Port:     5432,
	},
}

func main() {
	err := config.LoadOrCreateConfig("example.toml", defaultConfig)
	if err != nil {
		// In this case, you might want to use your default config instead, or keep the application from starting up!
		log.Println("cannot load or create config:", err.Error())
		return
	}

	// Unfortunately a type conversion is required due to the lack of generics in Go at this time.
	conf := config.GetConfig().(*Config)
	log.Println(conf.Database.Database)
}
