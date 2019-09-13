package gonada

import (
    "testing"
    "crypto/rand"
    "fmt"
)

func min(i, j int) int {
    if i > j {
        return j
    }
    return i
}

func generateUUID() string {
    b := make([]byte, 16)
    _, err := rand.Read(b)
    if err != nil {
        fmt.Println(err)
    }
    uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
    return uuid
}

func TestGenerateUrls(t *testing.T) {
    expected := BASE_URL + "/hello/world"
    url := generateURL("hello","world")
    if expected != url {
        t.Errorf("Expected %s but got %s\n", expected, url)
    }
}

func TestRefreshDomains(t *testing.T) {
    gn := GetNada{}

    if gn.domains != nil {
        t.Errorf("Expected gn.domains to be nil\n")
    }

    gn.RefreshDomains()
    if len(gn.domains) <= 0 {
        t.Errorf("Expected gn.domains to be more than 0 and got %d\n", len(gn.domains))
    }
}

func TestGetDomains(t *testing.T) {
    gn := GetNada{}

    domains := gn.GetDomains()
    if len(domains) <= 0 {
        t.Errorf("Expected at least 1 domain got %d\n", len(domains))
    }
}

func TestGetInbox(t *testing.T) {
    gn := GetNada{}

    domains := gn.GetDomains()
    domain := domains[0]

    email := generateUUID() + "@" + domain
    gni := gn.GetInbox(email)
    if gni.Msgs == nil || len(gni.Msgs) > 0 || !gni.IsEmpty() {
        t.Errorf("Unexpected messages in inbox.. %s\n", email)
    }
}

func TestGetContents(t *testing.T) {
    gn := GetNada{}
    email := "asdf@getnada.com"
    gni := gn.GetInbox(email)
    if gni.IsEmpty() {
        t.Errorf("Expecting messages in inbox: %s. It seems empty", email)
        return
    }

    // check first 3 messages in case some are empty
    j := min(len(gni.Msgs), 3)

    valid := false
    for i:=0;i<j;i++ {
        if len(gni.Msgs[i].GetContents()) > 0 {
            valid = true
        }
    }
    if !valid {
        t.Errorf("Expected non empty messages but first %d are empty...", j)
    }
}

func TestCompilation(t *testing.T) {

}
