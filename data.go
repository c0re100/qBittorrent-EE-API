package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"
)

func (qB *Client) getData() []string {
    dataURL := qB.ipAddr + "/api/v2/sync/maindata"
    req, err := http.NewRequest("GET", dataURL, nil)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(0)
    }

    req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.139 Safari/537.36")
    req.Header.Add("content-type", "application/x-www-form-urlencoded")
    req.Header.Add("referer", qB.ipAddr)
    req.Header.Add("origin", qB.ipAddr)
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
