package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forge_help_forgetCmd = &cobra.Command{
	Use:   "forget",
	Short: "Forget a previously authenticated forge account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forge_help_forgetCmd).Standalone()

	forge_helpCmd.AddCommand(forge_help_forgetCmd)
}
