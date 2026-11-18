package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/brew_completer/cmd/action"
	"github.com/spf13/cobra"
)

var unaliasCmd = &cobra.Command{
	Use:     "unalias",
	Short:   "Remove aliases",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unaliasCmd).Standalone()

	unaliasCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	unaliasCmd.Flags().BoolP("help", "h", false, "Show this message.")
	unaliasCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	unaliasCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(unaliasCmd)

	carapace.Gen(unaliasCmd).PositionalAnyCompletion(
		action.ActionAliases().FilterArgs(),
	)
}
