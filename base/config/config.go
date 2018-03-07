package config

import (
    "github.com/go-ini/ini"
)

var path = "C:/Users/Seiya/workspace/go/src/github.com/sayyeah-t/take2systembot/etc/take2systembot/"

type BotConfig struct {
    var IsLoaded bool
    var Token string
    var Channel string
    var EnabledCommands []string
}

func (c BotConfig) Init() {
    c.IsLoaded = false

    //
    cfg, err := ini.InsensitiveLoad(path + filename)
    if err != nil {
        println(err.Error())
        return
    }
    println("Path: " + path + filename)
    section, err := cfg.GetSection("default")
    if err != nil {
        println(err.Error())
        return
    }

    // load default attributes
    token, err := section.GetKey("token").MustString("none")
    if err != nil {
        println(err.Error())
        return
    }
    c.Token = token
    channel, err := section.GetKey("channel").MustString("none")
    if err != nil {
        println(err.Error())
        return
    }
    c.Channel = channel
    commands, err := section.GetKey("commands").Strings(",")
    if err != nil {
        println(err.Error())
        return
    }
    c.IsLoaded = true
}

func GetConfig() *ini.Section {


    return
}
