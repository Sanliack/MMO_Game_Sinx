package conf

import (
	"fmt"
	"gopkg.in/ini.v1"
)

type MMOConfigModel struct {
	MinX   int
	MaxX   int
	CountX int
	MinY   int
	MaxY   int
	CountY int
}

var MMOConfig = &MMOConfigModel{
	MinX:   0,
	MaxX:   300,
	CountX: 5,
	MinY:   0,
	MaxY:   300,
	CountY: 5,
}

func init() {
	config, err := ini.Load("mmo/conf/conf.ini")
	if err != nil {
		fmt.Println("[Warning] Config file open error", err)
	}
	con := config.Section("mmo")
	MMOConfig.MinX, _ = con.Key("MinX").Int()
	MMOConfig.MaxX, _ = con.Key("MaxX").Int()
	MMOConfig.CountX, _ = con.Key("CountX").Int()
	MMOConfig.MinY, _ = con.Key("MinY").Int()
	MMOConfig.MaxY, _ = con.Key("MaxY").Int()
	MMOConfig.CountY, _ = con.Key("CountY").Int()

}
