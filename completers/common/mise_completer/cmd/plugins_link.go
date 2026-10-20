// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).
// Completions below the flag definitions were wired by hand.

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var plugins_linkCmd = &cobra.Command{
	Use:     "link",
	Short:   "Link a local plugin directory into mise for development",
	Aliases: []string{"ln"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugins_linkCmd).Standalone()

	plugins_linkCmd.Flags().BoolP("force", "f", false, "Replace an existing plugin with the same name")
	plugins_linkCmd.Flags().BoolP("help", "h", false, "Print help")
	pluginsCmd.AddCommand(plugins_linkCmd)

	carapace.Gen(plugins_linkCmd).PositionalCompletion(
		carapace.ActionValues(),
		carapace.ActionDirectories(),
	)
}
