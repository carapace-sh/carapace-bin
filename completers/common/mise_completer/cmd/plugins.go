package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pluginsCmd = &cobra.Command{
	Use:     "plugins",
	Short:   "Manage plugins",
	Aliases: []string{"p"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pluginsCmd).Standalone()

	rootCmd.AddCommand(pluginsCmd)
}
