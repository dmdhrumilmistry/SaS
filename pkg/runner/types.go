package runner

import (
	"sync"

	"github.com/dmdhrumilmistry/sas/pkg/reader"
)

type PipelineResult struct {
	FileVarValue string
	Results      []PipelineStepResult
}

type PipelineStepResult struct {
	Stdout   string
	Stderr   string
	ExitCode int
}

type WorkerResult struct {
	WorkerNumber   int
	FileVarValue   string
	PipelineResult PipelineResult
	Success        bool
}

type Runner struct {
	Pipeline reader.Pipeline
	Results  []WorkerResult

	scanQueue chan string
	wg        sync.WaitGroup
	mu        sync.Mutex
}
