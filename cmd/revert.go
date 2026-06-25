package cmd

import (
	"fmt"

	"github.com/RohitRavindra-dev/devlocal/internal/filesystem"
	"github.com/RohitRavindra-dev/devlocal/internal/service/revert"
	"github.com/spf13/cobra"
)

var revertCmd = &cobra.Command{
	Use:   "revert",
	Short: "Revert .devlocal changes that reverts overlooks and dev patches",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("[Started] to revert devlocal changes")

		return filesystem.ValidateDevLocalFilesystem()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return revert.Run()
	},
}

func init() {
	rootCmd.AddCommand(revertCmd)
}
