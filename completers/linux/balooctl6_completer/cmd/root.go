package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "balooctl6",
	Short: "Control the Baloo file indexer",
	Long:  "https://community.kde.org/Baloo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.PersistentFlags().StringP("format", "f", "multiline", "Output format for file status")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Show help")
	rootCmd.PersistentFlags().Bool("help-all", false, "Show help including generic Qt options")
	rootCmd.PersistentFlags().String("qmljsdebugger", "", "Activate the QML/JS debugger")
	rootCmd.PersistentFlags().BoolP("version", "v", false, "Show version information")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("multiline", "json", "simple"),
	})
}
