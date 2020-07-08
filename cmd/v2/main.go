package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/robfig/cron/v3"
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
	Run()
	CronExpression() string
	SetID(int)
	GetID() int
	SetName(string)
	GetName() string
	GetType() string
	SetType(string)
}

// JobDetailSpec is job detail spec
type JobDetailSpec struct {
	HCL hcl.Body `hcl:",remain"`
}

// HTTPJob httpjob
type HTTPJob struct {
	Name     string
	Type     string
	ID       int
	Driver   string `hcl:"driver"`
	WebHook  string `hcl:"webhook"`
	Location string `hcl:"location"`
	Query    string `hcl:"query"`
	Schedule string `hcl:"schedule"`
}

// Run  HTTPJob exec method
func (job *HTTPJob) Run() {
	log.Printf("DbJob job %s----%s----%s---%s-----%s--------%d", job.Driver, job.Query, job.Schedule, job.Location, job.Name, job.ID)
}

// SetID SetID
func (job *HTTPJob) SetID(id int) {
	job.ID = id
}

// GetID GetID
func (job *HTTPJob) GetID() int {
	return job.ID
}

// SetName setName
func (job *HTTPJob) SetName(name string) {
	job.Name = name
}

// GetName GetName
func (job *HTTPJob) GetName() string {
	return job.Name
}

// SetType SetType
func (job *HTTPJob) SetType(name string) {
	job.Type = name
}

// GetType GetType
func (job *HTTPJob) GetType() string {
	return job.Type
}

// CronExpression get cron expression
func (job *HTTPJob) CronExpression() string {
	return job.Schedule
}

// DbJob dbjob
type DbJob struct {
	Name     string
	Type     string
	ID       int
	WebHook  string `hcl:"webhook,optional"`
	Location string `hcl:"location"`
	Driver   string `hcl:"driver"`
	Query    string `hcl:"query"`
	Schedule string `hcl:"schedule"`
}

// Run  DbJob exec method
func (job *DbJob) Run() {
	log.Printf("DbJob job %s----%s----%s---%s-----%s--------%d", job.Driver, job.Query, job.Schedule, job.Location, job.Name, job.ID)
}

// CronExpression  get cron expression
func (job *DbJob) CronExpression() string {
	return job.Schedule
}

// SetID SetID
func (job *DbJob) SetID(id int) {
	job.ID = id
}

// GetID GetID
func (job *DbJob) GetID() int {
	return job.ID
}

// SetName setName
func (job *DbJob) SetName(name string) {
	job.Name = name
}

// GetName GetName
func (job *DbJob) GetName() string {
	return job.Name
}

// SetType SetType
func (job *DbJob) SetType(name string) {
	job.Type = name
}

// GetType GetType
func (job *DbJob) GetType() string {
	return job.Type
}

// JobContainer jobcontainer
type JobContainer map[string][]Job

const (
	// FILENAME  hcl filename
	FILENAME = "config-v2.hcl"
	// DEFAULTLOCATION default load location
	DEFAULTLOCATION = "demo"
)

var (
	cronhub *cron.Cron
)

func init() {
	cronhub = cron.New(cron.WithChain(
		cron.SkipIfStillRunning(cron.DefaultLogger),
		cron.Recover(cron.DefaultLogger),
	), cron.WithParser(cron.NewParser(
		cron.SecondOptional|cron.Minute|cron.Hour|cron.Dom|cron.Month|cron.Dow|cron.Descriptor,
	)))
}

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
	var buildJobs JobContainer = JobContainer{}
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
		log.Println("read config error")
	}
	for _, item := range jobsSpec.JobSpecs {
		log.Println("job type", item.Type)
		switch jobtype := item.Type; jobtype {
		case "http":
			httpjob := &HTTPJob{Name: item.Name, Type: item.Type, Location: "demo"}
			if item.JobDetailSpec != nil {
				if diag := gohcl.DecodeBody(item.JobDetailSpec.HCL, evalContext(), httpjob); diag.HasErrors() {
					log.Println(diag.Errs())
				}
			}
			buildJobs["http"] = append(buildJobs["http"], httpjob)
		case "db":
			dbjob := &DbJob{Name: item.Name, Type: item.Type, Location: "demo"}
			if item.JobDetailSpec != nil {
				if diag := gohcl.DecodeBody(item.JobDetailSpec.HCL, evalContext(), dbjob); diag.HasErrors() {
					log.Println(diag.Errs())
				}
			}
			buildJobs["db"] = append(buildJobs["db"], dbjob)
		default:
			log.Println("not support type")
		}
	}
	// print all map jobs
	for _, jobs := range buildJobs {
		for _, job := range jobs {
			myjob := job
			id, err := cronhub.AddJob(myjob.CronExpression(), myjob)
			if err != nil {
				fmt.Println("create job error")
			} else {
				myjob.SetID(int(id))
			}
		}
	}
	for name, jobs := range buildJobs {
		fmt.Println(name)
		for _, job := range jobs {
			job := job
			switch job.(type) {
			case *HTTPJob:
				b := job.(*HTTPJob)
				log.Println(b.Name, b.Type, b.ID, b.Driver, b.Schedule)
			case *DbJob:
				b := job.(*DbJob)
				log.Println(b.Name, b.Type, b.ID, b.Driver, b.Schedule)
			default:
				log.Println("not support type")
			}
		}
	}
	cronhub.Run()
}
