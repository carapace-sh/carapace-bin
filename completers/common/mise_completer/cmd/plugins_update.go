// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).
// Completions below the flag definitions were wired by hand.

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var plugins_updateCmd = &cobra.Command{
	Use:     "update",
	Short:   "Update installed plugins",
	Aliases: []string{"up", "upgrade"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugins_updateCmd).Standalone()

	plugins_updateCmd.Flags().BoolP("help", "h", false, "Print help")
	plugins_updateCmd.Flags().StringP("jobs", "j", "", "Number of plugins to update in parallel (default: the `jobs` setting)")
	pluginsCmd.AddCommand(plugins_updateCmd)

	carapace.Gen(plugins_updateCmd).PositionalAnyCompletion(
		mise.ActionPlugins(),
	)
}
