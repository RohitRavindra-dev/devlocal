package filesystem

import (
	"fmt"
	"path/filepath"

	"os"

	"github.com/RohitRavindra-dev/devlocal/internal/configs"
)

func createRootDirectory() error {
	_, err := os.Stat(configs.PROJECT_ROOT)
	if err == nil {
		return fmt.Errorf("%s already exists, please run cleanup", configs.PROJECT_ROOT)
	}

	if !os.IsNotExist(err) {
		return err
	}

	return os.Mkdir(configs.PROJECT_ROOT, 0755)
}

func createConfigFile() error {
	err := os.WriteFile(
		filepath.Join(configs.PROJECT_ROOT, "config.yaml"),
		[]byte(""),
		0644,
	)
	return err
}

func seedYamlConfigFile() error {
	return nil
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
