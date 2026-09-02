package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ciCmd = &cobra.Command{
	Use:     "ci",
	Short:   "Install a project with a clean slate",
	Aliases: []string{"clean-install", "ic", "install-clean", "isntall-clean"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ciCmd).Standalone()
	ciCmd.Flags().Bool("allow-directory", false, "allow installing dependencies from directories")
	ciCmd.Flags().Bool("allow-file", false, "allow installing dependencies from tarball files")
	ciCmd.Flags().Bool("allow-git", false, "allow fetching dependencies from git references")
	ciCmd.Flags().Bool("allow-remote", false, "allow fetching dependencies from urls")
	ciCmd.Flags().StringArray("allow-scripts", nil, "packages whose install-time lifecycle scripts are allowed to run")
	ciCmd.Flags().Bool("audit", false, "conduct security audit")
	ciCmd.Flags().Bool("bin-links", false, "create symlinks for package executables")
	ciCmd.Flags().Bool("dangerously-allow-all-scripts", false, "bypass the allowScripts policy entirely")
	ciCmd.Flags().Bool("dry-run", false, "only report changes")
	ciCmd.Flags().Bool("foreground-scripts", false, "run build scripts in the foreground process")
	ciCmd.Flags().Bool("fund", false, "display funding message")
	ciCmd.Flags().BoolP("global", "g", false, "operate in global mode")
	ciCmd.Flags().Bool("global-style", false, "use global layout")
	ciCmd.Flags().Bool("ignore-scripts", false, "do not run scripts specified in package.json")
	ciCmd.Flags().String("include", "", "include dependency types to install")
	ciCmd.Flags().String("install-strategy", "", "strategy for installing packages in node_modules")
	ciCmd.Flags().Bool("legacy-bundling", false, "use legacy bundling")
	ciCmd.Flags().StringArray("omit", nil, "omit dependency types")
	ciCmd.Flags().Bool("package-lock", false, "use package-lock.json")
	ciCmd.Flags().Bool("strict-allow-scripts", false, "turn install-script policy from warning into error")
	ciCmd.Flags().Bool("strict-peer-deps", false, "fail and abort for any conflicting peerDependencies")
	addWorkspaceFlags(ciCmd)

	rootCmd.AddCommand(ciCmd)

	carapace.Gen(ciCmd).FlagCompletion(carapace.ActionMap{
		"include":          carapace.ActionValues("prod", "dev", "optional", "peer"),
		"install-strategy": carapace.ActionValues("hoisted", "nested", "shallow", "linked"),
		"omit":             carapace.ActionValues("dev", "optional", "peer"),
	})
}
