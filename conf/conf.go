package conf

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

var cfg *ini.File

func readConfFile() {
	var err error
	dir := os.Getenv("GOPATH")
	env := os.Getenv("GOENV")
	ci := os.Getenv("CI_ENV")
	var cfgFile string

	if ci == "CIRCLECI" {
		dir = "/go"
	}

	if env == "docker" {
		cfgFile = fmt.Sprintf("%s/src/github.com/PhillBeck/ingresse-backend/conf/conf.docker.ini", dir)
	} else {
		cfgFile = fmt.Sprintf("%s/src/github.com/PhillBeck/ingresse-backend/conf/conf.ini", dir)
	}

	cfg, err = ini.Load(cfgFile)
	if err != nil {
		fmt.Printf("Failed to read conf file: %s", err.Error())
		os.Exit(1)
	}
}

func GetMongoURI() string {
	if cfg == nil {
		readConfFile()
	}

	host := cfg.Section("").Key("dbhost").String()
	port := cfg.Section("").Key("dbport").String()

	return fmt.Sprintf("%s:%s", host, port)
}

func GetMongoDatabaseName() string {
	if cfg == nil {
		readConfFile()
	}

	return cfg.Section("").Key("dbname").String()
}

func GetHttpPort() int {
	if cfg == nil {
		readConfFile()
	}

	port, err := cfg.Section("").Key("httpport").Int()

	if err != nil {
		return 5000
	}

	return port
}

func GetDocsHTTPPort() int {
	if cfg == nil {
		readConfFile()
	}

	port, err := cfg.Section("").Key("docsport").Int()

	if err != nil {
		return 8000
	}

	return port
}
