package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validates a composer.json and composer.lock",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(validateCmd).Standalone()

	validateCmd.Flags().Bool("check-lock", false, "Check if lock file is up to date (even when config.lock is false)")
	validateCmd.Flags().Bool("no-check-all", false, "Do not validate requires for overly strict/loose constraints")
	validateCmd.Flags().Bool("no-check-lock", false, "Do not check if lock file is up to date")
	validateCmd.Flags().Bool("no-check-publish", false, "Do not check for publish errors")
	validateCmd.Flags().Bool("no-check-version", false, "Do not report a warning if the version field is present")
	validateCmd.Flags().Bool("strict", false, "Return a non-zero exit code for warnings as well as errors")
	validateCmd.Flags().BoolP("with-dependencies", "A", false, "Also validate the composer.json of all installed dependencies")

	rootCmd.AddCommand(validateCmd)

	carapace.Gen(validateCmd).PositionalCompletion(
		carapace.ActionFiles(".json"),
	)
}
