package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var unpinCmd = &cobra.Command{
	Use:     "unpin",
	Short:   "Unpin <formula>, allowing them to be upgraded by `brew upgrade` <formula>",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unpinCmd).Standalone()

	unpinCmd.Flags().Bool("debug", false, "Display any debugging information.")
	unpinCmd.Flags().Bool("help", false, "Show this message.")
	unpinCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	unpinCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(unpinCmd)
}
