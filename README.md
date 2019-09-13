# gonada

gonada is a Golang wrapper around the [getnada.com](https://getnada.com) API. 

getnada.com is a temp email provider.

The Unofficial API Documentation can be found [here](https://github.com/fyxme/pynada#api).

Other wrappers:
- [Python](https://github.com/fyxme/pynada)

## Installation

`go get https://github.com/fyxme/gonada`

## How to use

```Golang
import "github.com/fyxme/gonada"

func main() {
    gn := GetNada{}

    // list of available domains to use as emails
    domains := gn.GetDomains()
    fmt.Println(domains)
    
    email := "test" + domains[0]
    gni := gn.GetInbox(email)
    if gni.IsEmpty() {
        fmt.Println("Inbox is empty")
    }
    
    // find a comfirmation link inside of an email
    r, _ := regexp.Compile("http://somewebsite.com/confirm_email/[0-9A-Za-z]+")
    fmt.Println(r.FindString(gni.Msgs[0].GetContents()))

    // Alternatively you can use it as such
    fmt.Println(r.FindString(gn.GetInbox(email).Msgs[0].GetContents()))
}
```

