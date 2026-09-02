package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Bump a package version",
	Aliases: []string{"verison"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionCmd).Standalone()
	versionCmd.Flags().Bool("allow-same-version", false, "prevent error when same version as the current one is used")
	versionCmd.Flags().Bool("commit-hooks", false, "run git commit hooks")
	versionCmd.Flags().Bool("git-tag-version", false, "tag the commit")
	versionCmd.Flags().Bool("ignore-scripts", false, "do not run scripts specified in package.json")
	versionCmd.Flags().Bool("json", false, "output as json")
	versionCmd.Flags().Bool("no-commit-hooks", false, "skip git commit hooks")
	versionCmd.Flags().Bool("no-git-tag-version", false, "skip tagging the commit")
	versionCmd.Flags().Bool("no-save", false, "do not save updated version to package.json")
	versionCmd.Flags().Bool("no-workspaces-update", false, "do not update all workspaces")
	versionCmd.Flags().String("preid", "", "prerelease identifier")
	versionCmd.Flags().BoolP("save", "S", false, "save updated version to package.json")
	versionCmd.Flags().Bool("save-bundle", false, "save updated version to bundleDependencies")
	versionCmd.Flags().Bool("save-dev", false, "save updated version to devDependencies")
	versionCmd.Flags().Bool("save-optional", false, "save updated version to optionalDependencies")
	versionCmd.Flags().Bool("save-peer", false, "save updated version to peerDependencies")
	versionCmd.Flags().Bool("save-prod", false, "save updated version to dependencies")
	versionCmd.Flags().Bool("sign-git-tag", false, "sign with gpg signature")
	versionCmd.Flags().Bool("workspaces-update", false, "update all workspaces")
	addWorkspaceFlags(versionCmd)

	rootCmd.AddCommand(versionCmd)

	carapace.Gen(versionCmd).PositionalCompletion(
		carapace.ActionValues("major", "minor", "patch", "premajor", "preminor", "prepatch", "prerelease", "from-git"),
	)
}
