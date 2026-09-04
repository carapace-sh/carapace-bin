package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/windows/scoop_completer/cmd/action"
	"github.com/spf13/cobra"
)

var virustotalCmd = &cobra.Command{
	Use:   "virustotal",
	Short: "look for app's hash or url on virustotal.com",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(virustotalCmd).Standalone()
	virustotalCmd.Flags().BoolP("all", "a", false, "check for all installed apps")
	virustotalCmd.Flags().BoolP("no-depends", "n", false, "don't check dependencies")
	virustotalCmd.Flags().BoolP("no-update-scoop", "u", false, "don't update Scoop before checking")
	virustotalCmd.Flags().BoolP("passthru", "p", false, "return reports as objects")
	virustotalCmd.Flags().BoolP("scan", "s", false, "submit download URL for analysis if no info exists")
	rootCmd.AddCommand(virustotalCmd)

	carapace.Gen(virustotalCmd).PositionalAnyCompletion(
		action.ActionInstalledApps(),
	)
}
