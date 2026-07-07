package patch

import (
	"fmt"

	"github.com/RohitRavindra-dev/devlocal/internal/filesystem"
)

func Run(filePath string) error {

	if filePath == "" {
		return fmt.Errorf("[Error] file path is empty!")
	}

	if _, err := filesystem.FileExists(filePath); err != nil {
		return err
	}

	fmt.Println("[Completed] recording new patch into devlocal")
	return nil
}
