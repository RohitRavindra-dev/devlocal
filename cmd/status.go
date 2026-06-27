package cmd

import (
	"github.com/RohitRavindra-dev/devlocal/internal/filesystem"
	"github.com/RohitRavindra-dev/devlocal/internal/service/status"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "View the current status of devlocal setup",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return filesystem.ValidateDevLocalFilesystem()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return status.Run()
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
