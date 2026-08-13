package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/vercel_completer/cmd/action"
	"github.com/spf13/cobra"
)

var alias_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Create a new alias",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(alias_setCmd).Standalone()

	aliasCmd.AddCommand(alias_setCmd)

	carapace.Gen(alias_setCmd).PositionalCompletion(
		action.ActionDeployments(alias_setCmd),
		carapace.ActionValues(),
	)
}
