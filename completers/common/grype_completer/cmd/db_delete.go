package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var db_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete the vulnerability database",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(db_deleteCmd).Standalone()

	dbCmd.AddCommand(db_deleteCmd)
}
