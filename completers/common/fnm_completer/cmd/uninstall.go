package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/fnm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall a Node.js version",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)

	addCommonFlags(uninstallCmd)

	uninstallCmd.Flags().String("using", "", "Either an explicit version, or a filename with the version written in it")

	carapace.Gen(uninstallCmd).Standalone()

	carapace.Gen(uninstallCmd).FlagCompletion(carapace.ActionMap{
		"using": action.ActionLocalVersions(),
	})

	carapace.Gen(unaliasCmd).PositionalAnyCompletion(
		action.ActionInstalledVersions(),
	)
}
