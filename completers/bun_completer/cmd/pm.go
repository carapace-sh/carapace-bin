package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var pmCmd = &cobra.Command{
	Use:   "pm",
	Short: "More commands for managing packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pmCmd).Standalone()

	pmCmd.Flags().String("backend", "", "Platform-specific optimizations for installing dependencies")
	pmCmd.Flags().String("cache-dir", "", "Store & load cached data from a specific directory path")
	pmCmd.Flags().StringP("config", "c", "", "Load config")
	pmCmd.Flags().String("cwd", "", "Set a specific cwd")
	pmCmd.Flags().Bool("dry-run", false, "Don't install anything")
	pmCmd.Flags().BoolP("force", "f", false, "Always request the latest versions from the registry & reinstall all dependencies")
	pmCmd.Flags().BoolP("global", "g", false, "Install globally")
	pmCmd.Flags().Bool("help", false, "Print this help menu")
	pmCmd.Flags().Bool("ignore-scripts", false, "Skip lifecycle scripts in the project's package.json")
	pmCmd.Flags().String("link-native-bins", "", "Link \"bin\" from a matching platform-specific \"optionalDependencies\" instead")
	pmCmd.Flags().String("lockfile", "", "Store & load a lockfile at a specific filepath")
	pmCmd.Flags().Bool("no-cache", false, "Ignore manifest cache entirely")
	pmCmd.Flags().Bool("no-progress", false, "Disable the progress bar")
	pmCmd.Flags().Bool("no-save", false, "Don't save a lockfile")
	pmCmd.Flags().Bool("no-summary", false, "Don't print a summary")
	pmCmd.Flags().Bool("no-verify", false, "Skip verifying integrity of newly downloaded packages")
	pmCmd.Flags().BoolP("production", "p", false, "Don't install devDependencies")
	pmCmd.Flags().Bool("silent", false, "Don't log anything")
	pmCmd.Flags().Bool("verbose", false, "Excessively verbose logging")
	pmCmd.Flags().BoolP("yarn", "y", false, "Write a yarn.lock file")
	rootCmd.AddCommand(pmCmd)

	carapace.Gen(pmCmd).FlagCompletion(carapace.ActionMap{
		"backend":   carapace.ActionValues("hadlink", "symlink", "copyfile"),
		"cache-dir": carapace.ActionDirectories(),
		"config":    carapace.ActionFiles(),
		"cwd":       carapace.ActionDirectories(),
		"lockfile":  carapace.ActionFiles(),
	})

	carapace.Gen(pmCmd).PreInvoke(func(cmd *cobra.Command, flag *pflag.Flag, action carapace.Action) carapace.Action {
		return action.Chdir(pmCmd.Flag("cwd").Value.String())
	})
}
