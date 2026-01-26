package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var generateLockfileCmd = &cobra.Command{
	Use:   "generate-lockfile",
	Short: "Generate the lockfile for a package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generateLockfileCmd).Standalone()

	generateLockfileCmd.Flags().BoolP("help", "h", false, "Print help")
	generateLockfileCmd.Flags().Bool("ignore-rust-version", false, "Ignore `rust-version` specification in packages")
	generateLockfileCmd.Flags().String("lockfile-path", "", "Path to Cargo.lock (unstable)")
	generateLockfileCmd.Flags().String("manifest-path", "", "Path to Cargo.toml")
	generateLockfileCmd.Flags().String("publish-time", "", "Latest publish time allowed for registry packages (unstable)")
	rootCmd.AddCommand(generateLockfileCmd)

	carapace.Gen(generateLockfileCmd).FlagCompletion(carapace.ActionMap{
		"lockfile-path": carapace.ActionFiles(),
		"manifest-path": carapace.ActionFiles(),
	})
}
