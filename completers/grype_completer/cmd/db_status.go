package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var db_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "display database status",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(db_statusCmd).Standalone()

	dbCmd.AddCommand(db_statusCmd)
}
