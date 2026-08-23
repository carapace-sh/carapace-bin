package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var uniquesizeCmd = &cobra.Command{
	Use:   "uniquesize",
	Short: "analyze a path's unique size in backups",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(uniquesizeCmd).Standalone()
	rootCmd.AddCommand(uniquesizeCmd)

	carapace.Gen(uniquesizeCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}