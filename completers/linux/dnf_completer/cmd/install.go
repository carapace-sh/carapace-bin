package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/dnf_completer/cmd/action"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:     "install [options] <package-spec>|@<group-spec>|@<environment-spec>...",
	Aliases: []string{"in"},
	Short:   "install packages",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()

	installCmd.Flags().Bool("allowerasing", false, "allow removing of installed packages")
	installCmd.Flags().Bool("skip-broken", false, "resolve dependency problems by removing problematic packages")
	installCmd.Flags().Bool("skip-unavailable", false, "allow skipping packages that are not available")
	installCmd.Flags().Bool("allow-downgrade", false, "enable downgrade of dependencies")
	installCmd.Flags().Bool("no-allow-downgrade", false, "disable downgrade of dependencies")
	installCmd.Flags().Bool("downloadonly", false, "download packages without executing transaction")

	rootCmd.AddCommand(installCmd)

	carapace.Gen(installCmd).PositionalAnyCompletion(
		action.ActionPackageSearch(installCmd),
	)
}
