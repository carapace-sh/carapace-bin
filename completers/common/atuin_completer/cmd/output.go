package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var outputCmd = &cobra.Command{
	Use:   "output",
	Short: "Work with captured command output",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outputCmd).Standalone()

	outputCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(outputCmd)
}
