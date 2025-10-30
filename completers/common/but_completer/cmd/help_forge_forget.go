package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_forge_forgetCmd = &cobra.Command{
	Use:   "forget",
	Short: "Forget a previously authenticated forge account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_forge_forgetCmd).Standalone()

	help_forgeCmd.AddCommand(help_forge_forgetCmd)
}
