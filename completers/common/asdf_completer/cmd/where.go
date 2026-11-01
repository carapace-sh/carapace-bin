package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/asdf"
	"github.com/spf13/cobra"
)

var whereCmd = &cobra.Command{
	Use:   "where",
	Short: "Display install path for an installed or current version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(whereCmd).Standalone()

	rootCmd.AddCommand(whereCmd)

	carapace.Gen(whereCmd).PositionalCompletion(
		asdf.ActionInstalledPlugins(),
		asdf.ActionInstalledVersions(""),
	)
}
