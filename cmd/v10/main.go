package main

import (
	"encoding/json"
	"log"

	"gihub.com/rongfengliang/hclv2-learning/cmd/v10/conf"
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
	var myjobs map[string][]*Job = make(map[string][]*Job)
	files, err := conf.AssetDir("jobsconf")
	if err != nil {
		log.Println("#err :" + err.Error())
	} else {
		for _, file := range files {
			log.Println("file name:" + file)
			var config struct {
				Jobs []*Job `hcl:"job,block"`
			}
			jobconfs, _ := conf.Asset("jobsconf/" + file)
			err = hclsimple.Decode(file, jobconfs, nil, &config)
			if err != nil {
				log.Println("#err :" + err.Error())
			}
			myjobs[file] = config.Jobs
		}
	}
	if err != nil {
		log.Println("some err: " + err.Error())
	}
	for _, item := range myjobs {
		for _, job := range item {
			v, _ := json.Marshal(&job)
			log.Printf("message: %s \r\n", string(v))
		}
	}
}
