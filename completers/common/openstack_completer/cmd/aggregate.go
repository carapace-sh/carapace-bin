package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var aggregateCmd = &cobra.Command{
	Use:   "aggregate",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(aggregateCmd).Standalone()

	rootCmd.AddCommand(aggregateCmd)
}
