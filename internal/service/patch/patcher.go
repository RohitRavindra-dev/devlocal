package patch

import (
	"fmt"

	"github.com/RohitRavindra-dev/devlocal/internal/filesystem"
)

func Run(path string) error {

	if path == "" {
		return fmt.Errorf("[Error] path is empty!")
	}

	if _, err := filesystem.PathExists(path); err != nil {
		return err
	}

	fmt.Println("[Completed] recording new patch into devlocal")
	return nil
}
