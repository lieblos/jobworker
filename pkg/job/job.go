package job

import (
	"os/exec"
	"sync"

	"github.com/google/uuid"
)

// Job is the generalized represenation for a process, task or other similar thing.
type Job struct {
	ID   uuid.UUID
	Done chan chan error

	completed bool
	status    string
	rw        sync.RWMutex
}

// manage manages the completion status of an exec.Cmd and
// requests to alter it through the job.
func (j *Job) manage(cmd *exec.Cmd) {
	completed := make(chan error)

	go func() { completed <- cmd.Wait() }()

	for {
		select {
		case <-completed:
			j.rw.Lock()
			j.status = cmd.ProcessState.String()
			j.completed = true
			j.Done = nil
			j.rw.Unlock()
			return
		case c := <-j.Done:
			c <- cmd.Process.Kill()
			j.Done = nil
		}
	}
}

// Status accesses status in a thread-safe manner.
func (j *Job) Status() string {
	j.rw.RLock()
	defer j.rw.RUnlock()
	return j.status
}

// Completed accesses completed in a thread-safe manner.
func (j *Job) Completed() bool {
	j.rw.RLock()
	defer j.rw.RUnlock()
	return j.completed
}
