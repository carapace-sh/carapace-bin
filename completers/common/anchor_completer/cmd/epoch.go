package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var epochCmd = &cobra.Command{
	Use:   "epoch",
	Short: "Get current epoch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(epochCmd).Standalone()

	epochCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(epochCmd)
}
