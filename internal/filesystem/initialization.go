package filesystem

import (
	"fmt"
	"path/filepath"

	"os"

	"github.com/RohitRavindra-dev/devlocal/internal/config"

	"gopkg.in/yaml.v3"
)

func createRootDirectory() error {
	_, err := os.Stat(config.PROJECT_ROOT)
	if err == nil {
		return fmt.Errorf("%s already exists, please run cleanup", config.PROJECT_ROOT)
	}

	if !os.IsNotExist(err) {
		return err
	}

	return os.Mkdir(config.PROJECT_ROOT, 0755)
}

func createConfigFile() error {
	err := os.WriteFile(
		filepath.Join(config.PROJECT_ROOT, config.CONFIG_FILE_NAME),
		[]byte(""),
		0644,
	)
	return err
}

func seedYamlConfigFile() error {
	initYamlConfig := config.DevlocalConfigYaml{
		Version:  1,
		Overlook: []string{},
		Patches:  []string{},
	}

	data, err := yaml.Marshal(initYamlConfig)

	if err != nil {
		return err
	}

	return os.WriteFile(filepath.Join(config.PROJECT_ROOT, config.CONFIG_FILE_NAME), data, 0644)
}

func InitilizeDevLocalFilesystem() error {
	// try to create root directory first
	if err := createRootDirectory(); err != nil {
		return err
	}

	// try to create config file
	if err := createConfigFile(); err != nil {
		return err
	}

	if err := seedYamlConfigFile(); err != nil {
		return err
	}
	return nil
}
