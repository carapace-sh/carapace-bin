package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var CheckCmd = &cobra.Command{
	Use:   "check",
	Short: "Check",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(CheckCmd).Standalone()
	rootCmd.AddCommand(CheckCmd)
}
