package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Runs the deploy migration script",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrateCmd).Standalone()

	migrateCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(migrateCmd)
}
