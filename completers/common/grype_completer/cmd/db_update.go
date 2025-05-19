package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var db_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "download the latest vulnerability database",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(db_updateCmd).Standalone()

	dbCmd.AddCommand(db_updateCmd)
}
