package yconfig
import (
    "encoding/json"
    "os"
    "fmt"
)

// читает файл и правращает его в стректуру Config
func Conf(config *Config, configPath *string) {
    if len(*configPath)<2{
	*configPath = "/etc/postcode/" + os.Args[0]
    }
    confFile, err := os.Open(*configPath)
    if err != nil {
        fmt.Printf("Error can't open config file: %s\n", err)
	os.Exit(-1)
    }
    err = json.NewDecoder(confFile).Decode(&config)
    if err != nil {
        fmt.Printf("Error decoding config file: %s\n", err)
	os.Exit(-2)
    }
    confFile.Close()
}
