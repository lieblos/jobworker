package job

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"sync"

	"github.com/google/uuid"
)

const (
	// NotFound is the error string non existent job is requested.
	NotFound = "job not found"
	// AlreadyCompleted is the error string when a job is concurrently killed or completed.
	AlreadyCompleted = "job is already exiting"
)

// Manager is an implementation of Service that uses unix processes.
type Manager struct {
	jobs sync.Map
}

// NewManager initializes a new ProcessService and its global resources.
func NewManager() Manager {
	err := os.MkdirAll(LogsPath, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	return Manager{}
}

// Create creates a job from the name and args provided, redirects its stdout a file to
// and starts a goroutine that monitors requests to kill the job and it's completion status.
func (m *Manager) Create(name string, args ...string) (*Job, error) {
	id := uuid.New()

	f, err := os.OpenFile(logPath(id), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return nil, err
	}

	cmd := exec.Command(name, args...)
	cmd.Stdout = f
	cmd.Stderr = f

	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	j := &Job{
		ID:     id,
		Status: make(chan string),
		Delete: make(chan chan string),
	}

	go j.manage(cmd)

	m.jobs.Store(id, j)

	return j, nil
}

// Get retrieves a process by uuid.
func (m *Manager) Get(id uuid.UUID) (*Job, error) {
	j, ok := m.jobs.Load(id)
	if !ok {
		return nil, errors.New(NotFound)
	}

	return j.(*Job), nil
}

// Delete creates a process by uuid, taking care to not kill killed processes.
func (m *Manager) Delete(id uuid.UUID) (string, error) {
	j, ok := m.jobs.Load(id)
	if !ok {
		return "", errors.New(NotFound)
	}

	m.jobs.Delete(id)

	c := make(chan string)
	j.(*Job).Delete <- c

	return <-c, nil
}

// Logs retrieves logs for the job corresponding to the uuid passed. Logs is intentionally
// not gated by validation of the uuid.UUID passed in so that logs of deleted jobs can be retrieved.
func (m *Manager) Logs(id uuid.UUID) (chan (string), error) {
	f, err := os.Open(logPath(id))
	if err != nil {
		return nil, err
	}

	// tailing log files correctly isn't the simplest thing decided to out of scope it,
	// github.com/hpcloud/tail is a drop-in replacement for the functionality I want
	c := make(chan string)
	go func() {
		defer f.Close()
		readToChannel(f, c)
	}()

	return c, nil
}
