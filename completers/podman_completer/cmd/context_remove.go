package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var context_removeCmd = &cobra.Command{
	Use:     "remove [options] NAME",
	Short:   "Delete named destination",
	Aliases: []string{"rm"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(context_removeCmd).Standalone()

	context_removeCmd.Flags().BoolP("all", "a", false, "Remove all connections")
	context_removeCmd.Flags().BoolP("force", "f", false, "Ignored: for Docker compatibility")
	context_removeCmd.Flag("force").Hidden = true
	contextCmd.AddCommand(context_removeCmd)
}
