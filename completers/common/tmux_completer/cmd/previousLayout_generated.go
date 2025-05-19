package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var previousLayoutCmd = &cobra.Command{
	Use:   "previous-layout",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(previousLayoutCmd).Standalone()

	previousLayoutCmd.Flags().StringS("t", "t", "", "target-window")
	rootCmd.AddCommand(previousLayoutCmd)
}
