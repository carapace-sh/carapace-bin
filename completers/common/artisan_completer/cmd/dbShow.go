package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dbShowCmd = &cobra.Command{
	Use:   "db:show [--database [DATABASE]] [--json] [--counts] [--views] [--types]",
	Short: "Display information about the given database",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dbShowCmd).Standalone()

	dbShowCmd.Flags().Bool("counts", false, "Show the table row count <bg=red;options=bold> Note: This can be slow on large databases </>")
	dbShowCmd.Flags().String("database", "", "The database connection")
	dbShowCmd.Flags().Bool("json", false, "Output the database information as JSON")
	dbShowCmd.Flags().Bool("types", false, "Show the user defined types")
	dbShowCmd.Flags().Bool("views", false, "Show the database views <bg=red;options=bold> Note: This can be slow on large databases </>")
	rootCmd.AddCommand(dbShowCmd)
}
