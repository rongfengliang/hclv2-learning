package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/zclconf/go-cty/cty"
)

// JobsSpec  JobsSpec
type JobsSpec struct {
	JobSpecs []*JobSpec `hcl:"job,block"`
}

// JobSpec  jobspec
type JobSpec struct {
	Type          string         `hcl:",label"`
	Name          string         `hcl:",label"`
	JobDetailSpec *JobDetailSpec `hcl:"jobinfo,block"`
}

// Job is a job runner
type Job interface {
	Exec() error
}

// JobDetailSpec is job detail spec
type JobDetailSpec struct {
	HCL hcl.Body `hcl:",remain"`
}

// HTTPJob httpjob
type HTTPJob struct {
	Name     string
	Type     string
	Driver   string `hcl:"driver"`
	WebHook  string `hcl:"webhook"`
	Location string `hcl:"location"`
	Query    string `hcl:"query"`
}

// Exec  HTTPJob exec method
func (job *HTTPJob) Exec() error {
	log.Printf("HTTPJob job %s----%s----%s", job.Driver, job.Query, job.Location)
	return errors.New("HTTPJob some wrong")
}

// DbJob dbjob
type DbJob struct {
	Name     string
	Type     string
	Location string `hcl:"location"`
	Driver   string `hcl:"driver"`
	Query    string `hcl:"query"`
}

// Exec  DbJob exec method
func (job *DbJob) Exec() error {
	log.Printf("DbJob job %s----%s----%s", job.Driver, job.Query, job.Location)
	return errors.New("DbJob some wrong")
}

const (
	// FILENAME  hcl filename
	FILENAME = "config-v2.hcl"
	// DEFAULTLOCATION default load location
	DEFAULTLOCATION = "demo"
)

// create eval context
func evalContext() *hcl.EvalContext {
	loca1, loca2 := DEFAULTLOCATION, DEFAULTLOCATION
	if os.Getenv("LOC1") != "" {
		loca1 = os.Getenv("LOC1")
	}
	if os.Getenv("LOC2") != "" {
		loca1 = os.Getenv("LOC2")
	}
	userinfo := cty.MapVal(map[string]cty.Value{
		"LOC1": cty.StringVal(loca1),
		"LOC2": cty.StringVal(loca2),
	})
	return &hcl.EvalContext{
		Variables: map[string]cty.Value{
			"env": userinfo,
		},
	}
}

func main() {

	var jobs []Job
	src, err := ioutil.ReadFile(FILENAME)
	if err != nil {
		panic("some wrong,can't load file content")
	}

	parse := hclparse.NewParser()
	srcHCL, diag := parse.ParseHCL(src, FILENAME)
	if diag.HasErrors() {
		fmt.Println("some wrong")
	}
	jobsSpec := &JobsSpec{}
	if diag := gohcl.DecodeBody(srcHCL.Body, evalContext(), jobsSpec); diag.HasErrors() {
		fmt.Errorf("error in ReadConfig decoding HCL configuration: %w", diag)
	}
	for _, item := range jobsSpec.JobSpecs {
		switch jobtype := item.Type; jobtype {
		case "http":
			httpjob := &HTTPJob{Name: item.Name, Type: item.Type, Location: "demo"}
			if item.JobDetailSpec != nil {
				if diag := gohcl.DecodeBody(item.JobDetailSpec.HCL, evalContext(), httpjob); diag.HasErrors() {
					fmt.Errorf(
						"error in ReadConfig decoding cat HCL configuration: %w", diag,
					)
				}
			}
			jobs = append(jobs, httpjob)
		case "db":
			dbjob := &DbJob{Name: item.Name, Type: item.Type, Location: "demo"}
			if item.JobDetailSpec != nil {
				if diag := gohcl.DecodeBody(item.JobDetailSpec.HCL, evalContext(), dbjob); diag.HasErrors() {
					fmt.Errorf(
						"error in ReadConfig decoding cat HCL configuration: %w", diag,
					)
				}
			}
			jobs = append(jobs, dbjob)
		default:
			// Error in the case of an unknown type. In the future, more types
			// could be added to the switch to support, for example, fish
			// owners.
			fmt.Errorf("error in ReadConfig: unknown pet type ")
		}
	}

	for _, job := range jobs {
		job.Exec()
	}

}
