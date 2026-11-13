package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pluginsCmd = &cobra.Command{
	Use:   "plugins",
	Short: "List installed plugins",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pluginsCmd).Standalone()

	pluginsCmd.Flags().Bool("all", false, "Include built-in default plugins")
	pluginsCmd.Flags().String("hook", "", "Filter for plugins that implement this hook")
	rootCmd.AddCommand(pluginsCmd)
}
