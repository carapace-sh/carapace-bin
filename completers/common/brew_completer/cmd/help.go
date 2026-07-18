package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var helpCmd = &cobra.Command{
	Use:     "help",
	Short:   "Output the usage instructions for `brew` <command>",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(helpCmd).Standalone()

	helpCmd.Flags().Bool("debug", false, "Display any debugging information.")
	helpCmd.Flags().Bool("help", false, "Show this message.")
	helpCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	helpCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(helpCmd)
}
