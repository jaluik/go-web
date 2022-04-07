package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type MongoConfig struct {
	MongoAddr       string
	MongoPoolLimit  int
	MongoDb         string
	MongoCollection string
}

type JsonStruct struct {
}

func (js *JsonStruct) Load(filename string, v interface{}) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = json.Unmarshal(file, v)
	if err != nil {
		log.Fatal(err)
		return
	}
}

type Config struct {
	Port  string
	Mongo MongoConfig
}

func main() {
	JsonParse := &JsonStruct{}
	v := &Config{}
	JsonParse.Load("chapter6/6.2/json_parse.json", v)
	fmt.Println("Port", v.Port)
	fmt.Println(v.Mongo)
}
