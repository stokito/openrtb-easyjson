# OpenRTB

OpenRTB protocol definitions and DTOs for Golang.
Parsed/Serialized with EasyJSON.
Based on https://github.com/bsm/openrtb

## Requirements

Requires Go 1.17

## Installation

To install, use `go get`:

```shell
go get -u github.com/stokito/openrtb-easyjson/v3@v3.5.1
```

## Usage

```go
package main

import (
  "log"
  "github.com/stokito/openrtb-easyjson/v3"
)

func main() {
  file, err := os.Open("stored.json")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  bidReq := &openrtb.BidRequest{}
  err = easyjson.Unmarshal(file, bidReq)
  if err != nil {
    log.Fatal(err)
  }

  log.Printf("%+v\n", bidReq)
}
```
