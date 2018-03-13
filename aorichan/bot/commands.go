package bot

import(
    "errors"
    "github.com/sayyeah-t/base/splatoon/iksmsession"
    "github.com/sayyeah-t/base/splatoon/schedule"
)

var (
    iksmCommand = {
        "!test",
        "!set_iksm",
    }
    generalCommand = {
        "!test",
        "ガチマ",
        "リグマ",
        "!restart",
        "リスタート"
    }
)



func iksmAction(cmd string) string, error{
    switch cmd[0] {
    case "!set_iksm":
        if len(cmd) != 2{
            return "", error.New("invalid subcommand")
        }
        // Do iksm
        return "ToDo: set iksm_session", nil
    case "!test":
        return "テスト", nil
    }
}

func generalAction(cmd []string) string error{
    return "未実装", nil
}
