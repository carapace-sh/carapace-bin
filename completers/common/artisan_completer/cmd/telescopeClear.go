package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var telescopeClearCmd = &cobra.Command{
	Use:   "telescope:clear",
	Short: "Delete all Telescope data from storage",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(telescopeClearCmd).Standalone()

	rootCmd.AddCommand(telescopeClearCmd)
}
