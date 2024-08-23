package config

import (
    "os"
)

type Config struct {
    Port        string
    CertFile    string
    KeyFile     string
}

func LoadConfig() Config {
    return Config{
        Port:        os.Getenv("PORT"),
        CertFile:    os.Getenv("CERT_FILE"),
        KeyFile:     os.Getenv("KEY_FILE"),
    }
}
