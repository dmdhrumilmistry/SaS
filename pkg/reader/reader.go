package reader

import (
	"os"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
)

type Reader struct {
	filePath string
	Pipeline Pipeline
}

func NewReader(filePath string) *Reader {
	return &Reader{
		filePath: filePath,
	}
}

func (r *Reader) Load() error {
	contents, err := os.ReadFile(r.filePath)
	if err != nil {
		log.Error().Err(err).Msgf("failed to load file %s contents", r.filePath)
		return err
	}

	if err := yaml.Unmarshal(contents, &r.Pipeline); err != nil {
		log.Error().Err(err).Msgf("failed to load yaml content from file %s", r.filePath)
		return err
	}

	return nil
}
