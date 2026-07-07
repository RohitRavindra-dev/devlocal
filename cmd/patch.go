package cmd

import (
	"github.com/spf13/cobra"
)

var patchCmd = &cobra.Command{
	Use:   "patch",
	Short: "Patch(es) management root command.",
}

func init() {
	rootCmd.AddCommand(patchCmd)
}
