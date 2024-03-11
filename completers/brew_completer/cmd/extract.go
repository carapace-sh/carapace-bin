package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var extractCmd = &cobra.Command{
	Use:     "extract",
	Short:   "Look through repository history to find the most recent version of <formula> and create a copy in <tap>",
	GroupID: "developer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(extractCmd).Standalone()

	extractCmd.Flags().Bool("debug", false, "Display any debugging information.")
	extractCmd.Flags().Bool("force", false, "Overwrite the destination formula if it already exists.")
	extractCmd.Flags().Bool("help", false, "Show this message.")
	extractCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	extractCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	extractCmd.Flags().Bool("version", false, "Extract the specified <version> of <formula> instead of the most recent.")
	rootCmd.AddCommand(extractCmd)
}
