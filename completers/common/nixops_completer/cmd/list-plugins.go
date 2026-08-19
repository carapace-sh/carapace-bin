package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ListPluginsCmd = &cobra.Command{
	Use:   "list-plugins",
	Short: "List Plugins",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ListPluginsCmd).Standalone()
	rootCmd.AddCommand(ListPluginsCmd)
}
