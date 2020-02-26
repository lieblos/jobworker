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
		ID:     uuid.New(),
		Status: make(chan string),
		Delete: make(chan chan string),
	}

	go j.manage(cmd)

	// When we kill it
	c := make(chan string)

	j.Delete <- c

	// Then it should stop without error and indicate it was killed shortly after.
	expected := "signal: killed"
	if s := <-c; s != expected {
		t.Errorf("s = %v; want %v", s, expected)
	}
}

func TestManageCompletedJob(t *testing.T) {
	// Given we start a process
	cmd := exec.Command("go", "help")

	err := cmd.Start()
	if err != nil {
		t.Error(err)
	}

	j := &Job{
		ID:     uuid.New(),
		Status: make(chan string),
		Delete: make(chan chan string),
	}

	go j.manage(cmd)

	// When we let it finish
	time.Sleep(250 * time.Millisecond)

	// Then it should indicate it is completed and exited successfully
	expected := "exit status 0"
	if s := <-j.Status; s != expected {
		t.Errorf("s = %v; want %v", s, expected)
	}
}

func TestManageErroredJob(t *testing.T) {
	// Given we start a process that ends poorly
	cmd := exec.Command("python", "-c", "import time; time.sleep(0.1); raise Exception(\"hell\")")

	err := cmd.Start()
	if err != nil {
		t.Error(err)
	}

	j := &Job{
		ID:     uuid.New(),
		Status: make(chan string),
		Delete: make(chan chan string),
	}

	go j.manage(cmd)

	// When we let it finish
	time.Sleep(250 * time.Millisecond)

	// Then it should indicate it is completed and exited horribly
	expected := "exit status 1"
	if s := <-j.Status; s != expected {
		t.Errorf("s = %v; want %v", s, expected)
	}
	return
}
