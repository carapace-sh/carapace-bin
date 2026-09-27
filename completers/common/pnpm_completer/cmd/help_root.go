package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_rootCmd = &cobra.Command{
	Use:   "root",
	Short: "Print the effective `node_modules` directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_rootCmd).Standalone()

	helpCmd.AddCommand(help_rootCmd)
}
