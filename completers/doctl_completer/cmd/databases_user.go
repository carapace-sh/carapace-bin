package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var databases_userCmd = &cobra.Command{
	Use:   "user",
	Short: "Display commands for managing database users",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databases_userCmd).Standalone()
	databasesCmd.AddCommand(databases_userCmd)
}
