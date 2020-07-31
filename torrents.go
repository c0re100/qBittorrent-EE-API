package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"
    "strings"
)

func (qB *Client) getTorrents(torrentList []string) {
    for _, hash := range torrentList {
        torrentURL := qB.ipAddr + "/api/v2/sync/torrentPeers?hash=" + hash
        req, err := http.NewRequest("GET", torrentURL, nil)
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
