package job

import (
	"testing"
	"time"
)

func TestManagerSuccess(t *testing.T) {
	// Given a manager and a normal command to execute
	m := NewManager()

	// When we walk through the lifecycle of a process
	// Then manager should manage the process correctly
	j, err := m.Create("go", "help")
	if err != nil {
		t.Error(err)
	}

	j, err = m.Get(j.ID)
	if err != nil {
		t.Error(err)
	}

	// Give the program a second to run
	time.Sleep(250 * time.Millisecond)
	c, err := m.Logs(j.ID)
	if err != nil {
		t.Error(err)
	}

	expected := "Go is a tool for managing Go source code."
	l := <-c
	if l != expected {
		t.Errorf("<-logs = %v; want %v", l, expected)
	}

	_, err = m.Delete(j.ID)
	if err != nil {
		t.Error(err)
	}

	_, err = m.Get(j.ID)
	if err == nil {
		t.Errorf("m.Get(j.ID).err = %v; want %v", nil, NotFound)
	}
}
