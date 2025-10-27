package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate configuration file from v1 to v2.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrateCmd).Standalone()

	migrateCmd.Flags().StringP("config", "c", "", "Read config from file path `PATH`")
	migrateCmd.Flags().String("format", "", "Output file format.")
	migrateCmd.Flags().Bool("no-config", false, "Don't read config file")
	migrateCmd.Flags().Bool("skip-validation", false, "Skip validation of the configuration file against the JSON Schema for v1.")
	rootCmd.AddCommand(migrateCmd)

	carapace.Gen(migrateCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionFiles(),
		"format": carapace.ActionValues("yml", "yaml", "toml", "json"),
	})
}
