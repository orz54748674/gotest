package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type temp struct {
	Province []detail `yaml:"province"`
}

type detail struct {
	Code int    `yaml:"code"`
	Name string `yaml:"name"`
}

func main() {
	content, err := ioutil.ReadFile("yaml/areaCode.yaml")
	if err != nil {
		fmt.Println("YamlLoader ReadFile err ", err)
		return
	}
	var data temp
	err = yaml.Unmarshal(content, &data)
	if err != nil {
		fmt.Println("YamlLoader Unmarshal err ", err)
		return
	}

	fmt.Println(data)
}
