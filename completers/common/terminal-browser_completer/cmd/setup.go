package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Configure terminals and agents so terminal-browser works best",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setupCmd).Standalone()

	rootCmd.AddCommand(setupCmd)
}
