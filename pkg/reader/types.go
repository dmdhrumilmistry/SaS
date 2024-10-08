package reader

type Pipeline struct {
	Name   string             `yaml:"name"`
	Config PipelineConfig     `yaml:"config"`
	Vars   []PipelineVariable `yaml:"vars"`
	Steps  []PipelineStep     `yaml:"pipeline"`
}

type PipelineConfig struct {
	Workers int
}

type PipelineVariable struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
	Type  string `yaml:"type"`
}

type PipelineStep struct {
	Name          string `yaml:"name"`
	Cmd           string `yaml:"cmd"`
	Store         bool   `yaml:"store"`
	IgnoreFailure bool   `yaml:"ignore_failure"`
}
