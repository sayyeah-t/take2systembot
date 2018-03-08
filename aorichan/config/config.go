package config

import (
    "github.com/sayyeah-t/take2systembot/base/config"
)

type AoriConfig struct {
    baseconfig.BaseConfig
    IksmChannel string
    WorkChannels []string
    Commands []string
}


func (c *AoriConfig) AdditionalInit() bool {
    section, err := c.Cfg.GetSection("splatoon")
    if err != nil {
        print("cannot read section splatoon")
        return false
    }
    iksmchannel, err := section.GetKey("iksmchannel")
    if err != nil {
        println(err.Error())
        return false
    }
    c.IksmChannel = iksmchannel.MustString("none")

    workchannels, err := section.GetKey("workchannels")
    if err != nil {
        println(err.Error())
        return false
    }
    c.WorkChannels = workchannels.Strings(",")

    commands, err := section.GetKey("commands")
    if err != nil {
        println(err.Error())
        return false
    }
    c.Commands = commands.Strings(",")

    return true
}
