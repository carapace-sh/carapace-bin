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

	formulaeCmd.Flags().Bool("debug", false, "Display any debugging information.")
	formulaeCmd.Flags().Bool("help", false, "Show this message.")
	formulaeCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	formulaeCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(formulaeCmd)
}
