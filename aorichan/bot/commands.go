package bot

import (
    "errors"
    "github.com/sayyeah-t/take2systembot/base/splatoon/iksmsession"
    //"github.com/sayyeah-t/base/splatoon/schedule"
)

var (
    iksmCommand  = []string{
        "!test",
        "!set_iksm",
    }
    generalCommand = []string{
        "!test",
        "ガチマ",
        "リグマ",
        "!restart",
        "リスタート",
    }
)


func iksmAction(cmd []string) (string, error) {
    var msg string
    var err error

    switch cmd[0] {
    case "!set_iksm":
        if len(cmd) != 2{
            msg = ""
            err = errors.New("invalid subcommand")
            return msg, err
        }
        // Do iksm
        iksm.SetIksm(cmd[1])
        msg = "Succeeded to set iksm_session"
        err = nil
    case "!test":
        print("iksm_session: " + iksm.Iksm)
        msg = iksm.Iksm
        err = nil
    }
    return msg, err
}

func generalAction(cmd []string) (string, error) {
    return "未実装", nil
}
