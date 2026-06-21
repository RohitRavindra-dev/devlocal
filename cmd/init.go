package cmd

import (
	"github.com/RohitRavindra-dev/devlocal/internal/filesystem"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize devlocal in this project workspace",
	RunE: func(cmd *cobra.Command, args []string) error {
		return filesystem.InitilizeDevLocalFilesystem()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
