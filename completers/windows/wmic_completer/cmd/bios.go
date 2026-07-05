package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var biosCmd = &cobra.Command{
	Use:   "bios",
	Short: "BIOS management",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(biosCmd).Standalone()
	rootCmd.AddCommand(biosCmd)
}
