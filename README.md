# goenvi - Simple golang environment initializer

[![Go Reference](https://pkg.go.dev/badge/github.com/morkid/goenvi.svg)](https://pkg.go.dev/github.com/morkid/goenvi)
[![Github Actions](https://github.com/morkid/goenvi/workflows/Go/badge.svg)](https://github.com/morkid/goenvi/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/morkid/goenvi)](https://goreportcard.com/report/github.com/morkid/goenvi)
[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/morkid/goenvi)](https://github.com/morkid/goenvi/releases)

Initialize your environment variables in one shot. goenvi is built on top of [Viper](https://github.com/spf13/viper).

Install dependency:
```bash
go get github.com/morkid/goenvi
```

Supported config file:
- .env
- json
- yaml / yml
- toml
- java properties

## How to use goenvi

`cat .env`
```bash
MESSAGE="hello world"
```

`cat main.go`
```go
package main

import (
    "os"
    "fmt",
    "github.com/morkid/goenvi"
    "github.com/spf13/viper"
)

func main() {
    goenvi.Initialize()

    fmt.Println(os.Getenv("MESSAGE"))
    fmt.Println(viper.GetString("MESSAGE"))
}
```

`go run main.go`
```bash
hello world
hello world
```

By default goenvi autoload `.env` file in current working directory.

## Load custom config file

`cat config.json`
```json
{
    "version": {
        "number": 1,
        "name": "v1.0.0"
    }
}
```

```go
package main

import (
    "os"
    "fmt"
    "github.com/morkid/goenvi"
    "github.com/spf13/viper"
)

func main() {
    goenvi.Add("json", "config.json")
    goenvi.Initialize()

    fmt.Println(os.Getenv("VERSION_NUMBER"))
    fmt.Println(os.Getenv("VERSION_NAME"))
    fmt.Println(viper.GetInt("version.number"))
    fmt.Println(viper.GetInt("version.name"))
}
```

## License

Published under the [MIT License](https://github.com/morkid/goenvi/blob/master/LICENSE).

