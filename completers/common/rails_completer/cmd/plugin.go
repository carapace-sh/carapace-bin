package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rails_completer/cmd/common"
	"github.com/spf13/cobra"
)

var pluginCmd = &cobra.Command{
	Use:   "plugin",
	Short: "Create a new Rails plugin or engine",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pluginCmd).Standalone()

	carapace.Gen(pluginCmd).PositionalCompletion(
		carapace.ActionValues().Usage("plugin name"),
	)

	pluginCmd.Flags().String("dummy_path", "test/dummy", "Path for dummy test application")
	pluginCmd.Flags().Bool("full", false, "Create a full engine")
	pluginCmd.Flags().BoolP("help", "h", false, "Show help")
	pluginCmd.Flags().Bool("mountable", false, "Create a mountable engine")
	pluginCmd.Flags().Bool("skip-gemfile-entry", false, "Skip adding entry to parent Gemfile")
	pluginCmd.Flags().Bool("skip-gemspec", false, "Skip gemspec file")

	common.AddAppGeneratorFlags(pluginCmd)
}
