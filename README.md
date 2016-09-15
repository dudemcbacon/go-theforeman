# go-theforeman

A small sling wrapper for [The Foreman](https://theforeman.org) API in Golang

## Usage

```
package main

import (
  "fmt"
  "github.com/dudemcbacon/go-theforeman"
)

func main() {
  config := &theforeman.Config{
    BaseURL:  "http://foreman.mgmt.nweacolo.pvt/api/",
    Username: "api-user",
    Password: "snip",
  }

  foreman := theforeman.NewClient(config, nil)

  names, _, _ := foreman.EnvironmentsService.ListEnvironmentNames(nil)
  for _, name := range names {
    fmt.Println(name)
  }
}
```
