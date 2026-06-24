package apply

import (
	"github.com/RohitRavindra-dev/devlocal/internal/filesystem"
	"github.com/RohitRavindra-dev/devlocal/internal/git"
)

func Run() error {
	config, err := filesystem.LoadDevlocalConfig()

	if err != nil {
		return err
	}

	//TODO
	for _, file := range config.Overlook {
		git.SkipWorkTree()

	}
}
