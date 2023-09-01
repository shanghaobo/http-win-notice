package main

type ConfigType struct {
	ServerPort int    `yaml:"server_port"`
	Token      string `yaml:"token"`
	ApiToken   string `yaml:"api_token"`
	ApiPort    int    `yaml:"api_port"`
}
