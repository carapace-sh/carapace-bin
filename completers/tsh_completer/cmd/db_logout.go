package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var db_logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Remove database credentials.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(db_logoutCmd).Standalone()

	dbCmd.AddCommand(db_logoutCmd)
}
