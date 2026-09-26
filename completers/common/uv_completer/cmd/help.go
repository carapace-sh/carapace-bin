package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Display documentation for a command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(helpCmd).Standalone()

	helpCmd.Flags().Bool("no-pager", false, "Disable pager when printing help")
	rootCmd.AddCommand(helpCmd)
}
