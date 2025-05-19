package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var utils_semverCmd = &cobra.Command{
	Use:   "semver",
	Short: "Work with semantic version strings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(utils_semverCmd).Standalone()
	utilsCmd.AddCommand(utils_semverCmd)
}
