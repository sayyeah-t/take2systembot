package baseconfig

import (
    "github.com/go-ini/ini"
)

//var path = "C:/Program Files/Git/etc/take2systembot/"
var path = "/etc/take2systembot/"

type BaseConfig struct {
    IsLoaded bool
    Token string
    //Channel string
    //EnabledCommands []string
    Cfg *ini.File
}

func (c *BaseConfig) Init(filename string) {
    c.IsLoaded = false
    cfg, err := ini.InsensitiveLoad(path + filename)
    if err != nil {
        println(err.Error())
        return
    }
    c.Cfg = cfg
    println("Path: " + path + filename)
    section, err := cfg.GetSection("default")
    if err != nil {
        println(err.Error())
        return
    }

    // load default attributes
    token, err := section.GetKey("token")
    if err != nil {
        println(err.Error())
        return
    }
    c.Token = token.MustString("none")
    c.IsLoaded = true
}


func (c *BaseConfig) AdditionalInit() bool {
    println("You can override this function")
    return true
}
/*
func (c BaseConfig) GetConfig(secName string) *ini.Section {
    if c.IsLoaded {
        println("Config file unloaded")
        return nil
    }
    sec, err := c.GetSection(secName)
    if err != nil {
        println(err.Error())
        return nil
    }
    return sec
}
*/
