package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_xmasCmd = &cobra.Command{
	Use:   "xmas",
	Short: "Not implemented in pnpm. Use the npm CLI directly",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_xmasCmd).Standalone()

	helpCmd.AddCommand(help_xmasCmd)
}
