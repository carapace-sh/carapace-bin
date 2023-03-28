package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/apt"
	"github.com/spf13/cobra"
)

var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify the integrity of package(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(verifyCmd).Standalone()

	carapace.Gen(verifyCmd).PositionalAnyCompletion(
		apt.ActionPackages(),
	)
}
