package filesystem

import (
	"fmt"
	"os"

	"github.com/RohitRavindra-dev/devlocal/internal/config"
)

func TearDownDevLocalFilesystem() error {
	fmt.Printf("Removing %s and all underlying evidence\n", config.PROJECT_ROOT)
	return os.RemoveAll(config.PROJECT_ROOT)
}
