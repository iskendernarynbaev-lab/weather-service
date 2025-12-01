package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-co-op/gocron/v2"
)

func runJobs(scheduler gocron.Scheduler) ([]gocron.Job, error) {
	j, err := scheduler.NewJob(
		gocron.DurationJob(
			10*time.Second,
		),
		gocron.NewTask(
			func(a string, b int) {
				log.Println("job run", a, b)
			},
			"hello",
			1,
		),
	)
	if err != nil {
		log.Fatalf("cron run error: %s", err)
		return nil, err
	}
	// each job has a unique id
	fmt.Println(j.ID())

	return []gocron.Job{j}, nil
}
