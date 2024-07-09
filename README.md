## CapBypassWrapped

**CapBypassWrapped** is a Go package to interact with the CapBypass API. Supporting all task currently (7-8-2024).

## Installation
To use **CapBypassWrapped** in your Go Project, you can simply run:

```bash
go get github.com/backchatted/CapBypassWrapped
```

## Usage
To use **CapBypassWrapped** here is the following example.

```go
package main

import (
	"fmt"
	"github.com/backchatted/CapBypassWrapped"
)

func main() {
	CapBypass := capbypass.New("Key Here")

	balance, err := CapBypass.Balance()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(balance.Credits)

	var task capbypass.CapBypassPayload

	task.Task.Type = "FunCaptchaProxylessTask"
	task.Task.WebsiteURL = "https://www.example.com"
	task.Task.WebsitePublicKey = "SITE_PUBLIC_KEY"
	// Optional
	task.Task.WebsiteSubdomain = "https://client-api.arkoselabs.com"
	task.Task.Data = "{\"blob\": \"F5ebftFmrdZV9isL.cv4/ihbdWsMD4C1GhDkXD1f72AV4f1Rkhk1QMcNmOeM3XGS7EO8M1CBUmhsFirfsPdmbEtq/cplAQQEbIOo01KI0dq9kyqq0ZqBjUI+7W4L8ePOIxDqLr/Rk/eQLt95+0vos1TFDijvhRLc5YYTK3C4l436NKctdkrBhULF1mJleNUevF6oTxZPpm/A+DQYLPbT+37fwMgreXI5D8KGC2eaEdIXS32EfT39KSO/HJblN9tAVqpqRJRm1laijQfDvjWZq2Dtkfmjo/fnCI8kJoN63fSMtzlgWkmxO/wNkYVSR2RYu1+xCullLzZ9VPuqxst1wq44CK8GDrTL5BQ6y+1XEJrLIwpnBM/9KJl/GY6By9wax9YKn1Hqjl8TET+L4KP3GgScMotuZNKRjE8rDO1eeyCXH07BmzIG5ANDPc7iPw0HsvwTYjCWyDBE8+p1XepTcLRsm6Tv8+fpWjkRGNggv3EA=\""
	// Only optional, when data is not set
	task.Task.Headers = &map[string]interface{}{
		"acceptLanguage": "en-US,en;q=0.5",
		"userAgent":      "CapBypass",
	}

	solve, err := CapBypass.Solve(task)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(solve.Solution)
}

```

## Contact
Join the [Discord](https://discord.gg/capbypass)
