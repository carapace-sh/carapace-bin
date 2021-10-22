package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var databases_poolCmd = &cobra.Command{
	Use:   "pool",
	Short: "Display commands for managing connection pools",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databases_poolCmd).Standalone()
	databasesCmd.AddCommand(databases_poolCmd)
}
