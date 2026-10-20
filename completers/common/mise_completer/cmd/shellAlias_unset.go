package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shellAlias_unsetCmd = &cobra.Command{
	Use:     "unset",
	Short:   "Remove a shell alias from the global config",
	Aliases: []string{"rm", "remove", "delete", "del"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shellAlias_unsetCmd).Standalone()

	shellAlias_unsetCmd.Flags().BoolP("help", "h", false, "Print help")
	shellAliasCmd.AddCommand(shellAlias_unsetCmd)
}
