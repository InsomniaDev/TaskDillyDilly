package utilities

import (
	"fmt"
)

type Worker struct {
	name        string
	instruction string
	sleeper     Sleeper
}

var (
	queuedJobs []Worker
)

func Enqueue(jobs ...Worker) {
	if jobs == nil {
		queuedJobs = make([]Worker, len(jobs))
	}
	for _, job := range jobs {
		queuedJobs = append(queuedJobs, job)
	}
}

func StartProcessing() {
	for _, job := range queuedJobs {
		wg.Add(1)
		go startJob(job)
	}

	wg.Wait()
}

func startJob(job Worker) {
	defer wg.Done()

	SleepUntil(job.sleeper)
	fmt.Println(job.name)
	// TODO: Add logic to execute job in the future
}
