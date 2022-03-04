package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var conf Config

type Config struct {
	Bid     string `mapstructure:"bid"`
	StaffId string `mapstructure:"staff_id"`
	Mode    string `mapstructure:"mode"`
	Num     int    `mapstructure:"num"`
	Threads int    `mapstructure:"threads"`
	Retry   int    `mapstructure:"retry"`
}

var ins *viper.Viper

func initConfig() {
	ins = viper.New()
	ins.SetConfigType("yaml")
	ins.AddConfigPath("config")
	fmt.Println("config_file_name:", os.Getenv("CONFIG_FILE_NAME"))
	ins.SetConfigName(os.Getenv("config_file_name"))
	if err := ins.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := ins.UnmarshalKey("basic", &conf); err != nil {
		panic(err)
	}
}
