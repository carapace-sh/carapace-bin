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
	carapace.Gen(uninstallCmd).Standalone()

	uninstallCmd.Flags().String("using", "", "Either an explicit version, or a filename with the version written in it")
	rootCmd.AddCommand(uninstallCmd)

	carapace.Gen(uninstallCmd).FlagCompletion(carapace.ActionMap{
		"using": action.ActionLocalVersions(),
	})

	carapace.Gen(uninstallCmd).PositionalAnyCompletion(
		action.ActionInstalledVersions().Filter("system"),
	)
}
