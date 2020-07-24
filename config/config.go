package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Message struct {
	Employee struct {
		PhysicalDisability string `yaml:"physicalDisability"`
		Name               string `yaml:"name"`
		Birthday           string `yaml:"birthday"`
		ExperienceTime     string `yaml:"experienceTime"`
		Salary             string `yaml:"salary"`
	}
}

var (
	message    *Message
)

func Config() *Message {
	pwd, _ := os.Getwd()
	data, err := ioutil.ReadFile(pwd+"/src/github.com/cimarinho/golang/cadastro_funcionario/error.yaml")
	if err != nil {
		fmt.Println("Arquivo não existe")
	}
	err = yaml.Unmarshal(data, &message)
	return message
}