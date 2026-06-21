package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "balooctl6",
	Short: "Command-line interface for the Baloo file indexer",
	Long:  "https://community.kde.org/Baloo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().StringP("format", "f", "multiline", "Output format <multiline|json|simple>")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Displays help on commandline options")
	rootCmd.PersistentFlags().Bool("help-all", false, "Displays help, including generic Qt options")
	rootCmd.PersistentFlags().String("qmljsdebugger", "", "Activates the QML/JS debugger with a specified port")
	rootCmd.PersistentFlags().BoolP("version", "v", false, "Displays version information")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("multiline", "json", "simple"),
	})
}
