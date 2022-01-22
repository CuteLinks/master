package config

import (
	"github.com/tkanos/gonfig"
)

type Configuration struct {
	INMEMORYSTORAGE bool
	PORT string
	DOMAIN string
	DB_CONN_STR string
}

func GetConfig(params ...string) Configuration{
	configuration := Configuration{}
	fileName:="config.json"
	gonfig.GetConf(fileName, &configuration)
	return configuration
}