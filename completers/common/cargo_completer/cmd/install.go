package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/cargo_completer/cmd/action"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:     "install",
	Short:   "Install a Rust binary",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groupFor("install"),
}

func init() {
	carapace.Gen(installCmd).Standalone()

	installCmd.Flags().Bool("all-features", false, "Activate all available features")
	installCmd.Flags().StringSlice("bin", nil, "Install only the specified binary")
	installCmd.Flags().Bool("bins", false, "Install all binaries")
	installCmd.Flags().String("branch", "", "Branch to use when installing from git")
	installCmd.Flags().Bool("debug", false, "Build in debug mode (with the 'dev' profile) instead of release mode")
	installCmd.Flags().StringSlice("example", nil, "Install only the specified example")
	installCmd.Flags().Bool("examples", false, "Install all examples")
	installCmd.Flags().StringSliceP("features", "F", nil, "Space or comma separated list of features to activate")
	installCmd.Flags().BoolP("force", "f", false, "Force overwriting existing crates or binaries")
	installCmd.Flags().String("git", "", "Git URL to install the specified crate from")
	installCmd.Flags().BoolP("help", "h", false, "Print help")
	installCmd.Flags().Bool("ignore-rust-version", false, "Ignore `rust-version` specification in packages")
	installCmd.Flags().String("index", "", "Registry index to install from")
	installCmd.Flags().StringP("jobs", "j", "", "Number of parallel jobs, defaults to # of CPUs.")
	installCmd.Flags().Bool("keep-going", false, "Do not abort the build as soon as there is an error")
	installCmd.Flags().Bool("list", false, "List all installed packages and their versions")
	installCmd.Flags().StringSlice("message-format", nil, "Error format")
	installCmd.Flags().Bool("no-default-features", false, "Do not activate the `default` feature")
	installCmd.Flags().Bool("no-track", false, "Do not save tracking information")
	installCmd.Flags().String("path", "", "Filesystem path to local crate to install from")
	installCmd.Flags().String("profile", "", "Install artifacts with the specified profile")
	installCmd.Flags().String("registry", "", "Registry to use")
	installCmd.Flags().String("rev", "", "Specific commit to use when installing from git")
	installCmd.Flags().String("root", "", "Directory to install packages into")
	installCmd.Flags().String("tag", "", "Tag to use when installing from git")
	installCmd.Flags().StringSlice("target", nil, "Build for the target triple")
	installCmd.Flags().String("target-dir", "", "Directory for all generated artifacts")
	installCmd.Flags().String("timings", "", "Timing output formats (unstable) (comma separated): html, json")
	installCmd.Flags().String("version", "", "Specify a version to install")
	installCmd.Flag("timings").NoOptDefVal = " "
	rootCmd.AddCommand(installCmd)

	// TODO missing flag completion
	carapace.Gen(installCmd).FlagCompletion(carapace.ActionMap{
		"bin":            action.ActionTargets(installCmd, action.TargetOpts{Bin: true}),
		"example":        action.ActionTargets(installCmd, action.TargetOpts{Example: true}),
		"features":       action.ActionFeatures(installCmd).UniqueList(","),
		"message-format": action.ActionMessageFormats(),
		"path":           carapace.ActionFiles(),
		"profile":        action.ActionProfiles(installCmd),
		"registry":       action.ActionRegistries(),
		"root":           carapace.ActionDirectories(),
		"target-dir":     carapace.ActionDirectories(),
	})
}
