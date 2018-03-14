package bot

import(
    "os"
    "os/signal"
    "syscall"
    "errors"
    //"strings"
    "github.com/sayyeah-t/take2systembot/base/util"
    "github.com/sayyeah-t/take2systembot/aorichan/config"
    "github.com/bwmarrin/discordgo"
)


var Conf config.AoriConfig


func Init(){
    Conf.Init("aorichan.conf")
    Conf.AdditionalInit()
    //splatnet2.Init()
}

func Run() {
    dg, err := discordgo.New("Bot " + Conf.Token)
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

func DumpConfig(){
    println("=== Bot Configuration ===")
    println("Bot Token: ", Conf.Token)
    println("=========================")
}

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
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

func getMessage(channel string, commands []string) (string, error) {
    if channel == Conf.IksmChannel {
        err := util.ValidateCommand(commands[0], iksmCommand)
        if err != nil {
            println(err.Error())
            return "", err
        }
        msg, err := iksmAction(commands)
        if err != nil {
            println(err.Error())
            return "", err
        }
        return msg, nil
    }

    for _, c := range Conf.WorkChannels {
        if channel == c {
            err := util.ValidateCommand(commands[0], generalCommand)
            if err != nil {
                println(err.Error())
                return "", err
            }
            return "test", nil
        }
    }

    return "", errors.New("Matched no channels")
}
/*
func ruleMessage(commands []string) string {
    msg := splatnet2.GetRuleMessage(commands[0], commands[1])
    return msg
}
*/
