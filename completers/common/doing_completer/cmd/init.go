package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "doing init [OPTIONS] [REFERENCE_ISSUE]",
	Short: "Create a .doing-cli-config file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initCmd).Standalone()

	initCmd.Flags().Bool("help", false, "Show this message and exit.")
	rootCmd.AddCommand(initCmd)

	// TODO positional completion
}
