package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinCmd = &cobra.Command{
	Use:     "pin",
	Short:   "Pin the specified <formula>, preventing them from being upgraded when issuing the `brew upgrade` <formula> command",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinCmd).Standalone()

	pinCmd.Flags().Bool("debug", false, "Display any debugging information.")
	pinCmd.Flags().Bool("help", false, "Show this message.")
	pinCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	pinCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(pinCmd)
}
