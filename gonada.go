package gonada

// https://gist.github.com/dmikalova/5693142

import (
    "net/http"
    "encoding/json"
    "io/ioutil"
    "strings"
)

const (
    BASE_URL = "https://getnada.com/api/v1"
)

func generateURL(subpaths ...string) string {
    return strings.Join(append([]string{BASE_URL}, subpaths...), "/")
}

type GetNada struct {
    domains []domainStruct
}

type domainStruct struct {
        Id string `json:"_id"`
        Name string
        Keep *int `json:",string"`
}

type getNadaInbox struct {
    Last int
    Total string
    Msgs []getNadaEmail
}

type getNadaEmail struct {
    Uid string
    FromName string `json:"f"`
    FromEmail string `json:"fe"`
    Subject string `json:"s"`
    Timestamp int `json:"r"`
    Html *string
}

func (gn *GetNada) RefreshDomains() {
    url := generateURL("domains")
    res, err := http.Get(url)
    if err != nil {
        panic(err)
    }
    defer res.Body.Close()

    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        panic(err)
    }

    err = json.Unmarshal(body, &gn.domains)
    if err != nil {
        panic(err)
    }
}

func (gn *GetNada) GetDomains() []string {
    if gn.domains == nil {
        gn.RefreshDomains()
    }

    var list []string
    for _, domain := range gn.domains {
        list = append(list, domain.Name)
    }

    return list
}


func (gn *GetNada) GetInbox(address string) getNadaInbox {
    url := generateURL("inboxes", address)
    gni := getNadaInbox{}
    res, err := http.Get(url)
    if err != nil {
        panic(err)
    }
    defer res.Body.Close()

    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        panic(err)
    }

    err = json.Unmarshal(body, &gni)
    if err != nil {
        panic(err)
    }
    return gni
}

func (gni *getNadaInbox) IsEmpty() bool {
    return len(gni.Msgs) == 0
}

func (gne *getNadaEmail) GetContents() string {
    if gne.Html != nil {
        return *gne.Html
    }

    url := generateURL("messages", gne.Uid)

    res, err := http.Get(url)
    if err != nil {
        panic(err)
    }
    defer res.Body.Close()

    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        panic(err)
    }

    err = json.Unmarshal(body, gne)
    if err != nil {
        panic(err)
    }
    return *gne.Html
}

