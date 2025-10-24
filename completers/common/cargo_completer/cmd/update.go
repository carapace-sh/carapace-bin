package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:     "update",
	Short:   "Update dependencies as recorded in the local lock file",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groupFor("update"),
}

func init() {
	carapace.Gen(updateCmd).Standalone()

	updateCmd.Flags().Bool("aggressive", false, "Force updating all dependencies of [SPEC]... as well")
	updateCmd.Flags().BoolP("breaking", "b", false, "Update [SPEC] to latest SemVer-breaking version (unstable)")
	updateCmd.Flags().BoolP("dry-run", "n", false, "Don't actually write the lockfile")
	updateCmd.Flags().BoolP("help", "h", false, "Print help")
	updateCmd.Flags().Bool("ignore-rust-version", false, "Ignore `rust-version` specification in packages")
	updateCmd.Flags().String("lockfile-path", "", "Path to Cargo.lock (unstable)")
	updateCmd.Flags().String("manifest-path", "", "Path to Cargo.toml")
	updateCmd.Flags().String("precise", "", "Update [SPEC] to exactly PRECISE")
	updateCmd.Flags().Bool("recursive", false, "Force updating all dependencies of [SPEC]... as well")
	updateCmd.Flags().BoolP("workspace", "w", false, "Only update the workspace packages")
	updateCmd.Flag("aggressive").Hidden = true
	rootCmd.AddCommand(updateCmd)

	carapace.Gen(updateCmd).FlagCompletion(carapace.ActionMap{
		"lockfile-path": carapace.ActionFiles(),
		"manifest-path": carapace.ActionFiles(),
	})
}
