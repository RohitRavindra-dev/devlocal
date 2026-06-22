package config

type DevlocalConfigYaml struct {
	Version  int      `yaml:"version"`
	Overlook []string `yaml:"overlook"`
	Patches  []string `yaml:"patches"`
}
