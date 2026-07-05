package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var unsetenvCmd = &cobra.Command{
	Use:   "unsetenv",
	Short: "Unset an environment variable in the specified domain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unsetenvCmd).Standalone()
	rootCmd.AddCommand(unsetenvCmd)
}
