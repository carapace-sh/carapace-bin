package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var formulaeCmd = &cobra.Command{
	Use:     "formulae",
	Short:   "List all locally installable formulae including short names",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(formulaeCmd).Standalone()

	formulaeCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	formulaeCmd.Flags().BoolP("help", "h", false, "Show this message.")
	formulaeCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	formulaeCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(formulaeCmd)
}
