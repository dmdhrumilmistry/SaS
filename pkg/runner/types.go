package runner

import "github.com/dmdhrumilmistry/sas/pkg/reader"

type PipelineResult struct {
	Stdout   string
	Stderr   string
	ExitCode int
}

type Runner struct {
	Pipeline reader.Pipeline
	Results  []PipelineResult
}
