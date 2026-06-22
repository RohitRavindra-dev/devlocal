package cmd

import (
	"fmt"

	"github.com/RohitRavindra-dev/devlocal/internal/filesystem"
	"github.com/spf13/cobra"
)

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply .devlocal changes that makes the project ready for local developement",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Attempting to apply devlocal changes")
		// check for setup
		return filesystem.ValidateDevLocalFilesystem()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		// apply local changes
		// git worktree untrack
		return nil
	},
}

func init() {
	rootCmd.AddCommand(applyCmd)
}
