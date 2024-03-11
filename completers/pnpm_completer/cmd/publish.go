package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pnpm"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publishes a package to the npm registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(publishCmd).Standalone()

	publishCmd.Flags().String("access", "", "Published as public or restricted")
	publishCmd.Flags().String("changed-files-ignore-pattern", "", "Defines files to ignore when filtering for changed projects")
	publishCmd.Flags().Bool("dry-run", false, "Does everything a publish would do except actually publishing")
	publishCmd.Flags().String("filter", "", "set filter")
	publishCmd.Flags().String("filter-prod", "", "Restricts the scope to package names matching the given pattern")
	publishCmd.Flags().Bool("force", false, "Publish packages even if their current version is already in the registry")
	publishCmd.Flags().Bool("ignore-scripts", false, "Ignores any publish related lifecycle scripts")
	publishCmd.Flags().Bool("json", false, "Show information in JSON format")
	publishCmd.Flags().Bool("no-git-checks", false, "Don't check if current branch is your publish branch, clean, and up to date")
	publishCmd.Flags().Bool("otp", false, "When publishing packages that require two-factor authentication")
	publishCmd.Flags().Bool("publish-branch", false, "Sets branch name to publish")
	publishCmd.Flags().BoolP("recursive", "r", false, "Publish all packages from the workspace")
	publishCmd.Flags().Bool("report-summary", false, "Save the list of the newly published")
	publishCmd.Flags().String("tag", "", "Registers the published package with the given tag")
	publishCmd.Flags().String("test-pattern", "", "Defines files related to tests ")
	rootCmd.AddCommand(publishCmd)

	carapace.Gen(publishCmd).FlagCompletion(carapace.ActionMap{
		"access":      carapace.ActionValues("public", "restricted").StyleF(style.ForKeyword),
		"filter":      pnpm.ActionFilter(),
		"filter-prod": pnpm.ActionFilter(),
	})

	carapace.Gen(publishCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
