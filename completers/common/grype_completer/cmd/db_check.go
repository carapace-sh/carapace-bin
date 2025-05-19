package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var db_checkCmd = &cobra.Command{
	Use:   "check",
	Short: "check to see if there is a database update available",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(db_checkCmd).Standalone()

	dbCmd.AddCommand(db_checkCmd)
}
