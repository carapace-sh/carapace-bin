package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var plugins_lsRemoteCmd = &cobra.Command{
	Use:     "ls-remote",
	Short:   "List registry tools that have an asdf or vfox plugin backend",
	Aliases: []string{"list-remote", "list-all"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugins_lsRemoteCmd).Standalone()

	plugins_lsRemoteCmd.Flags().BoolP("help", "h", false, "Print help")
	plugins_lsRemoteCmd.Flags().Bool("only-names", false, "Only show the name of each plugin, without the `*` marking installed plugins")
	plugins_lsRemoteCmd.Flags().BoolP("urls", "u", false, "Show the plugin source for each shorthand, e.g. vfox:jdx/vfox-poetry")
	pluginsCmd.AddCommand(plugins_lsRemoteCmd)
}
