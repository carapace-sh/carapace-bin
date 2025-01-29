package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var farmCmd = &cobra.Command{
	Use:   "farm",
	Short: "Farm out builds to remote machines",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(farmCmd).Standalone()

	rootCmd.AddCommand(farmCmd)
}
