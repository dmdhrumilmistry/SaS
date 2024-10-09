package main

import (
	"flag"

	_ "github.com/dmdhrumilmistry/sas/pkg/logging"
	"github.com/dmdhrumilmistry/sas/pkg/reader"
	"github.com/dmdhrumilmistry/sas/pkg/runner"
	"github.com/rs/zerolog/log"
)

func main() {
	file := flag.String("f", "", "path to pipeline file")
	flag.Parse()

	if *file == "" {
		log.Fatal().Msg("pipeline file path is required. for more info use -h flag")
	}

	r := reader.NewReader(*file)
	if err := r.Load(); err != nil {
		log.Error().Err(err).Msg("failed to load file")
	}

	rn := runner.NewRunner(r.Pipeline)
	if err := rn.Run(); err != nil {
		log.Error().Err(err).Msg("failed to execute pipeline file")
	}

	log.Print(r.Pipeline)
	log.Print(rn.Results)
}
