package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//CREATE STRUCT FOR CONFIG
type Config struct {
	AppRoot  string `json:"AppRoot"`
	Port     string `json:"port"`
	Tlscert  string `json:"TLSCert"`
	Tlskey   string `json:"TLSKey"`
	Database struct {
		Hostname string `json:"hostname"`
		Username string `json:"username"`
		Dbname   string `json:"dbname"`
		Password string `json:"password"`
	} `json:"database"`
}

var cfg Config

//FUNCTION FOR READING CONFIG FILE
func Readconfigfile(confpath string, t bool) (Config, error) {
	file, err := ioutil.ReadFile(confpath)
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}
	if t {
		fmt.Printf("%s\n", string(file))
	}

	json.Unmarshal(file, &cfg)
	return cfg, err
}

