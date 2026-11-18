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

	casksCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	casksCmd.Flags().BoolP("help", "h", false, "Show this message.")
	casksCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	casksCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(casksCmd)
}
