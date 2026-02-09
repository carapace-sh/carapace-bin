package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resolve_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show the status of conflict resolution, listing remaining conflicted files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resolve_statusCmd).Standalone()

	resolve_statusCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	resolveCmd.AddCommand(resolve_statusCmd)
}
