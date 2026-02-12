package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/atuin"
	"github.com/spf13/cobra"
)

var scripts_deleteCmd = &cobra.Command{
	Use:     "delete",
	Short:   "",
	Aliases: []string{"rm"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scripts_deleteCmd).Standalone()

	scripts_deleteCmd.Flags().BoolP("force", "f", false, "")
	scripts_deleteCmd.Flags().BoolP("help", "h", false, "Print help")
	scriptsCmd.AddCommand(scripts_deleteCmd)

	carapace.Gen(scripts_deleteCmd).PositionalCompletion(
		atuin.ActionScripts(),
	)
}
