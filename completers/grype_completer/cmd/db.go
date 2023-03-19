package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "vulnerability database operations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dbCmd).Standalone()

	rootCmd.AddCommand(dbCmd)
}
