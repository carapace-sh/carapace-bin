package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help___internalCmd = &cobra.Command{
	Use:    "__internal",
	Short:  "Internal subcommands, not for direct use by users",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help___internalCmd).Standalone()

	helpCmd.AddCommand(help___internalCmd)
}
