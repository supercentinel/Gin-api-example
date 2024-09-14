package config

import (
    "os"
    "github.com/pelletier/go-toml"
)

type Config struct {
    SQL struct {
        Host string `toml:"host"`
        Port int `toml:"port"`
        Database string `toml:"database"`
        Username string `toml:"username"`
        Password string `toml:"password"`
    } `toml:"sql"`
}

func ReadConfig(filepath string) Config {
    var cfg Config
    bf, ferr := os.ReadFile(filepath)

    if ferr != nil {
        panic(ferr)
    }

    err := toml.Unmarshal(bf, &cfg)

    if err != nil {
        panic(err)
    }

    return cfg
}
