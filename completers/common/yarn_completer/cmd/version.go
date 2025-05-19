package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Apply a new version to the current package",
	GroupID: "general",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionCmd).Standalone()

	versionCmd.Flags().BoolP("deferred", "d", false, "Prepare the version to be bumped during the next release cycle")
	versionCmd.Flags().BoolP("immediate", "i", false, "Bump the version immediately")
	rootCmd.AddCommand(versionCmd)

	carapace.Gen(versionCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"major", "the first number from the semver range will be increased (X.0.0)",
			"minor", "the second number from the semver range will be increased (0.X.0)",
			"patch", "the third number from the semver range will be increased (0.0.X)",
			"pre", "when prefixed with 'pre' a -0 suffix will be set (0.0.0-0)", // TODO prefix completion
			"prerelease", "the suffix will be increased (0.0.0-X)",
			"decline", "the nonce will be increased for yarn version check to pass without version bump",
		),
	)
}
