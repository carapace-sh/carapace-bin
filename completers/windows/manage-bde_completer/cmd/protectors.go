package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var protectorsCmd = &cobra.Command{
	Use:   "protectors",
	Short: "manage protection methods",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(protectorsCmd).Standalone()
	rootCmd.AddCommand(protectorsCmd)
}
