package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "port",
	Short: "MacPorts package manager",
	Long:  "https://guide.macports.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("quiet", "q", false, "Quiet mode")
	rootCmd.Flags().BoolP("verbose", "v", false, "Verbose mode")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"install", "Install a port",
			"uninstall", "Uninstall a port",
			"upgrade", "Upgrade a port",
			"activate", "Activate a port",
			"deactivate", "Deactivate a port",
			"search", "Search for ports",
			"list", "List available ports",
			"info", "Show information about a port",
			"deps", "Show dependencies of a port",
			"variants", "Show variants of a port",
			"contents", "Show contents of an installed port",
			"location", "Show location of an installed port",
			"installed", "List installed ports",
			"outdated", "List outdated ports",
			"sync", "Sync ports tree",
			"selfupdate", "Update MacPorts and ports tree",
			"clean", "Clean build files",
			"fetch", "Fetch distfiles",
			"build", "Build a port",
			"test", "Test a port",
			"destroot", "Destroot a port",
			"archive", "Archive a port",
			"unarchive", "Unarchive a port",
		),
	)
}
