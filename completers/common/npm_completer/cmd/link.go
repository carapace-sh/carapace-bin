package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/npm_completer/cmd/action"
	"github.com/carapace-sh/carapace/pkg/condition"
	"github.com/spf13/cobra"
)

var linkCmd = &cobra.Command{
	Use:     "link",
	Short:   "Symlink a package folder",
	Aliases: []string{"ln"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(linkCmd).Standalone()

	linkCmd.Flags().Bool("allow-directory", false, "allow installing dependencies from directories")
	linkCmd.Flags().Bool("allow-file", false, "allow installing dependencies from tarball files")
	linkCmd.Flags().Bool("allow-git", false, "allow fetching dependencies from git references")
	linkCmd.Flags().Bool("allow-remote", false, "allow fetching dependencies from urls")
	linkCmd.Flags().Bool("audit", false, "Conduct security audit")
	linkCmd.Flags().Bool("bin-links", false, "Create symlinks for package executables")
	linkCmd.Flags().Bool("dry-run", false, "Only report changes")
	linkCmd.Flags().Bool("fund", false, "Display funding message")
	linkCmd.Flags().BoolP("global", "g", false, "operate in global mode")
	linkCmd.Flags().Bool("global-style", false, "Use global layout")
	linkCmd.Flags().Bool("ignore-scripts", false, "Disable scripts")
	linkCmd.Flags().String("include", "", "include dependency types")
	linkCmd.Flags().String("install-strategy", "", "strategy for installing packages in node_modules")
	linkCmd.Flags().Bool("legacy-bundling", false, "Use legacy bundling")
	linkCmd.Flags().Bool("no-audit", false, "Skip audit")
	linkCmd.Flags().Bool("no-bin-links", false, "Skip symlinks for package executables")
	linkCmd.Flags().Bool("no-fund", false, "Skip funding message")
	linkCmd.Flags().Bool("no-package-lock", false, "Only update package-lock.json")
	linkCmd.Flags().Bool("no-save", false, "Prevents saving to `dependencies`")
	linkCmd.Flags().StringArray("omit", []string{""}, "Exclude package")
	linkCmd.Flags().Bool("package-lock", false, "Only update package-lock.json")
	linkCmd.Flags().BoolP("save", "S", false, "Package will appear in your `dependencies`")
	linkCmd.Flags().Bool("save-bundle", false, "Package will appear in your `bundleDependencies`")
	linkCmd.Flags().Bool("save-dev", false, "Package will appear in your `devDependencies`")
	linkCmd.Flags().BoolP("save-exact", "E", false, "Save exact package version")
	linkCmd.Flags().Bool("save-optional", false, "Package will appear in your `optionalDependencies`")
	linkCmd.Flags().Bool("save-peer", false, "Package will appear in your `peerDependencies`")
	linkCmd.Flags().Bool("save-prod", false, "Package will appear in your `dependencies`.")
	linkCmd.Flags().Bool("strict-peer-deps", false, "Fail and abort for any conflicting `peerDependencies`")
	addWorkspaceFlags(linkCmd)
	rootCmd.AddCommand(linkCmd)

	carapace.Gen(linkCmd).FlagCompletion(carapace.ActionMap{
		"allow-directory":  carapace.ActionValues("all", "none", "root"),
		"allow-file":       carapace.ActionValues("all", "none", "root"),
		"allow-git":        carapace.ActionValues("all", "none", "root"),
		"allow-remote":     carapace.ActionValues("all", "none", "root"),
		"include":          carapace.ActionValues("prod", "dev", "optional", "peer"),
		"install-strategy": carapace.ActionValues("hoisted", "nested", "shallow", "linked"),
		"omit":             carapace.ActionValues("dev", "optional", "peer"),
	})

	carapace.Gen(linkCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionFiles(),
			action.ActionPackages(linkCmd).UnlessF(condition.CompletingPath),
		).ToA(),
	)
}
