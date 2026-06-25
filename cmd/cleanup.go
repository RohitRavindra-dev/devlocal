package cmd

import (
	"fmt"

	"github.com/RohitRavindra-dev/devlocal/internal/filesystem"
	"github.com/RohitRavindra-dev/devlocal/internal/service/revert"
	"github.com/spf13/cobra"
)

var cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "Cleanup devlocal setup",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("[Started] to revert devlocal changes")
		// check for setup
		return filesystem.ValidateDevLocalFilesystem()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Cleaning up setup")
		//call revert
		if err := revert.Run(); err != nil {
			fmt.Println("[Errored] while trying to cleanup devlocal setup")
			return err
		}
		// remove devlocal files
		fileTeardownError := filesystem.TearDownDevLocalFilesystem()
		if fileTeardownError != nil {
			return fileTeardownError
		}
		fmt.Println("[Completed] Cleanup, ciao!")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(cleanupCmd)
}
