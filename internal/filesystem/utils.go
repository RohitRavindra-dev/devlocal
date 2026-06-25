package filesystem

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/RohitRavindra-dev/devlocal/internal/config"
	"gopkg.in/yaml.v3"
)

func FileExists(path string) (bool, error) {
	info, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false, fmt.Errorf("%s does not exist", path)
	}

	if err != nil {
		return false, err
	}

	if info.IsDir() {
		return false, fmt.Errorf("%s is a directory, expected a file", path)
	}

	return true, nil
}

func LoadDevlocalConfig() (*config.DevlocalConfigYaml, error) {
	data, err := os.ReadFile(filepath.Join(config.PROJECT_ROOT, config.CONFIG_FILE_NAME))

	if err != nil {
		return nil, err
	}

	var cfg config.DevlocalConfigYaml

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
