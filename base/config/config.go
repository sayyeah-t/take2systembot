package config

import (
    "github.com/go-ini/ini"
)

var path = "C:/Users/Seiya/workspace/go/src/github.com/sayyeah-t/take2systembot/etc/take2systembot/"

type BotConfig struct {
    var IsLoaded bool
    var Channel string
    var EnabledCommands []string
}

func (c BotConfig) Init() {
    c.IsLoaded = false
    cfg, err := ini.InsensitiveLoad(path + filename)
    if err != nil {
        println(err.Error())
        return
    }

}

func GetConfig(filename string) *ini.Section {
    println("Path: " + path + filename)
    sec1, err := cfg.GetSection("default")
    if err != nil {
        println(err.Error())
        return nil
    }
    return sec1
}
