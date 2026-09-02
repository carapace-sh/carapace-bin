package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/npm_completer/cmd/action"
	"github.com/carapace-sh/carapace/pkg/condition"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:     "install",
	Short:   "Install a package",
	Aliases: []string{"add", "i", "in", "ins", "inst", "insta", "instal", "isnt", "isnta", "isntal", "isntall"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()

	installCmd.Flags().String("allow-directory", "", "Limits installing dependencies from directories")
	installCmd.Flags().String("allow-file", "", "Limits installing dependencies from tarball files")
	installCmd.Flags().String("allow-git", "", "Limits fetching dependencies from git references")
	installCmd.Flags().String("allow-remote", "", "Limits fetching dependencies from urls")
	installCmd.Flags().StringArray("allow-scripts", nil, "Packages whose install-time lifecycle scripts are allowed to run")
	installCmd.Flags().Bool("audit", false, "Conduct security audit")
	installCmd.Flags().String("before", "", "Only install versions available on or before the given date")
	installCmd.Flags().Bool("bin-links", false, "Create symlinks for package executables")
	installCmd.Flags().String("cpu", "", "Override CPU architecture of native modules to install")
	installCmd.Flags().Bool("dangerously-allow-all-scripts", false, "Bypass the allowScripts policy entirely")
	installCmd.Flags().Bool("dry-run", false, "Only report changes")
	installCmd.Flags().Bool("foreground-scripts", false, "Run build scripts in the foreground process")
	installCmd.Flags().Bool("fund", false, "Display funding message")
	installCmd.Flags().BoolP("global", "g", false, "operate in global mode")
	installCmd.Flags().Bool("global-style", false, "Use global layout")
	installCmd.Flags().Bool("ignore-scripts", false, "Disable scripts")
	installCmd.Flags().String("include", "", "Include dependency types to install")
	installCmd.Flags().String("install-strategy", "", "Strategy for installing packages in node_modules")
	installCmd.Flags().Bool("legacy-bundling", false, "Use legacy bundling")
	installCmd.Flags().String("libc", "", "Override libc of native modules to install")
	installCmd.Flags().String("min-release-age", "", "Only install versions available more than the given number of days ago")
	installCmd.Flags().StringArray("min-release-age-exclude", nil, "Packages exempt from the min-release-age filter")
	installCmd.Flags().Bool("no-audit", false, "Skip audit")
	installCmd.Flags().Bool("no-bin-links", false, "Skip symlinks for package executables")
	installCmd.Flags().Bool("no-fund", false, "Skip funding message")
	installCmd.Flags().Bool("no-package-lock", false, "Only update package-lock.json")
	installCmd.Flags().Bool("no-save", false, "Prevents saving to `dependencies`")
	installCmd.Flags().String("omit", "", "Exclude package")
	installCmd.Flags().String("os", "", "Override OS of native modules to install")
	installCmd.Flags().Bool("package-lock", false, "Only update package-lock.json")
	installCmd.Flags().Bool("package-lock-only", false, "Only use the package-lock.json")
	installCmd.Flags().Bool("prefer-dedupe", false, "Prefer to deduplicate packages if possible")
	installCmd.Flags().BoolP("save", "S", false, "Package will appear in your `dependencies`")
	installCmd.Flags().Bool("save-bundle", false, "Package will appear in your `bundleDependencies`")
	installCmd.Flags().Bool("save-dev", false, "Package will appear in your `devDependencies`")
	installCmd.Flags().BoolP("save-exact", "E", false, "Save exact package version")
	installCmd.Flags().Bool("save-optional", false, "Package will appear in your `optionalDependencies`")
	installCmd.Flags().Bool("save-peer", false, "Package will appear in your `peerDependencies`")
	installCmd.Flags().Bool("save-prod", false, "Package will appear in your `dependencies`.")
	installCmd.Flags().Bool("strict-allow-scripts", false, "Turn install-script policy from warning into error")
	installCmd.Flags().Bool("strict-peer-deps", false, "Fail and abort for any conflicting `peerDependencies`")
	addWorkspaceFlags(installCmd)
	rootCmd.AddCommand(installCmd)

	carapace.Gen(installCmd).FlagCompletion(carapace.ActionMap{
		"allow-directory":  carapace.ActionValues("all", "none", "root"),
		"allow-file":       carapace.ActionValues("all", "none", "root"),
		"allow-git":        carapace.ActionValues("all", "none", "root"),
		"allow-remote":     carapace.ActionValues("all", "none", "root"),
		"include":          carapace.ActionValues("prod", "dev", "optional", "peer"),
		"install-strategy": carapace.ActionValues("hoisted", "nested", "shallow", "linked"),
		"omit":             carapace.ActionValues("dev", "optional", "peer"),
	})

	carapace.Gen(installCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionFiles(),
			action.ActionPackages(installCmd).UnlessF(condition.CompletingPath),
		).ToA(),
	)
}
