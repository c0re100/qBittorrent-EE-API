package main

import (
    "fmt"
    "log"
    "time"
)

func main() {
    fmt.Println("A example program for qBittorrent Enhanced Edition API")

    client := &Client{
        ipAddr: inputText("URL: "),
    }
    client.sId = client.login()
    if client.sId == "" {
        log.Fatal("Login fail.")
    }
    for {
        torrentList := client.getData()
        client.getTorrents(torrentList)
        time.Sleep(1000 * time.Millisecond)
    }
}
