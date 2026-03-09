package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/dnf_completer/cmd/action"
	"github.com/spf13/cobra"
)

var downgradeCmd = &cobra.Command{
	Use:     "downgrade [options] <package-spec>...",
	Aliases: []string{"dg"},
	Short:   "downgrade packages",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(downgradeCmd).Standalone()

	downgradeCmd.Flags().Bool("allow-downgrade", false, "enable downgrade of dependencies")
	downgradeCmd.Flags().Bool("allowerasing", false, "allow removing of installed packages")
	downgradeCmd.Flags().Bool("downloadonly", false, "download packages without executing transaction")
	downgradeCmd.Flags().Bool("no-allow-downgrade", false, "disable downgrade of dependencies")
	downgradeCmd.Flags().Bool("skip-broken", false, "resolve dependency problems")
	downgradeCmd.Flags().Bool("skip-unavailable", false, "allow skipping packages")

	rootCmd.AddCommand(downgradeCmd)

	carapace.Gen(downgradeCmd).PositionalAnyCompletion(
		action.ActionInstalledPackages(downgradeCmd),
	)
}
