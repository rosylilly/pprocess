package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	User    string `json:"user"`
	Token   string `json:"token"`
	Device  string `json:"device"`
	Sound   string `json:"sound"`
	Message string
}

func NewConfig() *Config {
	c := &Config{}

	homeDir := os.Getenv("HOME")
	bytes, err := ioutil.ReadFile(homeDir + "/.pprocess.json")
	if err == nil {
		err := json.Unmarshal(bytes, c)

		if err != nil {
			log.Fatal(err)
		}
	}

	return c
}
