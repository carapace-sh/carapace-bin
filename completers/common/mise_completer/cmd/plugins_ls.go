package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var plugins_lsCmd = &cobra.Command{
	Use:     "ls",
	Short:   "List installed plugins",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugins_lsCmd).Standalone()

	plugins_lsCmd.Flags().Bool("core", false, "Show core plugins")
	plugins_lsCmd.Flags().BoolP("urls", "u", false, "Show git urls")
	pluginsCmd.AddCommand(plugins_lsCmd)
}
