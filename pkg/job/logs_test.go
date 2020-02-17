package job

import (
	"strings"
	"testing"
)

func TestReadToChannel(t *testing.T) {
	r := strings.NewReader("hello, world\nwhats up")
	c := make(chan string)

	go readToChannel(r, c)

	expected := "hello, world"
	if l := <-c; l != expected {
		t.Errorf("<- logs = %v; want %v", l, expected)
	}

	expected = "whats up"
	if l := <-c; l != expected {
		t.Errorf("<- logs = %v; want %v", l, expected)
	}
}
