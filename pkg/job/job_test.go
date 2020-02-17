package job

import (
	"os/exec"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestManageKilledJob(t *testing.T) {
	// Given we start a long running process
	cmd := exec.Command("sleep", "60")

	err := cmd.Start()
	if err != nil {
		t.Error(err)
	}

	j := &Job{
		ID:        uuid.New(),
		Done:      make(chan chan error),
		completed: false,
		status:    cmd.ProcessState.String(),
	}

	go j.manage(cmd)

	// When we kill it
	c := make(chan error)

	j.Done <- c

	// Then it should stop without error and indicate it was killed shortly after.
	err = <-c
	if err != nil {
		t.Error(err)
	}

	for i := 0; i < 10; i++ {
		if j.Completed() {
			expected := "signal: killed"
			if j.Status() != expected {
				t.Errorf("j.Status() = %v; want %v", j.Status(), expected)
			}
			return
		}

		// Yes, time dependent tests are bad, but if this takes more than 2.5 seconds we have a problem.
		time.Sleep(250 * time.Millisecond)
	}
	t.FailNow()
}

func TestManageCompletedJob(t *testing.T) {
	// Given we start a process
	cmd := exec.Command("go", "help")

	err := cmd.Start()
	if err != nil {
		t.Error(err)
	}

	j := &Job{
		ID:        uuid.New(),
		Done:      make(chan chan error),
		completed: false,
		status:    cmd.ProcessState.String(),
	}

	go j.manage(cmd)

	// When we let it finish
	time.Sleep(250 * time.Millisecond)

	// Then it should indicate it is completed and exited successfully
	if !j.Completed() {
		t.Errorf("j.Completed() = %v; want %v", j.Status(), true)
	}

	expected := "exit status 0"
	if j.Status() != expected {
		t.Errorf("j.Status() = %v; want %v", j.Status(), expected)
	}
	return
}

func TestManageErroredJob(t *testing.T) {
	// Given we start a process that ends poorly
	cmd := exec.Command("python", "-c", "import time; time.sleep(0.1); raise Exception(\"hell\")")

	err := cmd.Start()
	if err != nil {
		t.Error(err)
	}

	j := &Job{
		ID:        uuid.New(),
		Done:      make(chan chan error),
		completed: false,
		status:    cmd.ProcessState.String(),
	}

	go j.manage(cmd)

	// When we let it finish
	time.Sleep(250 * time.Millisecond)

	// Then it should indicate it is completed and exited horribly
	if !j.Completed() {
		t.Errorf("j.Completed() = %v; want %v", j.Status(), true)
	}

	expected := "exit status 1"
	if j.Status() != expected {
		t.Errorf("j.Status() = %v; want %v", j.Status(), expected)
	}
	return
}
