package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var utils_semver_checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check version constraints",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(utils_semver_checkCmd).Standalone()
	utils_semver_checkCmd.Flags().StringP("constraint", "c", "", "the version constraint to check against")
	utils_semver_checkCmd.Flags().StringP("version", "v", "", "the semver version string to check")
	utils_semverCmd.AddCommand(utils_semver_checkCmd)
}
