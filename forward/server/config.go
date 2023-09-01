package main

type ConfigType struct {
	SocketPort int    `yaml:"socket_port"`
	Token      string `yaml:"token"`
	ApiToken   string `yaml:"api_token"`
	Port       int    `yaml:"port"`
}
