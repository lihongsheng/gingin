package config

import (
	"fmt"
	yaml2 "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type mysql struct {
	User string `yaml:"user"`
	Password string `yaml:"password"`
	Port int `yaml:"port"`
	Chart string `yaml:"chart"`
	Host string `yaml:"host"`
	DbName string `yaml:"db_name"`
}
type redis struct {
	User string `yaml:"user"`
	Password string `yaml:"password"`
	Host string `yaml:"host"`
	Port int `yaml:"port"`
	Optins map[string]string `yaml:"optins"`
}
type configMap struct {
	Port int `yaml:"port"`
	Redis redis `yaml:"redis"`
	Mysql mysql `yaml:"mysql"`
}

var ConfigMap *configMap
func init() {
	ConfigMap = new(configMap)
	fmt.Println("解析配置文件")
	yamlFile, err :=ioutil.ReadFile("config/config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
		panic("读取配置出差" + fmt.Sprintf("#%v ", err))
	}
	err = yaml2.Unmarshal(yamlFile, ConfigMap)
	if err != nil {
		log.Fatalf("解析配置出差: %v", err)
		panic("解析配置出差" + fmt.Sprintf("#%v ", err))
	}
}


