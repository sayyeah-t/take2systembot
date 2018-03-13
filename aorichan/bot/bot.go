package bot

import(
    "os"
    "os/signal"
    "syscall"
    "strings"
    "github.com/sayyeah-t/base/util"
    "github.com/sayyeah-t/aorichan/config"
    "github.com/bwmarrin/discordgo"
)

type BotInfo struct {
    Conf config.AoriConfig
}

func (b *BotInfo) Init(tk string, ch string){
    b.Conf.Init("aorichan.conf")
    b.Conf.AdditionalInit()
    //splatnet2.Init()
}

func (b *BotInfo) Run() {
    dg, err := discordgo.New("Bot " + b.Conf.Token)
    if err != nil {
        println("error creating Discord session,", err)////////////////
        return
    }

    dg.AddHandler(onMessageCreate)

    err = dg.Open()
    if err != nil {
        println("error opening connection,", err)
        return
    }
    sc := make(chan os.Signal, 1)
    signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
    <-sc

    dg.Close()
}

func (b *BotInfo) DumpConfig(){
    println("=== Bot Configuration ===")
    println("Bot Token: ", b.Token)
    println("=========================")
}

func (b *BotInfo) onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate)
    if m.Author.ID == s.State.User.ID {
        return
    }
    commands := util.ParseCommand(m.Content)
    msg, err := getMessage(m.ChannelID, commands)
    if err != nil {
        println(err.Error())
        return
    }
    s.ChannelMessageSend(m.ChannelID, msg)
}

func (b *BotInfo) getMassage(channel string, commands []string) string, error {
    if channel == b.Conf.IksmChannel {
        err := util.ValidateCommand(command[0], iksmCommand)
        if err != nil {
            println(err.Error())
            return "", err
        }

    }

    for _, c := range b.Conf.WorkChannels {
        if channel == c {
            err := util.ValidateCommand(command[0], generalCommand)
            if err != nil {
                println(err.Error())
                return msg, err
            }

        }
    }
}
/*
func ruleMessage(commands []string) string {
    msg := splatnet2.GetRuleMessage(commands[0], commands[1])
    return msg
}
*/
