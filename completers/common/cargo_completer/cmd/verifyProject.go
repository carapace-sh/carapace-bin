package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var verifyProjectCmd = &cobra.Command{
	Use:   "verify-project",
	Short: "Check correctness of crate manifest",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(verifyProjectCmd).Standalone()

	verifyProjectCmd.Flags().BoolP("help", "h", false, "Print help")
	verifyProjectCmd.Flags().String("manifest-path", "", "Path to Cargo.toml")
	rootCmd.AddCommand(verifyProjectCmd)

	carapace.Gen(verifyProjectCmd).FlagCompletion(carapace.ActionMap{
		"manifest-path": carapace.ActionFiles(),
	})
}
