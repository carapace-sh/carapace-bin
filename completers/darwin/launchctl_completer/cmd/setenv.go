package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setenvCmd = &cobra.Command{
	Use:   "setenv",
	Short: "Sets the specified environment variables",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	carapace.Gen(setenvCmd).Standalone()
	rootCmd.AddCommand(setenvCmd)
}
