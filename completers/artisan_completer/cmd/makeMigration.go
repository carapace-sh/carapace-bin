package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var makeMigrationCmd = &cobra.Command{
	Use:   "make:migration [--create [CREATE]] [--table [TABLE]] [--path [PATH]] [--realpath] [--fullpath] [--] <name>",
	Short: "Create a new migration file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(makeMigrationCmd).Standalone()

	makeMigrationCmd.Flags().String("create", "", "The table to be created")
	makeMigrationCmd.Flags().Bool("fullpath", false, "Output the full path of the migration (Deprecated)")
	makeMigrationCmd.Flags().String("path", "", "The location where the migration file should be created")
	makeMigrationCmd.Flags().Bool("realpath", false, "Indicate any provided migration file paths are pre-resolved absolute paths")
	makeMigrationCmd.Flags().String("table", "", "The table to migrate")
	rootCmd.AddCommand(makeMigrationCmd)
}
