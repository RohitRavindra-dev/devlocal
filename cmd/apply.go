package cmd

import (
	"fmt"

	"github.com/RohitRavindra-dev/devlocal/internal/filesystem"
	"github.com/spf13/cobra"

	"github.com/RohitRavindra-dev/devlocal/internal/orchestration/apply"
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
		return apply.Run()
	},
}

func init() {
	rootCmd.AddCommand(applyCmd)
}
