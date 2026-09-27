package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Not implemented in pnpm. Use the npm CLI directly",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_tokenCmd).Standalone()

	helpCmd.AddCommand(help_tokenCmd)
}
