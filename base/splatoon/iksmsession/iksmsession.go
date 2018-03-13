package iksm

import (
    ""
)

testPath = "aaaaaaa"

var (
    Iksm = ""
)

func DumpSessionToken() {
    print("dump session token: " + Iksm)
}

func readIksmFromFile() {

}

func WriteIksmToFile() {
    print("Write file...")
}

func SetIksm(token string) {
    if Iksm != token {
        Iksm = token
    }
    WriteTokenToFile
}
