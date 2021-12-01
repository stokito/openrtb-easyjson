# OpenRTB

OpenRTB protocol definitions and DTOs for Golang.
Parsed/Serialized with EasyJSON.
Based on github.com/bsm/openrtb

## Requirements

Requires Go 1.17

## Installation

To install, use `go get`:

```shell
go get -u github.com/stokito/openrtb-easyjson@v3.1.0
```

## Usage

```go
package main

import (
  "log"
  "github.com/stokito/openrtb-easyjson"
)

func main() {
  file, err := os.Open("stored.json")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  var req *openrtb.BidRequest
  if err := json.NewDecoder(file).Decode(&req); err != nil {
    log.Fatal(err)
  }

  log.Printf("%+v\n", req)
}
```
