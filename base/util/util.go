package util

import(
    "strings"
    "errors"
)

var (
    InvalidCommand = "Invalid Command"
)

func ParseCommand(msg string) []string {
    msg = strings.Replace(msg, "　", " ", -1)
    //msg = strings.Replace(msg, "！", "!", -1)
    commands := strings.Split(msg, " ")
    return commands
}

func ValidateCommand(msg string, enabledcommands []string) error {
    for _, c := range enabledcommands {
        if commands[0] == c {
            return nil
        }
    }
    return errors.New(InvalidCommand)
}
