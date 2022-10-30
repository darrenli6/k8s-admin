package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"k8s.io/klog"
	"sigs.k8s.io/yaml"
)

type KeyName string

const (
	ServerName KeyName = "server_name"
	ServerHost KeyName = "server_host"
	ServerPort KeyName = "server_port"
)

var keyMap map[KeyName]interface{}

type Config struct {
	Server Server
}

type Server struct {
	Name string `yaml:"name"`
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func init() {
	var config Config
	yamlFile, err := ioutil.ReadFile("./config/application.yaml")
	if err != nil {
		klog.Fatal(err)
		return
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		klog.Fatal(err)
		return
	}
	fmt.Println(os.Getwd())

	keyMap = make(map[KeyName]interface{})
	keyMap[ServerName] = config.Server.Name
	keyMap[ServerHost] = config.Server.Port
	keyMap[ServerPort] = config.Server.Port

}

func GetString(KeyName KeyName) string {
	return keyMap[KeyName].(string)
}

func GetInt(KeyName KeyName) int {
	return keyMap[KeyName].(int)
}
