package main

import (
	"log"

	"github.com/robfig/cron/v3"
)

type Job struct {
	Name     string
	Schedule string
}

func (job Job) Exec() {
	log.Println(job.Name)
}

func (job Job) Run() {
	job.Exec()
}

var (
	sches map[string][]Job
)

func init() {
	sches = map[string][]Job{}
	for i := 0; i < 2; i++ {
		sches["db"] = []Job{
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
		sches["http"] = []Job{
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
	cronhub := cron.New(cron.WithChain(
		cron.SkipIfStillRunning(cron.DefaultLogger),
		cron.Recover(cron.DefaultLogger),
	), cron.WithParser(cron.NewParser(
		cron.SecondOptional|cron.Minute|cron.Hour|cron.Dom|cron.Month|cron.Dow|cron.Descriptor,
	)))
	for jobType, jobs := range sches {
		log.Println(jobType)
		for _, job := range jobs {
			// not work as expect
			cronhub.AddFunc(job.Schedule, func() {
				job.Exec()
			})
			// this works
			// cronhub.AddJob(job.Schedule, job)
		}
	}
	cronhub.Run()
}
