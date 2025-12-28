package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/android/pkg_completer/cmd/action"
	"github.com/spf13/cobra"
)

var filesCmd = &cobra.Command{
	Use:   "files",
	Short: "Show all files installed by packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(filesCmd).Standalone()

	rootCmd.AddCommand(filesCmd)

	carapace.Gen(filesCmd).PositionalAnyCompletion(
		action.ActionInstalledPackages(),
	)
}
