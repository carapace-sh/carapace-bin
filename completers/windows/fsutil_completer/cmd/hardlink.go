package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hardlinkCmd = &cobra.Command{
	Use:   "hardlink",
	Short: "hardlink management",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hardlinkCmd).Standalone()
	rootCmd.AddCommand(hardlinkCmd)
}
