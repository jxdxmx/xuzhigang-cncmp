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
	curr, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	ins.AddConfigPath(curr)
	fmt.Println("config_name:", os.Getenv("config_file_name"))
	ins.SetConfigName(os.Getenv("config_file_name"))
	if err = ins.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := ins.UnmarshalKey("basic", &conf); err != nil {
		panic(err)
	}
}
