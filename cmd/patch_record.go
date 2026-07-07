package cmd

import (
	"fmt"

	"github.com/RohitRavindra-dev/devlocal/internal/filesystem"
	"github.com/RohitRavindra-dev/devlocal/internal/service/patch"
	"github.com/spf13/cobra"
)

var patchRecordCmd = &cobra.Command{
	Use:   "record <file>",
	Short: "Record local modifications as a DevLocal patch",
	Args:  cobra.ExactArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("[Started] to record a new patch")
		// check for setup
		return filesystem.ValidateDevLocalFilesystem()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		file := args[0]
		return patch.Run(file)
	},
}

func init() {
	patchCmd.AddCommand(patchRecordCmd)
}
