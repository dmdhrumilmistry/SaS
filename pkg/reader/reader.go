package reader

import (
	"os"
	"strings"

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

func (r *Reader) loadFileVar() error {
	if r.Pipeline.FileVariable.Separator == "" {
		r.Pipeline.FileVariable.Separator = "\n"
	}

	varsByte, err := os.ReadFile(r.Pipeline.FileVariable.Path)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to load %s var from file %s", r.Pipeline.FileVariable.Name, r.Pipeline.FileVariable.Path)
	}

	r.Pipeline.FileVariable.Values = strings.Split(string(varsByte), r.Pipeline.FileVariable.Separator)

	return nil
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

	if err := r.loadFileVar(); err != nil {
		log.Error().Err(err).Msgf("failed to load vars from file %s", r.filePath)
		return err
	}

	return nil
}
