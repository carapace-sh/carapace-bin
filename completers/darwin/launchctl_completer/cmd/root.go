package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "launchctl",
	Short: "control service instances",
	Long:  "https://keith.github.io/xcode-manpages/launchctl.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("domain", "D", "", "Apply the command to the specified domain")
	rootCmd.Flags().BoolP("help", "h", false, "Display a brief usage message and exit")
	rootCmd.Flags().BoolP("quiet", "q", false, "Suppress informational output")
	rootCmd.Flags().BoolP("verbose", "v", false, "Increase the verbosity of informational output")
	rootCmd.Flags().BoolP("version", "V", false, "Print the version of launchctl and exit")
}
