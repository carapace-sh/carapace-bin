package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var telescopePruneCmd = &cobra.Command{
	Use:   "telescope:prune [--hours [HOURS]] [--keep-exceptions]",
	Short: "Prune stale entries from the Telescope database",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(telescopePruneCmd).Standalone()

	telescopePruneCmd.Flags().String("hours", "", "The number of hours to retain Telescope data")
	telescopePruneCmd.Flags().Bool("keep-exceptions", false, "Retain exception data")
	rootCmd.AddCommand(telescopePruneCmd)
}
