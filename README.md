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
`jobclt create job ls -lr`

Retrieving a job:
`jobclt get job $JOB_UUID`

Deleting a job:
`jobclt delete job $JOB_UUID`

Retrieving the logs for a job:
`jobclt logs $JOB_UUID`