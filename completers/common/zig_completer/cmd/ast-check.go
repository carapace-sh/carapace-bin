package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var astCheckCmd = &cobra.Command{
	Use:   "ast-check",
	Short: "Look for simple compile errors in any set of files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(astCheckCmd).Standalone()

	rootCmd.AddCommand(astCheckCmd)

	carapace.Gen(astCheckCmd).PositionalAnyCompletion(
		carapace.ActionFiles(".zig"),
	)
}
