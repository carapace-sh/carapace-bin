package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shellAlias_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Show the command for a shell alias",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shellAlias_getCmd).Standalone()

	shellAlias_getCmd.Flags().BoolP("help", "h", false, "Print help")
	shellAliasCmd.AddCommand(shellAlias_getCmd)
}
