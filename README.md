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
package main

import (
    "regexp"
    "fmt"
    "github.com/fyxme/gonada"
)

func main() {
    gn := gonada.GetNada{}

    // list of available domains to use as emails
    domains := gn.GetDomains()
    fmt.Println(domains)

    email := fmt.Sprintf("%s@%s", "test", domains[0])
    gni := gn.GetInbox(email)
    if gni.IsEmpty() {
        fmt.Println("Inbox is empty")
    }

    // get the contents of the first email in the mailbox
    firstMail := gni.Msgs[0]
    fmt.Println(
        firstMail.FromName,
        firstMail.FromEmail,
        firstMail.Subject,
        firstMail.Timestamp,
        firstMail.GetContents()[:10],
    )

    // find a comfirmation link inside of an email
    r, _ := regexp.Compile("http://somewebsite.com/confirm_email/[0-9A-Za-z]+")
    fromEmail := "do-not-reply@somewebsite.com"
    for _, mail := range gni.Msgs {
        // skip email if not from right address
        if mail.FromEmail != fromEmail {
            continue
        }

        if confirmLink := r.FindString(mail.GetContents()); confirmLink != "" {
            // do something with confirmlink here
            fmt.Println(confirmLink)
            break
        }
    }
}
```

