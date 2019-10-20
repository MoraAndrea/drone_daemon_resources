package configuration

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type ConfigType struct {
	RabbitConf struct {
		BrokerAddress string `yaml:"broker_address"`
		BrokerPort string `yaml:"broker_port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		QueueResources string `yaml:"queue_resources"`
	} `yaml:"rabbit"`
	Federation struct {
		ExchangeName string `yaml:"exchange_name"`
	} `yaml:"federation"`
	Kubernetes struct {
		Namespace string `yaml:"namespace"`
		ClusterName string `yaml:"cluster_name"`
	} `yaml:"kubernetes"`
}

func Config() *ConfigType{
	c:= new(ConfigType)

	filename, _ := filepath.Abs("configuration/conf.yml")
	yamlFile, err := ioutil.ReadFile(filename)
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		panic(err)
	}

	return c
}