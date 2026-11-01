package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-jjlex/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show high-level repo status [default alias: st]",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(statusCmd).Standalone()

	statusCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")

	carapace.Gen(statusCmd).PositionalAnyCompletion(
		jj.ActionRevFiles("@"),
	)

	rootCmd.AddCommand(statusCmd)
}
