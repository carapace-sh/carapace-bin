package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var casksCmd = &cobra.Command{
	Use:     "casks",
	Short:   "List all locally installable casks including short names",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(casksCmd).Standalone()

	casksCmd.Flags().Bool("debug", false, "Display any debugging information.")
	casksCmd.Flags().Bool("help", false, "Show this message.")
	casksCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	casksCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(casksCmd)
}
