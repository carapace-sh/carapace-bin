package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var databasesCmd = &cobra.Command{
	Use:   "databases",
	Short: "Display commands that manage databases",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databasesCmd).Standalone()
	rootCmd.AddCommand(databasesCmd)
}
