package runner

import (
	"fmt"
	"time"

	"github.com/dmdhrumilmistry/sas/pkg/reader"
	"github.com/dmdhrumilmistry/sas/pkg/utils"
	"github.com/rs/zerolog/log"
)

func NewRunner(pipeline reader.Pipeline) *Runner {
	return &Runner{
		Pipeline:  pipeline,
		scanQueue: make(chan string),
	}
}

func (r *Runner) InitChannels() {
	log.Print(len(r.Pipeline.FileVariable.Values))
	r.scanQueue = make(chan string, len(r.Pipeline.FileVariable.Values))
}

func (r *Runner) CloseChannels() {
	close(r.scanQueue)
}

func (r *Runner) AddJobs() {
	for _, fileVar := range r.Pipeline.FileVariable.Values {
		r.scanQueue <- fileVar
	}

	r.CloseChannels()
}

func (r *Runner) RunWorker(workerNo int) {
	defer r.wg.Done()
	workerResult := WorkerResult{}

	for fileVarValue := range r.scanQueue {
		log.Info().Msgf("[Worker %d] Running job for %s -> %s", workerNo, r.Pipeline.FileVariable.Name, fileVarValue)
		workerResult.FileVarValue = fileVarValue

		pipelineResult, err := r.RunPipeline(fileVarValue)
		workerResult.PipelineResult = pipelineResult

		if err != nil {
			log.Error().Err(err).Msgf("[Worked %d] failed to execute pipeline for %s -> %s", workerNo, r.Pipeline.FileVariable.Name, fileVarValue)
		}

		r.mu.Lock()
		r.Results = append(r.Results, workerResult)
		r.mu.Unlock()

		log.Info().Msgf("[Worked %d] executed pipeline for %s -> %s", workerNo, r.Pipeline.FileVariable.Name, fileVarValue)
	}
}

func (r *Runner) RunWorkers() {
	varsLen := len(r.Pipeline.FileVariable.Values)
	if varsLen < r.Pipeline.Config.Workers {
		r.Pipeline.Config.Workers = varsLen
	}

	r.InitChannels()

	startTime := time.Now()
	for i := 1; i <= r.Pipeline.Config.Workers; i++ {
		r.wg.Add(1)
		go r.RunWorker(i)
	}

	r.AddJobs()

	r.wg.Wait()
	endTime := time.Now()
	totalTime := endTime.Sub(startTime)

	log.Info().Msgf("Completed running pipelines in %v", totalTime)
}

func (r *Runner) RunPipeline(fileVarValue string) (PipelineResult, error) {
	pipelineResult := PipelineResult{}

	for _, step := range r.Pipeline.Steps {
		stepResult, err := r.runStep(step, fileVarValue)
		if err != nil {
			return pipelineResult, err
		}
		pipelineResult.Results = append(pipelineResult.Results, stepResult)
		log.Info().Msgf("Ran step successfully: %s", step.Name)
	}

	return pipelineResult, nil
}

func (r *Runner) runStep(pipelineStep reader.PipelineStep, fileVar string) (PipelineStepResult, error) {
	pipelineResult := PipelineStepResult{}

	log.Info().Msgf("Running Pipeline Step: %s", pipelineStep.Name)

	// TODO: replace pipleline vars with original vars
	if fileVar == "" {
		log.Warn().Msgf("Value of %s is empty", r.Pipeline.FileVariable.Name)
	}

	stdout, stderr, exitCode, err := utils.RunCommand("/bin/bash", "-c", pipelineStep.Cmd)
	if err != nil {
		log.Error().Err(err).Msgf("failed to execute step: %s", pipelineStep.Name)
		return pipelineResult, err
	}

	if pipelineStep.Store {
		pipelineResult.Stdout = stdout
		pipelineResult.Stderr = stderr
		pipelineResult.ExitCode = exitCode
	}

	if !pipelineStep.IgnoreFailure && exitCode != 0 {
		err := fmt.Errorf("step returned exit code %d instead of 0", exitCode)
		return pipelineResult, err
	}

	return pipelineResult, nil
}
