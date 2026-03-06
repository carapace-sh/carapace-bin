package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/dnf_completer/cmd/action"
	"github.com/spf13/cobra"
)

var reinstallCmd = &cobra.Command{
	Use:     "reinstall [options] <package-spec>...",
	Aliases: []string{"rei"},
	Short:   "reinstall packages",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reinstallCmd).Standalone()

	reinstallCmd.Flags().Bool("allowerasing", false, "allow removing of installed packages")
	reinstallCmd.Flags().Bool("skip-broken", false, "resolve dependency problems by removing problematic packages")
	reinstallCmd.Flags().Bool("skip-unavailable", false, "allow skipping packages")
	reinstallCmd.Flags().Bool("allow-downgrade", false, "enable downgrade of dependencies")
	reinstallCmd.Flags().Bool("no-allow-downgrade", false, "disable downgrade of dependencies")
	reinstallCmd.Flags().Bool("downloadonly", false, "download packages without executing transaction")

	rootCmd.AddCommand(reinstallCmd)

	carapace.Gen(reinstallCmd).PositionalAnyCompletion(
		action.ActionPackageSearch(reinstallCmd),
	)
}
