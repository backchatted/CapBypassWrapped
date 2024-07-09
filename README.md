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
	"github.com/backchatted/CapBypassWrapped"
	"fmt"
)

func main() {
	CapBypass := capbypass.New("Key Here")

	balance, err := CapBypass.Balance()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(balance.Credits)

	solve, err := CapBypass.Solve(map[string]any{
		"type":             "FunCaptchaProxylessTask",
		"websiteURL":       "https://iframe.arkoselabs.com/",
		"websitePublicKey": "84E1DACC-3B8E-04D6-6E35-2A7D2B8ACFE1",
		"websiteSubdomain": "https://client-api.arkoselabs.com",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(solve.Solution)
}
```

## Contact
Join the [Discord](https://discord.gg/capbypass)
