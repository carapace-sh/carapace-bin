package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apps_configCmd = &cobra.Command{
	Use:   "config",
	Short: "Print app connection information.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apps_configCmd).Standalone()

	apps_configCmd.Flags().StringP("format", "f", "", "Optional print format, one of: \"uri\" to print app address, \"ca\" to print CA cert path, \"cert\" to print cert path, \"key\" print key path, \"curl\" to print example curl command, \"json\" or \"yaml\" to print everything as JSON or YAML.")
	appsCmd.AddCommand(apps_configCmd)

	carapace.Gen(apps_configCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValuesDescribed(
			"uri", "app address",
			"ca", "CA cert path",
			"cert", "cert path",
			"key", "key path",
			"curl", "example curl command",
			"json", "print everything as JSON",
			"yaml", "print everything as YAML",
		),
	})

	// TODO complete apps
}
