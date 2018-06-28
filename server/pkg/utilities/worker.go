package utilities

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type Worker struct {
	name        string
	instruction string
	sleeper     Sleeper
}

var (
	queuedJobs []Worker
	errs       chan error
)

func Enqueue(jobs ...Worker) error {
	if jobs == nil {
		return errors.New("Need to pass in value")
	}
	for _, job := range jobs {
		queuedJobs = append(queuedJobs, job)
	}
	return nil
}

func StartProcessing() <-chan error {
	errs = make(chan error, 1)

	for _, job := range queuedJobs {
		wg.Add(1)
		go startJob(job)
	}

	wg.Wait()
	close(errs)

	return errs
}

func startJob(job Worker) {
	defer wg.Done()

	SleepUntil(job.sleeper)
	fmt.Println(job.name + " beginning")
	command := strings.Fields(job.instruction)
	out, err := exec.Command(command[0], command[1:]...).Output()
	if err != nil {
		// log.Fatal(err)
		errs <- errors.New("Test " + job.name + " failed during execution")
	}
	log.Println(out)
}
