package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/npm_completer/cmd/action"
	"github.com/carapace-sh/carapace/pkg/condition"
	"github.com/spf13/cobra"
)

var installTestCmd = &cobra.Command{
	Use:     "install-test",
	Short:   "Install package(s) and run tests",
	Aliases: []string{"it"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installTestCmd).Standalone()

	installTestCmd.Flags().String("allow-directory", "", "Limits installing dependencies from directories")
	installTestCmd.Flags().String("allow-file", "", "Limits installing dependencies from tarball files")
	installTestCmd.Flags().String("allow-git", "", "Limits fetching dependencies from git references")
	installTestCmd.Flags().String("allow-remote", "", "Limits fetching dependencies from urls")
	installTestCmd.Flags().StringArray("allow-scripts", nil, "Packages whose install-time lifecycle scripts are allowed to run")
	installTestCmd.Flags().Bool("audit", false, "Conduct security audit")
	installTestCmd.Flags().String("before", "", "Only install packages available on or before the given date")
	installTestCmd.Flags().Bool("bin-links", false, "Create symlinks for package executables")
	installTestCmd.Flags().String("cpu", "", "Override CPU architecture of native modules to install")
	installTestCmd.Flags().Bool("dangerously-allow-all-scripts", false, "Bypass the allowScripts policy entirely")
	installTestCmd.Flags().Bool("dry-run", false, "Only report changes")
	installTestCmd.Flags().Bool("foreground-scripts", false, "Run build scripts in the foreground process")
	installTestCmd.Flags().Bool("fund", false, "Display funding message")
	installTestCmd.Flags().BoolP("global", "g", false, "operate in global mode")
	installTestCmd.Flags().Bool("global-style", false, "Use global layout")
	installTestCmd.Flags().Bool("ignore-scripts", false, "Disable scripts")
	installTestCmd.Flags().String("include", "", "Include dependency types to install")
	installTestCmd.Flags().String("install-strategy", "", "Strategy for installing packages in node_modules")
	installTestCmd.Flags().Bool("legacy-bundling", false, "Use legacy bundling")
	installTestCmd.Flags().String("libc", "", "Override libc of native modules to install")
	installTestCmd.Flags().String("min-release-age", "", "Only install packages available more than the given number of days ago")
	installTestCmd.Flags().StringArray("min-release-age-exclude", nil, "Packages exempt from the min-release-age filter")
	installTestCmd.Flags().Bool("no-audit", false, "Skip audit")
	installTestCmd.Flags().Bool("no-bin-links", false, "Skip symlinks for package executables")
	installTestCmd.Flags().Bool("no-fund", false, "Skip funding message")
	installTestCmd.Flags().Bool("no-package-lock", false, "Only update package-lock.json")
	installTestCmd.Flags().Bool("no-save", false, "Prevents saving to `dependencies`")
	installTestCmd.Flags().StringArray("omit", []string{""}, "Exclude package")
	installTestCmd.Flags().String("os", "", "Override OS of native modules to install")
	installTestCmd.Flags().Bool("package-lock", false, "Only update package-lock.json")
	installTestCmd.Flags().Bool("package-lock-only", false, "Only use the package-lock.json")
	installTestCmd.Flags().Bool("prefer-dedupe", false, "Prefer to deduplicate packages if possible")
	installTestCmd.Flags().BoolP("save", "S", false, "Package will appear in your `dependencies`")
	installTestCmd.Flags().Bool("save-bundle", false, "Package will appear in your `bundleDependencies`")
	installTestCmd.Flags().Bool("save-dev", false, "Package will appear in your `devDependencies`")
	installTestCmd.Flags().BoolP("save-exact", "E", false, "Save exact package version")
	installTestCmd.Flags().Bool("save-optional", false, "Package will appear in your `optionalDependencies`")
	installTestCmd.Flags().Bool("save-peer", false, "Package will appear in your `peerDependencies`")
	installTestCmd.Flags().Bool("save-prod", false, "Package will appear in your `dependencies`.")
	installTestCmd.Flags().Bool("strict-allow-scripts", false, "Turn install-script policy from warning into error")
	installTestCmd.Flags().Bool("strict-peer-deps", false, "Fail and abort for any conflicting `peerDependencies`")
	addWorkspaceFlags(installTestCmd)
	rootCmd.AddCommand(installTestCmd)

	carapace.Gen(installTestCmd).FlagCompletion(carapace.ActionMap{
		"allow-directory":  carapace.ActionValues("all", "none", "root"),
		"allow-file":       carapace.ActionValues("all", "none", "root"),
		"allow-git":        carapace.ActionValues("all", "none", "root"),
		"allow-remote":     carapace.ActionValues("all", "none", "root"),
		"include":          carapace.ActionValues("prod", "dev", "optional", "peer"),
		"install-strategy": carapace.ActionValues("hoisted", "nested", "shallow", "linked"),
		"omit":             carapace.ActionValues("dev", "optional", "peer"),
	})

	carapace.Gen(installTestCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionFiles(),
			action.ActionPackages(installTestCmd).UnlessF(condition.CompletingPath),
		).ToA(),
	)
}
