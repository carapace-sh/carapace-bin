package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var epochInfoCmd = &cobra.Command{
	Use:   "epoch-info",
	Short: "Get information about the current epoch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(epochInfoCmd).Standalone()

	epochInfoCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(epochInfoCmd)
}
