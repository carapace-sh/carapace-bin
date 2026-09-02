package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/npm"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:     "update",
	Short:   "Update packages",
	Aliases: []string{"u", "udpate", "up", "upgrade"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updateCmd).Standalone()
	updateCmd.Flags().Bool("allow-directory", false, "allow installing dependencies from directories")
	updateCmd.Flags().Bool("allow-file", false, "allow installing dependencies from tarball files")
	updateCmd.Flags().Bool("allow-git", false, "allow fetching dependencies from git references")
	updateCmd.Flags().Bool("allow-remote", false, "allow fetching dependencies from urls")
	updateCmd.Flags().StringArray("allow-scripts", nil, "packages whose install-time lifecycle scripts are allowed to run")
	updateCmd.Flags().Bool("audit", false, "submit audit reports")
	updateCmd.Flags().String("before", "", "only install versions available on or before the given date")
	updateCmd.Flags().Bool("bin-links", false, "Create symlinks for package executables")
	updateCmd.Flags().Bool("dangerously-allow-all-scripts", false, "bypass the allowScripts policy entirely")
	updateCmd.Flags().Bool("dry-run", false, "Only report changes")
	updateCmd.Flags().Bool("foreground-scripts", false, "run build scripts in the foreground process")
	updateCmd.Flags().Bool("fund", false, "Display funding message")
	updateCmd.Flags().BoolP("global", "g", false, "operate in global mode")
	updateCmd.Flags().Bool("global-style", false, "Use global layout")
	updateCmd.Flags().Bool("ignore-scripts", false, "do not run scripts specified in package.json")
	updateCmd.Flags().String("include", "", "include dependency types")
	updateCmd.Flags().String("install-strategy", "", "strategy for installing packages in node_modules")
	updateCmd.Flags().Bool("legacy-bundling", false, "Use legacy bundling")
	updateCmd.Flags().String("min-release-age", "", "only install versions available more than the given number of days ago")
	updateCmd.Flags().StringArray("min-release-age-exclude", nil, "packages exempt from the min-release-age filter")
	updateCmd.Flags().StringArray("omit", nil, "omit dependency types")
	updateCmd.Flags().Bool("package-lock", false, "Only update package-lock.json")
	updateCmd.Flags().BoolP("save", "S", false, "Package will appear in your dependencies")
	updateCmd.Flags().Bool("strict-allow-scripts", false, "turn install-script policy from warning into error")
	updateCmd.Flags().Bool("strict-peer-deps", false, "Fail and abort for any conflicting `peerDependencies`")
	addWorkspaceFlags(updateCmd)

	rootCmd.AddCommand(updateCmd)

	carapace.Gen(updateCmd).PositionalAnyCompletion(
		npm.ActionModules(), // TODO support global
	)
}
