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
	generateLockfileCmd.Flags().Bool("ignore-rust-version", false, "Ignore `rust-version` specification in packages (unstable)")
	generateLockfileCmd.Flags().String("manifest-path", "", "Path to Cargo.toml")
	rootCmd.AddCommand(generateLockfileCmd)

	carapace.Gen(generateLockfileCmd).FlagCompletion(carapace.ActionMap{
		"manifest-path": carapace.ActionFiles(),
	})
}
