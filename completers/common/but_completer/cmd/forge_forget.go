package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forge_forgetCmd = &cobra.Command{
	Use:   "forget",
	Short: "Forget a previously authenticated forge account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forge_forgetCmd).Standalone()

	forge_forgetCmd.Flags().BoolP("help", "h", false, "Print help")
	forgeCmd.AddCommand(forge_forgetCmd)
}
