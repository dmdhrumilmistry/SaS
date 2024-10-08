package runner

import (
	"fmt"

	"github.com/dmdhrumilmistry/sas/pkg/reader"
	"github.com/dmdhrumilmistry/sas/pkg/utils"
	"github.com/rs/zerolog/log"
)

func NewRunner(pipeline reader.Pipeline) *Runner {
	return &Runner{
		Pipeline: pipeline,
	}
}

func (r *Runner) Run() error {
	for _, step := range r.Pipeline.Steps {
		if err := r.runStep(step); err != nil {
			return err
		} else {
			log.Info().Msgf("Ran step successfully: %s", step.Name)
		}
	}

	return nil
}

func (r *Runner) runStep(pipelineStep reader.PipelineStep) error {
	log.Info().Msgf("Running Pipeline Step: %s", pipelineStep.Name)

	// TODO: replace pipleline vars with original vars

	stdout, stderr, exitCode, err := utils.RunCommand("/bin/bash", "-c", pipelineStep.Cmd)
	if err != nil {
		log.Error().Err(err).Msgf("failed to execute step: %s", pipelineStep.Name)
		return err
	}

	var pipelineResult PipelineResult
	if pipelineStep.Store {
		pipelineResult = PipelineResult{
			Stdout:   stdout,
			Stderr:   stderr,
			ExitCode: exitCode,
		}
	} else {
		pipelineResult = PipelineResult{}
	}

	r.Results = append(r.Results, pipelineResult)

	if !pipelineStep.IgnoreFailure && exitCode != 0 {
		err := fmt.Errorf("step returned exit code %d instead of 0", exitCode)
		return err
	}

	return nil
}
