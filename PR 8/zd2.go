package main

import (
    "encoding/json"
    "fmt"
    "os"
)

type Config struct {
    DatabaseURL string
    Port        string
}

type ConfigProvider interface {
    Load() Config
}

type FileConfigProvider struct{}

func (f FileConfigProvider) Load() Config {
    data, _ := os.ReadFile("config.json")
    var c Config
    json.Unmarshal(data, &c)
    return c
}

type EnvConfigProvider struct{}

func (e EnvConfigProvider) Load() Config {
    return Config{os.Getenv("DB_URL"), os.Getenv("PORT")}
}

func main() {
    os.WriteFile("config.json", []byte(`{"DatabaseURL":"Фото_Анапа_2003","Port":"1710"}`), 0)
    
    os.Setenv("DB_URL", "pizza://dodo:1710/Hawaiian")
    os.Setenv("PORT", "1710")
    
    fmt.Println("File:", FileConfigProvider{}.Load())
    fmt.Println("Env:", EnvConfigProvider{}.Load())
}