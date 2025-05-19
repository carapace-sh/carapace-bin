package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var untapCmd = &cobra.Command{
	Use:     "untap",
	Short:   "Remove a tapped formula repository",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(untapCmd).Standalone()

	untapCmd.Flags().Bool("debug", false, "Display any debugging information.")
	untapCmd.Flags().Bool("force", false, "Untap even if formulae or casks from this tap are currently installed.")
	untapCmd.Flags().Bool("help", false, "Show this message.")
	untapCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	untapCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(untapCmd)
}
