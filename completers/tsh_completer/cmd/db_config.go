package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var db_configCmd = &cobra.Command{
	Use:   "config",
	Short: "Print database connection information. Useful when configuring GUI clients.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(db_configCmd).Standalone()

	db_configCmd.Flags().String("db", "", "Print information for the specified database.")
	db_configCmd.Flags().StringP("format", "f", "", "Print format: \"text\" to print in table format (default), \"cmd\" to print connect command, \"json\" or \"yaml\" to print in JSON or YAML.")
	db_configCmd.Flag("db").Hidden = true
	dbCmd.AddCommand(db_configCmd)

	carapace.Gen(db_configCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValuesDescribed(
			"text", "print in table format",
			"cmd", "print connect command",
			"json", "print int JSON",
			"yaml", "print in YAML",
		),
	})
}
