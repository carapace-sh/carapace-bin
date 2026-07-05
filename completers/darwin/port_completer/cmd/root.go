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

	rootCmd.Flags().BoolP("auto-clean", "c", false, "Autoclean mode (clean after install)")
	rootCmd.Flags().BoolP("binary-only", "b", false, "Binary-only mode (install from archives)")
	rootCmd.Flags().BoolP("debug", "d", false, "Debug mode (implies -v)")
	rootCmd.Flags().BoolP("dry-run", "y", false, "Dry run (compute steps but do not execute)")
	rootCmd.Flags().BoolP("force", "f", false, "Force mode (ignore state file)")
	rootCmd.Flags().BoolP("keep", "k", false, "Keep mode (do not autoclean)")
	rootCmd.Flags().StringP("location", "D", "", "Change to directory before processing")
	rootCmd.Flags().BoolP("no-follow-dependencies", "n", false, "Do not follow dependencies")
	rootCmd.Flags().BoolP("noninteractive", "N", false, "Non-interactive mode")
	rootCmd.Flags().BoolP("proceed", "p", false, "Proceed despite errors")
	rootCmd.Flags().BoolP("quiet", "q", false, "Quiet mode (suppress informational messages)")
	rootCmd.Flags().BoolP("source-only", "s", false, "Source-only mode (build from source)")
	rootCmd.Flags().BoolP("trace", "t", false, "Enable trace mode debug facilities")
	rootCmd.Flags().BoolP("uninstall-inactive", "u", false, "Uninstall inactive ports when upgrading")
	rootCmd.Flags().BoolP("upgrade-dependents", "R", false, "Also upgrade dependents")
	rootCmd.Flags().BoolP("verbose", "v", false, "Verbose mode")
}
