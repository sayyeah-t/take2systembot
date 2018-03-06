package base

import(
    "os"
    "os/signal"
    "syscall"
    "strings"

    "github.com/bwmarrin/discordgo"
)

type DiscordInfo struct {
    Token string
    EnabledChannel string
}

func (info DiscordInfo) Init(tk string, ch string){
    info.Token = tk
    info.EnabledChannel = ch
    //splatnet2.Init()
}

func (info DiscordInfo) Run() {
    dg, err := discordgo.New("Bot " + info.Token)
    if err != nil {
        println("error creating Discord session,", err)
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

func (info DiscordInfo) DumpConfig(){
    println("=== Bot Configuration ===")
    println("Bot Token: ", info.Token)
    println("Channel:   ", info.EnabledChannel)
    println("=========================")
}

func (info DiscordInfo) onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
    if m.ChannelID == info.EnabledChannel {
        if m.Author.ID == s.State.User.ID {
            return
        }
        commands, err := parseCommand(m.Content)
        if err == "not command" {
            return
        }
        if err != "" {
            s.ChannelMessageSend(m.ChannelID, err)
            return
        }
        s.ChannelMessageSend(m.ChannelID, firstResponse)
        msg := ruleMessage(commands)
        s.ChannelMessageSend(m.ChannelID, msg)
    }
}

func parseCommand(msg string) ([]string, string) {
    msg = strings.Replace(msg, "　", " ", -1)
    commands := strings.Split(msg, " ")

    matchFirst := false
    matchSecond := false
    // 1つ目のコマンドが合っているか確認
    for _, val := range first {
        if val == commands[0] {
            matchFirst = true
        }
    }
    if !matchFirst {
        return nil, "not command"
    }
    // サブコマンド無しの時はnoneをサブコマンドにしてreturn
    if len(commands) == 1 {
        commands = append(commands, "none")
        return commands, ""
    }

    // 2つ目のコマンドが合っているか確認
    for _, val := range second {
        if val == commands[1] {
            matchSecond = true
        }
    }
    if !matchSecond {
        return nil, commandError + "\n" + secondError
    }

    return commands, ""

}

func ruleMessage(commands []string) string {
    msg := splatnet2.GetRuleMessage(commands[0], commands[1])
    return msg
}
