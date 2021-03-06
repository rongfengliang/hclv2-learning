package main

import (
	"log"
	"os"

	"github.com/robfig/cron/v3"
)

type Job struct {
	Name     string
	Schedule string
}

func (job *Job) Exec() {
	log.Println(job.Name)
}

func (job *Job) GetName() string {
	return job.Name
}

func (job *Job) Run() {
	job.Exec()
}

var (
	sches  map[string][]*Job
	myjobs []*Job
)

func init() {
	sches = map[string][]*Job{}
	for i := 0; i < 2; i++ {
		sches["db"] = []*Job{
			{
				Name:     "demo1",
				Schedule: "*/5 * * * * *",
			},
			{
				Name:     "demo2",
				Schedule: "*/10 * * * * *",
			},
		}
	}
	for i := 0; i < 2; i++ {
		sches["http"] = []*Job{
			{
				Name:     "http1",
				Schedule: "*/15 * * * * *",
			},
			{
				Name:     "http2",
				Schedule: "*/20 * * * * *",
			},
		}
	}
}

func main() {
	log := cron.VerbosePrintfLogger(log.New(os.Stdout, "cron: ", log.LstdFlags))
	cronhub := cron.New(cron.WithParser(cron.NewParser(
		cron.SecondOptional|cron.Minute|cron.Hour|cron.Dom|cron.Month|cron.Dow|cron.Descriptor,
	)), cron.WithLogger(log))
	for jobType, jobs := range sches {
		log.Info(jobType)
		for _, job := range jobs {
			myjobs = append(myjobs, job)
			if err := (func(job *Job) error {
				_, err := cronhub.AddFunc(job.Schedule, func() {
					job.Exec()
				})
				return err
			})(job); err != nil {
				log.Info("some wrong")
			}
		}
	}
	cronhub.Run()

}
