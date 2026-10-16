package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var plugins_linkCmd = &cobra.Command{
	Use:   "link",
	Short: "Symlink a plugin into mise",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugins_linkCmd).Standalone()

	plugins_linkCmd.Flags().BoolP("force", "f", false, "Overwrite existing symlink")
	pluginsCmd.AddCommand(plugins_linkCmd)

	carapace.Gen(plugins_linkCmd).PositionalCompletion(
		carapace.ActionValues(),
		carapace.ActionDirectories(),
	)
}
