package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var repo_license_viewCmd = &cobra.Command{
	Use:   "view {<license-key> | <spdx-id>}",
	Short: "View a specific repository license",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_license_viewCmd).Standalone()

	repo_license_viewCmd.Flags().BoolP("web", "w", false, "Open https://choosealicense.com/ in the browser")
	repo_licenseCmd.AddCommand(repo_license_viewCmd)

	carapace.Gen(repo_license_viewCmd).PositionalCompletion(
		gh.ActionLicenses(gh.HostOpts{}),
	)
}
