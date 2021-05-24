package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type ServerConfiguration struct {
	Host           string `json:"host"`
	Port           int    `json:"port"`
	RequestPattern string `json:"requestPattern"`
}

func (serverConfig *ServerConfiguration) GetHost() string {
	return serverConfig.Host
}

func (serverConfig *ServerConfiguration) GetPort() int {
	return serverConfig.Port
}

func (serverConfig *ServerConfiguration) GetRequestPattern() string {
	return serverConfig.RequestPattern
}

func (serverConfig *ServerConfiguration) LoadConfigFile() {
	data, err := ioutil.ReadFile("config/server.json")
	if err != nil {
		fmt.Println("Error parsing:", err.Error())
	}

	err = json.Unmarshal(data, serverConfig)
	if err != nil {
		fmt.Println("Error setting data:", err.Error())
	}
}
