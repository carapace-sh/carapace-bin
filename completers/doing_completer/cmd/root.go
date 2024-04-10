package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "doing [OPTIONS] COMMAND [ARGS]...",
	Short: "CLI for repository/issue workflow on Azure Devops",
	Long:  "https://github.com/ing-bank/doing-cli",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("help", false, "Show this message and exit.")
	rootCmd.Flags().Bool("version", false, "Show the version and exit.")
}
