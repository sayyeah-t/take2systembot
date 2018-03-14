package iksm

import (
    "errors"
)



var (
    Iksm = ""
    filePath = ""
)

func DumpSessionToken() {
    print("dump session token: " + Iksm)
}

func readIksmFromFile() error {
    if filePath == ""{
        return errors.New("file path is empty")
    }
    return nil
}

func SetDumpPath(path string) {
    filePath = path
}

func writeIksmToFile() error {
    print("Write file...")
    if filePath == ""{
        return errors.New("file path is empty")
    }
    return nil
}

func SetIksm(token string) {
    if Iksm != token {
        Iksm = token
    }
    err := writeIksmToFile()
    if err != nil {
        print(err.Error())
    }
}
