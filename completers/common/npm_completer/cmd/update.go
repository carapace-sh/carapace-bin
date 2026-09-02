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
	updateCmd.Flags().StringArray("allow-scripts", nil, "packages whose install-time lifecycle scripts are allowed to run")
	updateCmd.Flags().String("before", "", "only install versions available on or before the given date")
	updateCmd.Flags().Bool("dangerously-allow-all-scripts", false, "bypass the allowScripts policy entirely")
	updateCmd.Flags().Bool("dry-run", false, "Only report changes")
	updateCmd.Flags().Bool("foreground-scripts", false, "run build scripts in the foreground process")
	updateCmd.Flags().BoolP("global", "g", false, "operate in global mode")
	updateCmd.Flags().Bool("global-style", false, "Use global layout")
	updateCmd.Flags().Bool("ignore-scripts", false, "do not run scripts specified in package.json")
	updateCmd.Flags().String("include", "", "include dependency types")
	updateCmd.Flags().String("install-strategy", "", "strategy for installing packages in node_modules")
	updateCmd.Flags().Bool("legacy-bundling", false, "Use legacy bundling")
	updateCmd.Flags().String("min-release-age", "", "only install versions available more than the given number of days ago")
	updateCmd.Flags().StringArray("min-release-age-exclude", nil, "packages exempt from the min-release-age filter")
	updateCmd.Flags().Bool("audit", false, "submit audit reports")
	updateCmd.Flags().Bool("bin-links", false, "Create symlinks for package executables")
	updateCmd.Flags().Bool("fund", false, "Display funding message")
	updateCmd.Flags().Bool("no-audit", false, "skip audit")
	updateCmd.Flags().Bool("no-bin-links", false, "Skip symlinks for package executables")
	updateCmd.Flags().Bool("no-fund", false, "Skip funding message")
	updateCmd.Flags().Bool("no-package-lock", false, "Only update package-lock.json")
	updateCmd.Flags().Bool("no-save", false, "Prevents saving to `dependencies`")
	updateCmd.Flags().Bool("package-lock", false, "Only update package-lock.json")
	updateCmd.Flags().StringArray("omit", nil, "omit dependency types")
	updateCmd.Flags().BoolP("save", "S", false, "Package will appear in your dependencies")
	updateCmd.Flags().Bool("save-bundle", false, "Package will appear in your `bundleDependencies`")
	updateCmd.Flags().Bool("save-dev", false, "Package will appear in your `devDependencies`")
	updateCmd.Flags().Bool("save-optional", false, "Package will appear in your `optionalDependencies`")
	updateCmd.Flags().Bool("save-peer", false, "Package will appear in your `peerDependencies`")
	updateCmd.Flags().Bool("save-prod", false, "Package will appear in your `dependencies`.")
	updateCmd.Flags().Bool("strict-allow-scripts", false, "turn install-script policy from warning into error")
	updateCmd.Flags().Bool("strict-peer-deps", false, "Fail and abort for any conflicting `peerDependencies`")
	addWorkspaceFlags(updateCmd)

	rootCmd.AddCommand(updateCmd)

	carapace.Gen(updateCmd).FlagCompletion(carapace.ActionMap{
		"include":          carapace.ActionValues("prod", "dev", "optional", "peer"),
		"install-strategy": carapace.ActionValues("hoisted", "nested", "shallow", "linked"),
		"omit":             carapace.ActionValues("dev", "optional", "peer"),
	})

	carapace.Gen(updateCmd).PositionalAnyCompletion(
		npm.ActionModules(), // TODO support global
	)
}
