# guolai
Golang bindings for the [Wolai](https://www.wolai.com/product) open API.

## Get Started
Install the package with the following command:
```shell
go get github.com/lemonnekogh/guolai
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/lemonnekogh/guolai"
)

func main() {
	client := guolai.New("YOUR_API_KEY")
	blockApiResp, err := client.GetBlocks("your block id")
	if err != nil {
		panic(err)
	}

	fmt.Printf("blockApiResp: %+v\n", blockApiResp)
}
```
