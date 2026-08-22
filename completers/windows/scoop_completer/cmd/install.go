package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/windows/scoop_completer/cmd/action"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "install apps",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()
	installCmd.Flags().StringP("arch", "a", "", "use the specified architecture (32bit|64bit|arm64)")
	installCmd.Flags().BoolP("global", "g", false, "install the app globally")
	installCmd.Flags().BoolP("independent", "i", false, "don't install dependencies automatically")
	installCmd.Flags().BoolP("no-cache", "k", false, "don't use the download cache")
	installCmd.Flags().BoolP("no-update-scoop", "u", false, "don't update Scoop before installing")
	installCmd.Flags().BoolP("skip-hash-check", "s", false, "skip hash validation")
	rootCmd.AddCommand(installCmd)

	carapace.Gen(installCmd).FlagCompletion(carapace.ActionMap{
		"arch": carapace.ActionValues("32bit", "64bit", "arm64"),
	})

	carapace.Gen(installCmd).PositionalAnyCompletion(
		action.ActionInstalledApps(),
	)
}
