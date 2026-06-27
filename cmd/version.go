package cmd

import (
	"fmt"

	"github.com/RohitRavindra-dev/devlocal/internal/config"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[devlocal]", config.Version, config.Commit)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
