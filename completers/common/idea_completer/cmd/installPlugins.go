package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var installPluginsCmd = &cobra.Command{
	Use:   "installPlugins",
	Short: "Install plugins by plugin ID from JetBrains Marketplace or a custom plugin repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installPluginsCmd).Standalone()

	rootCmd.AddCommand(installPluginsCmd)
}
