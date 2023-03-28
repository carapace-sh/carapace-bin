package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/apt"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Display package status details",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(statusCmd).Standalone()

	carapace.Gen(statusCmd).PositionalAnyCompletion(
		apt.ActionPackages(),
	)
}
