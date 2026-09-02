package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var installCiTestCmd = &cobra.Command{
	Use:     "install-ci-test",
	Short:   "Install a project with a clean slate and run tests",
	Aliases: []string{"cit", "clean-install-test", "sit"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCiTestCmd).Standalone()
	installCiTestCmd.Flags().Bool("allow-directory", false, "allow installing dependencies from directories")
	installCiTestCmd.Flags().Bool("allow-file", false, "allow installing dependencies from tarball files")
	installCiTestCmd.Flags().Bool("allow-git", false, "allow fetching dependencies from git references")
	installCiTestCmd.Flags().Bool("allow-remote", false, "allow fetching dependencies from urls")
	installCiTestCmd.Flags().StringArray("allow-scripts", nil, "packages whose install-time lifecycle scripts are allowed to run")
	installCiTestCmd.Flags().Bool("audit", false, "conduct security audit")
	installCiTestCmd.Flags().Bool("bin-links", false, "create symlinks for package executables")
	installCiTestCmd.Flags().Bool("dangerously-allow-all-scripts", false, "bypass the allowScripts policy entirely")
	installCiTestCmd.Flags().Bool("dry-run", false, "only report changes")
	installCiTestCmd.Flags().Bool("foreground-scripts", false, "run build scripts in the foreground process")
	installCiTestCmd.Flags().Bool("fund", false, "display funding message")
	installCiTestCmd.Flags().BoolP("global", "g", false, "operate in global mode")
	installCiTestCmd.Flags().Bool("global-style", false, "use global layout")
	installCiTestCmd.Flags().Bool("ignore-scripts", false, "do not run scripts specified in package.json")
	installCiTestCmd.Flags().String("include", "", "include dependency types to install")
	installCiTestCmd.Flags().String("install-strategy", "", "strategy for installing packages in node_modules")
	installCiTestCmd.Flags().Bool("legacy-bundling", false, "use legacy bundling")
	installCiTestCmd.Flags().StringArray("omit", nil, "omit dependency types")
	installCiTestCmd.Flags().Bool("package-lock", false, "use package-lock.json")
	installCiTestCmd.Flags().Bool("strict-allow-scripts", false, "turn install-script policy from warning into error")
	installCiTestCmd.Flags().Bool("strict-peer-deps", false, "fail and abort for any conflicting peerDependencies")
	addWorkspaceFlags(installCiTestCmd)

	rootCmd.AddCommand(installCiTestCmd)

	carapace.Gen(installCiTestCmd).FlagCompletion(carapace.ActionMap{
		"include":          carapace.ActionValues("prod", "dev", "optional", "peer"),
		"install-strategy": carapace.ActionValues("hoisted", "nested", "shallow", "linked"),
		"omit":             carapace.ActionValues("dev", "optional", "peer"),
	})
}
