package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/scoop"
	"github.com/spf13/cobra"
)

var alias_rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "remove an alias",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(alias_rmCmd).Standalone()
	aliasCmd.AddCommand(alias_rmCmd)

	carapace.Gen(alias_rmCmd).PositionalCompletion(
		scoop.ActionAliases(),
	)
}
