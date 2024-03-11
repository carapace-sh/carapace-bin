package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var duplicateCmd = &cobra.Command{
	Use:   "duplicate [OPTIONS] [REVISIONS]...",
	Short: "Create a new change with the same content as an existing one",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(duplicateCmd).Standalone()

	duplicateCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(duplicateCmd)

	carapace.Gen(duplicateCmd).PositionalAnyCompletion(
		jj.ActionRevs(jj.RevOption{}.Default()).FilterArgs(),
	)
}
