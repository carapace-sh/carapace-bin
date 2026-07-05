package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setenvCmd = &cobra.Command{
	Use:   "setenv",
	Short: "Set an environment variable in the specified domain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setenvCmd).Standalone()
	rootCmd.AddCommand(setenvCmd)
}
