# Glex

I write too many command parsers so this is a simple package for me to not rewrite this all the time.

## Installation

```bash
go get -u github.com/vcokltfre/glex
```

## Usage

```go
package main

import (
    "fmt"

    "github.com/vcokltfre/glex"
)

func main() {
    words, err := glex.SplitCommand("test 123 'abc def'")
    if err != nil {
        panic(err)
    }

    fmt.Println(words)
}
```
