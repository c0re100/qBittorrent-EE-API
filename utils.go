package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "syscall"

    "golang.org/x/crypto/ssh/terminal"
)

type Client struct {
    ipAddr string
    sId    string
}

type Torrent struct {
    Torrents map[string]interface{} `json:"torrents"`
}

type Peer struct {
    Peers map[string]interface{} `json:"peers"`
}

func inputText(msg string) string {
    fmt.Print(msg)
    scanner := bufio.NewScanner(os.Stdin)
    if scanner.Scan() {
        return scanner.Text()
    }
    return ""
}

func inputPassword() string {
    fmt.Print("Password: ")
    bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
    if err != nil {
        log.Fatal(err)
    }
    return string(bytePassword)
}
