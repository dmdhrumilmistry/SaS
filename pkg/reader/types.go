package reader

type Pipeline struct {
	Name         string               `yaml:"name"`
	Config       PipelineConfig       `yaml:"config"`
	Vars         []PipelineVariable   `yaml:"vars"`
	Steps        []PipelineStep       `yaml:"pipeline"`
	FileVariable PipelineFileVariable `yaml:"file_var"`
}

type PipelineConfig struct {
	Workers int `yaml:"workers"`
}

type PipelineVariable struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
	Type  string `yaml:"type"`
}

type PipelineFileVariable struct {
	Name      string   `yaml:"name"`
	Path      string   `yaml:"path"`
	Separator string   `yaml:"separator"`
	Values    []string `yaml:"-"`
}

type PipelineStep struct {
	Name          string `yaml:"name"`
	Cmd           string `yaml:"cmd"`
	Store         bool   `yaml:"store"`
	IgnoreFailure bool   `yaml:"ignore_failure"`
}
