package config

import (
	"fmt"
	"gopkg.in/ini.v1"
)

type ConfigObject struct {
	Name           string
	Host           string
	Port           int
	MaxConn        int
	MaxTranSize    int
	version        string
	MsgQueueNum    int
	MsgQueueLen    int
	MaxConnections int
}

var SinxConfig *ConfigObject

func init() {
	SinxConfig = &ConfigObject{
		Name:           "Sinx_Server",
		Host:           "127.0.0.1",
		Port:           33366,
		MaxConn:        3,
		MaxTranSize:    1024,
		version:        "V0.9",
		MsgQueueNum:    10,
		MsgQueueLen:    1024,
		MaxConnections: 1024,
	}

	ConfReload()
}

func ConfReload() {
	conf, err := ini.Load("../../../config/conf.ini")
	if err != nil {
		fmt.Println("[Warning] config reload error", err)
		return
	}
	section := conf.Section("Sinx_Server")
	SinxConfig.Name = section.Key("Name").String()
	SinxConfig.Host = section.Key("Host").String()

	theconf, _ := section.Key("Port").Int()
	SinxConfig.Port = theconf

	theconf, _ = section.Key("MaxConn").Int()
	SinxConfig.MaxConn = theconf

	theconf, _ = section.Key("MaxTranSize").Int()
	SinxConfig.MaxTranSize = theconf

	theconf, _ = section.Key("MsgQueueNum").Int()
	SinxConfig.MsgQueueNum = theconf

	theconf, _ = section.Key("MsgQueueLen").Int()
	SinxConfig.MsgQueueLen = theconf

	theconf, _ = section.Key("MaxConnections").Int()
	SinxConfig.MaxConnections = theconf

	fmt.Println("conf_reload========", SinxConfig)
}
