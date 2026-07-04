package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appleRAIDCmd = &cobra.Command{
	Use:   "appleRAID",
	Short: "Apple RAID set operations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appleRAIDCmd).Standalone()
	rootCmd.AddCommand(appleRAIDCmd)
}
