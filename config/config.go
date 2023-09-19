package config

import (
	"github.com/spf13/viper"
	"log"
)

func getViper() *viper.Viper{
	viper := viper.New()
	viper.SetConfigFile("/root/src/go/gopl.io/config/config.yaml")
	if err := viper.ReadInConfig() ; err != nil {
		log.Println(err)
	}
	return viper
}


type NetCat1 struct {
	IP string
	Port string
}

func GetNetCat1() *NetCat1{

	viper := getViper()

	return &NetCat1{
		IP:   viper.GetString("netcat1.ip"),
		Port: viper.GetString("netcat1.port"),
	}
}
