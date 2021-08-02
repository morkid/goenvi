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
- see more about [supported config](https://github.com/spf13/viper#what-is-viper)

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

`cat main.go`
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

## Multiple variations in one shot

`cat main.go`
```go
package main

import (
    "os"
    "fmt"
    "github.com/morkid/goenvi"
    ...
)

func main() {
    goenvi.Add("properties", "config.properties")
    goenvi.Add("json", "config.json")
    goenvi.Add("toml", "config.toml")
    goenvi.Add("yaml", "config.yaml")
    goenvi.Add("dotenv", ".env")
    goenvi.Initialize()

    fmt.Println(os.Getenv("VERSION_NUMBER"))
    fmt.Println(os.Getenv("VERSION_NAME"))
    fmt.Println(viper.GetInt("version.number"))
    fmt.Println(viper.GetInt("version.name"))
}
```

## Register custom viper instance
```go
func main() {
    myEnv := viper.New()
    goenvi.Register(myEnv, true)
    goenvi.Initialize()
}
```

## Register command-line parameters as environment

by implementing `goenvi.FlagSetProvider` interface, you can register command-line parameters as environment variables.

```go
import (
    "github.com/spf13/pflag"
)

type myFlagSet struct {}
func (myFlagSet) VisitAll(fn func(*pflag.FlagSet)) {
    defaultValue := viper.GetString("message")
    pflag.String("message", defaultValue, "message to show")
    pflag.Parse()

    fn(pflag.CommandLine)
}

func main() {
    goenvi.AddFlagSetProvider(myFlagSet{})
    goenvi.Initialize()
}
```

> Note:  
> `pflag` will override some environment variables if command-line parameters specified

## License

Published under the [MIT License](https://github.com/morkid/goenvi/blob/master/LICENSE).

