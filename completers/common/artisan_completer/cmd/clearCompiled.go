package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clearCompiledCmd = &cobra.Command{
	Use:   "clear-compiled",
	Short: "Remove the compiled class file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clearCompiledCmd).Standalone()

	rootCmd.AddCommand(clearCompiledCmd)
}
