package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "terminal-browser",
	Short: "A browser inside your terminal",
	Long:  "https://terminal-browser.com/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "show help")
	rootCmd.Flags().Bool("version", false, "print the installed version")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
