package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var containerRegistry_tagCmd = &cobra.Command{
	Use:     "tag <command> [flags]",
	Short:   "Manage container registry tags.",
	Aliases: []string{"tags"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(containerRegistry_tagCmd).Standalone()

	containerRegistryCmd.AddCommand(containerRegistry_tagCmd)
}
