package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "Cleanup devlocal setup",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Cleaning up ")
		//call revert
		// remove devlocal files
	},
}

func init() {
	rootCmd.AddCommand(cleanupCmd)
}
