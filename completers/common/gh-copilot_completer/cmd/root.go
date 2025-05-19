package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gh-copilot",
	Short: "Your AI command line copilot",
	Long:  "https://github.com/github/gh-copilot",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "help for copilot")
	rootCmd.Flags().BoolP("version", "v", false, "version for copilot")
}
