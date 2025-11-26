package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show high-level repo status [default alias: st]",
	Aliases: []string{"st"},
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(statusCmd).Standalone()

	statusCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")

	carapace.Gen(statusCmd).PositionalAnyCompletion(carapace.ActionFiles())

	rootCmd.AddCommand(statusCmd)
}
