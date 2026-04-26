package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:   "up [n]",
	Short: "Check out a branch further up in the stack (further from the trunk)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upCmd).Standalone()

	rootCmd.AddCommand(upCmd)

	// TODO positional completion
}
