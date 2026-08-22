package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/asdf_completer/cmd/action"
	"github.com/spf13/cobra"
)

var reshimCmd = &cobra.Command{
	Use:   "reshim",
	Short: "Recreate shims for version of a package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reshimCmd).Standalone()

	rootCmd.AddCommand(reshimCmd)

	carapace.Gen(reshimCmd).PositionalCompletion(
		action.ActionPlugins(),
		action.ActionInstalledVersions(""),
	)
}
