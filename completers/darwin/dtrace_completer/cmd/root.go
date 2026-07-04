package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dtrace",
	Short: "Dynamic tracing framework",
	Long:  "https://keith.github.io/xcode-manpages/dtrace.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("h", "h", false, "Display help")
	rootCmd.Flags().BoolS("q", "q", false, "Quiet mode")
	rootCmd.Flags().BoolS("v", "v", false, "Verbose mode")

	rootCmd.Flags().StringS("c", "c", "", "Run command")
	rootCmd.Flags().StringS("o", "o", "", "Output file")
	rootCmd.Flags().StringS("p", "p", "", "Process ID")
	rootCmd.Flags().StringS("s", "s", "", "Script file")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionValues())
}
