package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apfsCmd = &cobra.Command{
	Use:   "apfs",
	Short: "APFS container and volume operations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apfsCmd).Standalone()
	rootCmd.AddCommand(apfsCmd)
}
