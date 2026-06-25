package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var unsetenvCmd = &cobra.Command{
	Use:   "unsetenv",
	Short: "Unsets the specified environment variables",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	carapace.Gen(unsetenvCmd).Standalone()
	rootCmd.AddCommand(unsetenvCmd)
}
