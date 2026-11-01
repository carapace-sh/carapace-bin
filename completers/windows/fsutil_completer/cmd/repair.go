package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repairCmd = &cobra.Command{
	Use:   "repair",
	Short: "NTFS self-healing management",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repairCmd).Standalone()
	rootCmd.AddCommand(repairCmd)
}
