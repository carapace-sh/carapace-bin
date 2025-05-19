package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var horizonForgetCmd = &cobra.Command{
	Use:   "horizon:forget [--all] [--] [<id>]",
	Short: "Delete a failed queue job",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(horizonForgetCmd).Standalone()

	horizonForgetCmd.Flags().Bool("all", false, "Delete all failed jobs")
	rootCmd.AddCommand(horizonForgetCmd)
}
