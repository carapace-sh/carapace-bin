package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_pluginsCmd = &cobra.Command{
	Use:   "plugins",
	Short: "Manage package manager plugins declared in `[bootstrap.plugins]`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_pluginsCmd).Standalone()

	bootstrap_pluginsCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrapCmd.AddCommand(bootstrap_pluginsCmd)
}
