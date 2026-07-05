package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var luckyCmd = &cobra.Command{
	Use:   "lucky",
	Short: "Install the first search result",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(luckyCmd).Standalone()
	rootCmd.AddCommand(luckyCmd)
}
