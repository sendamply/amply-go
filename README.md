# Amply

This is the Amply Go SDK that integrates with the [v1 API](https://docs.sendamply.com/docs/api/docs/Introduction.md).

__Table of Contents__

- [Install](#install)
- [Quick Start](#quick-start)
- [Methods](#methods)
	- [Email](#Email)

## Install

### Prerequisites
- Go
- Amply account, [sign up here.](https://sendamply.com/plans)

### Access Token

Obtain your access token from the [Amply UI.](https://sendamply.com/home/settings/access_tokens)

### Install Package
```
go get github.com/sendamply/amply-go
```

### Domain Verification
Add domains you want to send `from` via the [Verified Domains](https://sendamply.com/home/settings/verified_domains) tab on your dashboard.

Any emails you attempt to send from an unverified domain will be rejected.  Once verified, Amply immediately starts warming up your domain and IP reputation.  This warmup process will take approximately one week before maximal deliverability has been reached.

## Quick Start
The following is the minimum needed code to send a simple email. Use this example, and modify the `to` and `from` variables:

```go
package main

import (
        "fmt"
        "os"

        "github.com/sendamply/amply-go"
)

func main() {
        amply.SetAccessToken(os.Getenv("AMPLY_ACCESS_TOKEN"))

        resp, err := amply.Email.Create(amply.EmailData{
                To: "test@example.com",
                From: "test@verifieddomain.com",
                Subject: "My first Amply email!",
                Text: "This is easy",
                Html: "<strong>and fun :)</strong>",
        })

        if err != nil {
                fmt.Println(err)
        }

        fmt.Println(resp.StatusCode)
        fmt.Println(resp.Body)
}
```

Once you execute this code, you should have an email in the inbox of the recipient.  You can check the status of your email in the UI from the [Search](https://sendamply.com/home/analytics/searches/basic/new), [SQL](https://sendamply.com/home/analytics/searches/sql/new), or [Users](https://sendamply.com/home/analytics/users) page.

## Methods

### Email

Parameter(s)         | Description
:---------------- | :---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
To, Cc, Bcc | Email address of the recipient(s).  This may be a string `"Test <test@example.com>"`, a struct `amply.EmailAddress{Name: "Test", Email: "test@example.com"}`, or an array of strings/structs.
Personalizations | For fine tuned access, you may override the To, Cc, and Bcc keys and use advanced personalizations.  See the API guide [here](https://docs.sendamply.com/docs/api/Mail-Send.v1.yaml/paths/~1email/post).
From | Email address of the sender.  This may be formatted as a string or `amply.EmailAddress{}`.  An array of senders is not allowed.
Subject | Subject of the message.
Html | HTML portion of the message.
Text | Text portion of the message.
Content | An array of `amply.Content{}` structs containing the following fields: `Type` (required), `Value` (required).
Template | The template to use. This may be a string (the UUID of the template), an array of UUID strings (useful for A/B/... testing where one is randomly selected), or a map of the format `map[string]float{"template1Uuid": 0.25, "template2Uuid": 0.75}` (useful for weighted A/B/... testing).
DynamicTemplateData | The dynamic data to be replaced in your template.  This is a map of the format `map[string]string{"variable1": "replacement1", ...}`. Variables should be defined in your template body as `${variable1}`.
ReplyTo |Email address of who should receive replies.  This may be a string or an amply.EmailAddress with `Name` (optional) and `Email` (required) fields.
Headers | A `map[string]string` where the header name is the key and header value is the value.
IpOrPoolUuid | The UUID string of the IP address or IP pool you want to send from.  Default is your Global pool.
UnsubscribeGroupUuid | The UUID string of the unsubscribe group you want to associate with this email.
Attachments[] | An array of Attachment structs `[]amply.Attachments{}`.
Attachments[][Content] | A base64 encoded string of your attachment's content (string).
Attachments[][Type] | The MIME type of your attachment (string).
Attachments[][Filename] | The filename of your attachment (string).
Attachments[][Disposition] | The disposition of your attachment (`inline` or `attachment`) (string).
Attachments[][ContentId] | The content ID of your attachment (string).
Clicktracking | Enable or disable clicktracking (bool).
Categories | A `[]string` of email categories you can associate with your message.
Substitutions | A `map[string]string` of the format `map[string]string{"subFrom": "subTo", ...}` of substitutions.

__Example__

```go
amply.Email.Create(amply.EmailData{
        To: "example@test.com",
        Cc: amply.EmailAddress{ Name: "Billy", Email: "Smith" },
        From: "From <example@verifieddomain.com>",
        Text: "Text part",
        Html: "HTML part",
        Content: []amply.Content{
                amply.Content{
                        Type: "text/testing",
                        Value: "Test!",
                },
        },
        Subject: "A new email!",
        ReplyTo: "Reply To <test@example.com>",
        Template: "faecb75b-371e-4062-89d5-372b8ff0effd",
        DynamicTemplateData: map[string]string{"name": "Jimmy"},
        UnsubscribeGroupUuid: "5ac48b43-6e7e-4c51-817d-f81ea0a09816",
        IpOrPoolUuid: "2e378fc9-3e23-4853-bccb-2990fda83ca9",
        Attachments: []amply.Attachment{
                amply.Attachment{
                        Content: "dGVzdA==",
                        Filename: "test.txt",
                },
        },
        Headers: map[string]string{"X-Testing": "Test"},
        Categories: []string{"Test"},
        Clicktracking: true,
        Substitutions: map[string]string{"sub1": "replacement1"},
})
```
