package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate *.templ files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrateCmd).Standalone()

	migrateCmd.Flags().StringS("f", "f", "", "Optionally migrate a single file")
	migrateCmd.Flags().BoolS("help", "help", false, "Print help and exit")
	migrateCmd.Flags().StringS("path", "path", "", "Migrates code for all files in path")
	rootCmd.AddCommand(migrateCmd)

	carapace.Gen(migrateCmd).FlagCompletion(carapace.ActionMap{
		"f":    carapace.ActionFiles(".templ"),
		"path": carapace.ActionDirectories(),
	})
}
