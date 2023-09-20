package conf

import (
	"gopkg.in/ini.v1"
	"log"
)

var (
	Cfg      *ini.File
	RunMode  string
	HttpPort string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	LoadServer()
	LoadMysql()
	LoadRedis()
}

func LoadServer() {
	RunMode = Cfg.Section("").Key("RUN_MODE").String()
	HttpPort = Cfg.Section("server").Key("HttpPort").String()
}

func LoadMysql() {

}

func LoadRedis() {

}
