package cmd

import (
	"fmt"

	"github.com/RohitRavindra-dev/devlocal/internal/filesystem"
	"github.com/spf13/cobra"
)

var cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "Cleanup devlocal setup",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Cleaning up setup")
		//call revert
		// remove devlocal files
		fileTeardownError := filesystem.TearDownDevLocalFilesystem()
		if fileTeardownError != nil {
			return fileTeardownError
		}
		fmt.Println("Cleanup completed, ciao!")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(cleanupCmd)
}
