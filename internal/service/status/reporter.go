package status

import (
	"fmt"

	"github.com/RohitRavindra-dev/devlocal/internal/config"
	"github.com/RohitRavindra-dev/devlocal/internal/filesystem"
	"github.com/RohitRavindra-dev/devlocal/internal/git"
)

func displayConfigVersion(version int) {
	fmt.Printf("[Version] %d \n", version)
}

func displayOverlookStatus(fileNames []string) {
	fmt.Println("[Overlooked files]")
	if len(fileNames) == 0 {
		fmt.Println("\t- No files registed in config for overlooking!")
		return
	}
	for _, filename := range fileNames {
		exists, err := filesystem.FileExists(filename)
		overlookStatus := "✗"
		overlookComment := "Being tracked"
		if exists {

			if isOverlooked, _ := git.IsSkipWorktree(filename); isOverlooked {
				overlookStatus = "✓"
				overlookComment = "Overlooked"
			}

		} else {
			overlookComment = err.Error()
		}

		fmt.Printf("\t%s %s \n\t\t- %s\n", overlookStatus, filename, overlookComment)

	}
}

func displayPatchesStatus(patches []string) {
	fmt.Println("[Patches]")
	if len(patches) == 0 {
		fmt.Println("\t- No patches registed in config for patching!")
		return
	}
}

func displayDevlocalConfig(config *config.DevlocalConfigYaml) error {
	fmt.Println("\n===============================")
	fmt.Println("|   devlocal [Status]         |")
	fmt.Printf("===============================\n\n")
	displayConfigVersion(config.Version)
	fmt.Println()
	displayOverlookStatus(config.Overlook)
	fmt.Println()
	displayPatchesStatus(config.Patches)
	fmt.Printf("===============================\n\n")

	return nil
}

func Run() error {

	//load config
	config, err := filesystem.LoadDevlocalConfig()
	if err != nil {
		return err
	}

	if err := displayDevlocalConfig(config); err != nil {
		return err
	}

	return nil
}
