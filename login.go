package main

import (
    "bytes"
    "fmt"
    "log"
    "net/http"
    "net/url"
    "strings"
)

func (qB *Client) login() string {
    data := url.Values{}
    data.Add("username", inputText("Username: "))
    data.Add("password", inputPassword())

    loginUrl := qB.ipAddr + "/api/v2/auth/login"
    req, err := http.NewRequest("POST", loginUrl, strings.NewReader(data.Encode()))
    if err != nil {
        log.Fatal(err)
    }

    req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.139 Safari/537.36")
    req.Header.Add("content-type", "application/x-www-form-urlencoded")
    req.Header.Add("referer", qB.ipAddr)
    req.Header.Add("origin", qB.ipAddr)

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
