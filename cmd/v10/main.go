package main

import (
	"encoding/json"
	"log"

	"github.com/hashicorp/hcl/v2/hclsimple"
)

// Job job
type Job struct {
	Type string     `hcl:",label"`
	SQLs []*SQLType `hcl:"sqls,block"`
}

// SQLType SqlType
type SQLType struct {
	Name    string `hcl:",label"`
	SQLType string `hcl:"sqltype"`
	SQL     string `hcl:"sql"`
}

func main() {
	var config struct {
		Jobs []*Job `hcl:"job,block"`
	}
	err := hclsimple.DecodeFile("sqljobs.hcl", nil, &config)
	if err != nil {
		log.Println("some err: " + err.Error())
	}
	for _, item := range config.Jobs {
		v, _ := json.Marshal(item)
		log.Printf("message: %s\r\n", string(v))
	}
}
