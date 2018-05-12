package main

import (
    "bufio"
    "bytes"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "net/url"
    "os"
    "strings"
    "syscall"
    "time"

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

func (qB *Client) login() string {
    data := url.Values{}
    data.Add("username", inputText("Username: "))
    data.Add("password", inputPassword())

    loginUrl := qB.ipAddr + "/login"
    req, err := http.NewRequest("POST", loginUrl, strings.NewReader(data.Encode()))
    if err != nil {
        log.Fatal(err)
    }

    req.Header.Add("accept", "text/javascript, text/html, application/xml, text/xml, */*")
    req.Header.Add("x-requested-with", "XMLHttpRequest")
    req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.139 Safari/537.36")
    req.Header.Add("content-type", "application/x-www-form-urlencoded")
    req.Header.Add("referer", qB.ipAddr)

    res, err := http.DefaultClient.Do(req)
    if err != nil {
        log.Fatal(err)
    }

    defer res.Body.Close()
    body := &bytes.Buffer{}
    _, err = body.ReadFrom(res.Body)
    if err != nil {
        log.Fatal(err)
    }

    if body.String() == "Ok." {
        fmt.Println()
        fmt.Println("Login success.")
        return res.Cookies()[0].Value
    }
    return ""
}

func (qB *Client) getData() []string {
    dataURL := qB.ipAddr + "/sync/maindata"
    req, err := http.NewRequest("GET", dataURL, nil)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(0)
    }

    req.Header.Add("accept", "text/javascript, text/html, application/xml, text/xml, */*")
    req.Header.Add("x-requested-with", "XMLHttpRequest")
    req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.139 Safari/537.36")
    req.Header.Add("content-type", "application/x-www-form-urlencoded")
    req.Header.Add("referer", qB.ipAddr)
    req.Header.Add("cookie", "SID="+qB.sId)

    res, err := http.DefaultClient.Do(req)
    if err != nil {
        log.Fatal(err)
    }

    defer res.Body.Close()
    body := &bytes.Buffer{}
    _, err = body.ReadFrom(res.Body)
    if err != nil {
        log.Fatal(err)
    }

    var syncData Torrent
    err = json.Unmarshal(body.Bytes(), &syncData)
    if err != nil {
        log.Fatal(err)
    }

    var hashList []string
    for hash := range syncData.Torrents {
        hashList = append(hashList, hash)
    }
    return hashList
}

func (qB *Client) getTorrents(torrentList []string) {
    for _, hash := range torrentList {
        torrentURL := qB.ipAddr + "/sync/torrent_peers?hash=" + hash
        req, err := http.NewRequest("GET", torrentURL, nil)
        if err != nil {
            fmt.Println(err.Error())
            os.Exit(0)
        }

        req.Header.Add("accept", "text/javascript, text/html, application/xml, text/xml, */*")
        req.Header.Add("x-requested-with", "XMLHttpRequest")
        req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.139 Safari/537.36")
        req.Header.Add("content-type", "application/x-www-form-urlencoded")
        req.Header.Add("referer", qB.ipAddr)
        req.Header.Add("cookie", "SID="+qB.sId)

        res, err := http.DefaultClient.Do(req)
        if err != nil {
            log.Fatal(err)
        }

        body := &bytes.Buffer{}
        _, err = body.ReadFrom(res.Body)
        if err != nil {
            log.Fatal(err)
        }

        var peerList Peer
        err = json.Unmarshal(body.Bytes(), &peerList)
        if err != nil {
            log.Fatal(err)
        }

        for peer := range peerList.Peers {
            clientName := peerList.Peers[peer].(map[string]interface{})["client"].(string)
            if strings.Contains(clientName, "BitComet") {
                qB.tempBlockPeer(peerList.Peers[peer].(map[string]interface{})["ip"].(string))
            }
        }
    }
}

func (qB *Client) tempBlockPeer(ip string) {
    blockURL := qB.ipAddr + "/command/tempblockPeer"
    req, err := http.NewRequest("POST", blockURL, strings.NewReader("ip="+ip))
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(0)
    }

    req.Header.Add("accept", "text/javascript, text/html, application/xml, text/xml, */*")
    req.Header.Add("x-requested-with", "XMLHttpRequest")
    req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.139 Safari/537.36")
    req.Header.Add("content-type", "application/x-www-form-urlencoded")
    req.Header.Add("referer", qB.ipAddr)
    req.Header.Add("cookie", "SID="+qB.sId)

    res, err := http.DefaultClient.Do(req)
    if err != nil {
        log.Fatal(err)
    }

    defer res.Body.Close()
    body := &bytes.Buffer{}
    _, err = body.ReadFrom(res.Body)
    if err != nil {
        log.Fatal(err)
    }

    if body.String() == "Done." {
        fmt.Printf("Peer '%v' banned.\n", ip)
    }
}

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
