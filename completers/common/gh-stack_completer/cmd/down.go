package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var downCmd = &cobra.Command{
	Use:   "down [n]",
	Short: "Check out a branch further down in the stack (closer to the trunk)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(downCmd).Standalone()

	rootCmd.AddCommand(downCmd)

	// TODO positional completion
}
