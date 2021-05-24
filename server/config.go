package server

type ServerConfig struct {
	Host           string `json:"host"`
	Port           int    `json:"port"`
	RequestPattern string `json:"requestPattern"`
}

func (serverConfig *ServerConfig) getHost() string {
	return serverConfig.Host
}

func (serverConfig *ServerConfig) getPort() int {
	return serverConfig.Port
}

func (serverConfig *ServerConfig) getRequestPattern() string {
	return serverConfig.RequestPattern
}
