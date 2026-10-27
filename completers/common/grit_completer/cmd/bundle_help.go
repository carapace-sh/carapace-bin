package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bundle_helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Print this message or the help of the given subcommand(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bundle_helpCmd).Standalone()

	bundleCmd.AddCommand(bundle_helpCmd)

	carapace.Gen(bundle_helpCmd).PositionalCompletion(
		carapace.ActionCommands(bundleCmd),
	)

	carapace.Gen(bundle_helpCmd).PositionalAnyCompletion(
		carapace.ActionCommands(bundleCmd),
	)
}
