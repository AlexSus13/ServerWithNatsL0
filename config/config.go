package config

import (
	"github.com/go-yaml/yaml"
	"github.com/pkg/errors"

	"io/ioutil"
)

type DB struct {
	User     string `yaml:"user"`
	Host     string `yaml:"hostdb"`
	Port     string `yaml:"portdb"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
	Sslmode  string `yaml:"sslmode"`
}

type Nats struct {
	ClusterID string `yaml:"clusterid"`
	ClientID  string `yaml:"clientid"`
	NatsUrl   string `yaml:"natsurl"`
	Subject   string `yaml:"subject"`
}

type Conf struct {
	Host             string `yaml:"host"`
	Port             string `yaml:"port"`
	PathToHTMLhome   string `yaml:"pathtohtmlhome"`
	PathToHTMLorders string `yaml:"pathtohtmlorders"`
	DB
	Nats
}

func Get() (*Conf, error) {

	var dconf Conf
	//Reading the yaml file
	yamlFile, err := ioutil.ReadFile("/home/ubuntu/ServerForWb/ServerWithNatsL0/etc/etc.yaml")
	if err != nil {
		return nil, errors.Wrap(err, "Read .yaml File, func Get")
	}
	//Save the received data in the Conf structure
	err = yaml.Unmarshal(yamlFile, &dconf)
	if err != nil {
		return nil, errors.Wrap(err, "Unmarshal .yaml File, func Get")
	}

	return &dconf, nil
}
