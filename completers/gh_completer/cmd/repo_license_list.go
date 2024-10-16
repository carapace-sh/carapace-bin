package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repo_license_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List common repository licenses",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_license_listCmd).Standalone()

	repo_licenseCmd.AddCommand(repo_license_listCmd)
}
