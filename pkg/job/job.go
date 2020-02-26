package job

import (
	"os/exec"

	"github.com/google/uuid"
)

const (
	// RunningState is the initial state of a job, when start was successful and it hasn't returned
	RunningState = "running"
)

// Job is the generalized represenation for a process, task or other similar thing.
type Job struct {
	ID     uuid.UUID
	Status chan string
	Delete chan chan string

	completed bool
	status    string
}

// manage manages the completion status of an exec.Cmd and
// requests to alter it through the job.
func (j *Job) manage(cmd *exec.Cmd) {
	j.status = RunningState
	j.completed = false

	completed := make(chan error)
	go func() { completed <- cmd.Wait() }()

	for {
		select {
		case j.Status <- j.status:
		case <-completed:
			j.completed = true
			j.status = cmd.ProcessState.String()
		case c := <-j.Delete:
			if j.completed {
				c <- j.status
				return
			}

			err := cmd.Process.Kill()
			if err != nil {
				c <- err.Error()
				return
			}

			<-completed

			c <- cmd.ProcessState.String()
			return
		}
	}
}
