package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/zclconf/go-cty/cty"
)

// Job type
type Job struct {
	Type          string `hcl:",label"`
	Name          string `hcl:",label"`
	Driver        string `hcl:"driver"`
	DSN           string `hcl:"dsn"`
	Query         string `hcl:"query"`
	Webhook       string `hcl:"webhook"`
	MyInfo        string `hcl:"myinfo"`
	Schedule      string `hcl:"schedule"`
	MessageString string `hcl:"message"`
	EngineName    string `hcl:"jsengine"`
}

// MyJob myJob
type MyJob map[string]*Job

func main() {
	var config struct {
		Jobs []*Job `hcl:"job,block"`
	}
	userinfo := cty.MapVal(map[string]cty.Value{
		"USERNAME2": cty.StringVal(os.Getenv("USERNAME2")),
	})
	err := hclsimple.DecodeFile("config-v1.hcl", &hcl.EvalContext{
		Variables: map[string]cty.Value{
			"env": userinfo,
		},
	}, &config)
	if err != nil {
		log.Fatalf("Failed to load configuration: %s", err)
	}
	var jobs MyJob = MyJob{}
	for _, item := range config.Jobs {
		job := &Job{
			Type:          item.Type,
			Name:          item.Name,
			Driver:        item.Driver,
			MessageString: item.MessageString,
			EngineName:    item.EngineName,
			Query:         item.Query,
		}
		jobs[item.Name] = job
	}
	for key, v := range jobs {
		msg := fmt.Sprintf("name:%s---type:%s--------dirver:%s----query:%s", key, v.Type, v.Driver, v.Query)
		fmt.Println(msg)
	}
}
