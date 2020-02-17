package job

import (
	"bufio"
	"io"

	"github.com/google/uuid"
)

const (
	// LogsPath is the base path where the logs are
	LogsPath = "/tmp/log/jobworker/"
	// LogSuffix is suffix for log files
	LogSuffix = ".log"
)

// ReadToChannel reads the logs for a job to a channel
func ReadToChannel(r io.Reader, c chan string) {
	buf := bufio.NewReader(r)

	for {
		// TODO: we drop isPrefix here, to make logs look as expected for
		// long lines, we should send it to the client so they can know when
		// to print a \n and when to not.
		b, _, err := buf.ReadLine()
		if err == io.EOF {
			close(c)
			return
		}

		if err != nil {
			c <- "error reading log: " + err.Error()
			close(c)
			return
		}

		c <- string(b)
	}
}

// LogPath is the log path
// TODO use os.Path for windows and logs should go to /var/log or user home
func LogPath(id uuid.UUID) string {
	return LogsPath + id.String() + LogSuffix
}
