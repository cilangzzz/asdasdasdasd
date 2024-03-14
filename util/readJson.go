package util

import (
	"encoding/json"
	"os"
)

type Config struct {
	App struct {
		AppName string `json:"appName"`
	} `json:"app"`
	Server struct {
		Address string `json:"address"`
		Port    string `json:"port"`
	} `json:"server"`
	Mysql struct {
		User         string `json:"user"`
		Password     string `json:"password"`
		Ip           string `json:"ip"`
		Port         string `json:"port"`
		DatabaseName string `json:"databaseName"`
		Charset      string `json:"Charset"`
	} `json:"mysql"`
}

var Cfg *Config

func ReadJson(path string) *Config {
	bytes, err := os.ReadFile(path)

	if err != nil {
		panic("读取配置文件发生错误")
	}
	var config Config
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		panic("解码文件发生错误")
	}
	Cfg = &config
	return &config
}
