package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deltaCmd = &cobra.Command{
	Use:     "delta",
	Short:   "Create a delta package",
	Aliases: []string{"dt"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deltaCmd).Standalone()

	deltaCmd.Flags().StringP("output-dir", "O", "", "Output directory for produced packages")

	rootCmd.AddCommand(deltaCmd)

	carapace.Gen(deltaCmd).PositionalAnyCompletion(
		carapace.ActionFiles(".eopkg"),
	)
}
