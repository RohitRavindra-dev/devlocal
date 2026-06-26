package filesystem

import (
	"os"
	"path/filepath"

	"github.com/RohitRavindra-dev/devlocal/internal/config"
	"gopkg.in/yaml.v3"
)

func exists(path string) (bool, error) {
	_, err := os.Stat(path)

	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

func LoadDevlocalConfig() (*config.DevlocalConfigYaml, error) {
	data, err := os.ReadFile(filepath.Join(config.PROJECT_ROOT, config.CONFIG_FILE_NAME))
	var x
	if err != nil {
		return nil, err
	}

	var cfg config.DevlocalConfigYaml

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
