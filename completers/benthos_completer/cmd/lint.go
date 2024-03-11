package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lintCmd = &cobra.Command{
	Use:   "lint",
	Short: "Parse Benthos configs and report any linting errors",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lintCmd).Standalone()

	lintCmd.Flags().Bool("deprecated", false, "Print linting errors for the presence of deprecated fields.")
	lintCmd.Flags().Bool("labels", false, "Print linting errors when components do not have labels.")
	lintCmd.Flags().Bool("skip-env-var-check", false, "Do not produce lint errors when environment interpolations exist without defaults within configs but aren't defined.")
	rootCmd.AddCommand(lintCmd)

	carapace.Gen(lintCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
