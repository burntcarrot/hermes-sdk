<div align="center">
    <br>
    <br>
    <p>
    <img src="https://github.com/deepsourcelabs/hermes/raw/master/logo.svg" alt="Hermes" width="300">
    </p>
    <h2>Hermes SDK</h2>
</div>

## Hermes SDK

SDK for [Hermes](https://github.com/deepsourcelabs/hermes); the open-source notification API.

## Quickstart

Import the SDK into your project and get started:

```go
import "github.com/deepsourcelabs/hermes-sdk/sdk"
```

## Example

Here is an example which uses the SDK to send a message to a Slack channel:

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/deepsourcelabs/hermes-sdk/sdk"
)

func main() {
	c, err := sdk.NewClient("http://localhost:7272/messages")
	if err != nil {
		log.Fatalln(err)
	}

	// set up slack service with token
	token := os.Getenv("SLACK_TOKEN")
	c.Slack.Setup(token)

	// get template
	tmp := sdk.GetTemplate("template-1")
	payload := &map[string]interface{}{
		"name": "Apollo",
	}

	// send message
	ctx := context.Background()
	resp, err := c.Slack.Send(ctx, tmp, payload, "general")
	if err != nil {
		log.Fatalln(err)
	}
}
```

## Providers

Here is a list of providers which are currently supported:
- [x] Slack
- [x] Discord
- [ ] Linear
- [ ] Jira
- [ ] GitHub Issues
- [ ] Mailgun
