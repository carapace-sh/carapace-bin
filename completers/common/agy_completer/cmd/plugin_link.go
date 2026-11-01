package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pluginLinkCmd = &cobra.Command{
	Use:   "link",
	Short: "Generate link to a marketplace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pluginLinkCmd).Standalone()
	pluginCmd.AddCommand(pluginLinkCmd)

	carapace.Gen(pluginLinkCmd).PositionalCompletion(
		carapace.ActionValues().Usage("marketplace name"),
		carapace.ActionValues().Usage("target"),
	)
}
