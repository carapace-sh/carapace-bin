package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dbSchemaCmd = &cobra.Command{
	Use:   "schema",
	Short: "Database schema tasks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dbSchemaCmd).Standalone()
	dbSchemaCmd.Flags().BoolP("help", "h", false, "Show help")
}
