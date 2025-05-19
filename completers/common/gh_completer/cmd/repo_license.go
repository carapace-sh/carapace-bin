package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repo_licenseCmd = &cobra.Command{
	Use:     "license <command>",
	Short:   "Explore repository licenses",
	GroupID: "Targeted commands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_licenseCmd).Standalone()

	repoCmd.AddCommand(repo_licenseCmd)
}
