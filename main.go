package main

import (
	_ "github.com/dmdhrumilmistry/sas/pkg/logging"
	"github.com/dmdhrumilmistry/sas/pkg/reader"
	"github.com/rs/zerolog/log"
)

func main() {
	r := reader.NewReader("pipeline.sample.yml")
	r.Load()

	log.Print(r.Pipeline)
}
