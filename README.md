# Job-agent

## Building

To build, just run `make` (`go install ./...`).

If your GOBIN is on your path, the commands `joblet` and `jobclt` should now be available.

## Example usage

### Starting the job agent

Run:
`joblet`

### Using the command line tool:

Creating a job:
`jobclt create job ls -lr`,
`jobclt create job bash -c "while sleep 1; do echo "hello"; done"`,
`jobclt create job python3 -c "import time; time.sleep(0.1); raise Exception(\"terrible code\")"`,
`jobclt create job python3 -c "import sys; print(\"out\"); print(\"err\", file=sys.stderr)"` 

Retrieving a job:
`jobclt get job $JOB_UUID`

Deleting a job:
`jobclt delete job $JOB_UUID`

Retrieving the logs (up to now) for a job:
`jobclt logs $JOB_UUID`