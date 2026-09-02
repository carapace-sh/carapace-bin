package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dedupeCmd = &cobra.Command{
	Use:     "dedupe",
	Short:   "Reduce duplication in the package tree",
	Aliases: []string{"ddp"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dedupeCmd).Standalone()

	dedupeCmd.Flags().Bool("allow-directory", false, "allow installing dependencies from directories")
	dedupeCmd.Flags().Bool("allow-file", false, "allow installing dependencies from tarball files")
	dedupeCmd.Flags().Bool("allow-git", false, "allow fetching dependencies from git references")
	dedupeCmd.Flags().Bool("allow-remote", false, "allow fetching dependencies from urls")
	dedupeCmd.Flags().Bool("audit", false, "submit audit reports")
	dedupeCmd.Flags().Bool("bin-links", false, "create symlinks")
	dedupeCmd.Flags().Bool("dry-run", false, "only report changes")
	dedupeCmd.Flags().Bool("fund", false, "show funding message")
	dedupeCmd.Flags().Bool("global-style", false, "use global layout")
	dedupeCmd.Flags().Bool("ignore-scripts", false, "do not run scripts specified in package.json")
	dedupeCmd.Flags().String("include", "", "include dependency types")
	dedupeCmd.Flags().String("install-strategy", "", "strategy for installing packages in node_modules")
	dedupeCmd.Flags().Bool("legacy-bundling", false, "support older npm version")
	dedupeCmd.Flags().Bool("no-audit", false, "skip audit")
	dedupeCmd.Flags().Bool("no-bin-links", false, "skip symlinks")
	dedupeCmd.Flags().Bool("no-fund", false, "skip funding message")
	dedupeCmd.Flags().Bool("no-package-lock", false, "ignore package-lock.json")
	dedupeCmd.Flags().String("omit", "", "omit dependency types")
	dedupeCmd.Flags().Bool("package-lock", false, "ignore package-lock.json")
	dedupeCmd.Flags().Bool("strict-peer-deps", false, "any conflicting `peerDependencies` will be treated as install failure")
	addWorkspaceFlags(dedupeCmd)
	rootCmd.AddCommand(dedupeCmd)

	carapace.Gen(dedupeCmd).FlagCompletion(carapace.ActionMap{
		"allow-directory":  carapace.ActionValues("all", "none", "root"),
		"allow-file":       carapace.ActionValues("all", "none", "root"),
		"allow-git":        carapace.ActionValues("all", "none", "root"),
		"allow-remote":     carapace.ActionValues("all", "none", "root"),
		"include":          carapace.ActionValues("prod", "dev", "optional", "peer"),
		"install-strategy": carapace.ActionValues("hoisted", "nested", "shallow", "linked"),
		"omit":             carapace.ActionValues("dev", "optional", "peer"),
	})
}
