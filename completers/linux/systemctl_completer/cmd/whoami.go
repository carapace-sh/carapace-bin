package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/ps"
	"github.com/spf13/cobra"
)

var whoamiCmd = &cobra.Command{
	Use:     "whoami",
	Short:   "Return unit caller or specified PIDs are part of",
	GroupID: "unit",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(whoamiCmd).Standalone()

	rootCmd.AddCommand(whoamiCmd)

	carapace.Gen(whoamiCmd).PositionalAnyCompletion(
		ps.ActionProcessIds(),
	)
}
