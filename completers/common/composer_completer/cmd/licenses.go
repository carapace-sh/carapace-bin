package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licensesCmd = &cobra.Command{
	Use:   "licenses",
	Short: "Shows information about licenses of dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licensesCmd).Standalone()

	licensesCmd.Flags().StringP("format", "f", "", "Format of the output: text, json or summary")
	licensesCmd.Flags().Bool("locked", false, "Shows licenses from the lock file instead of installed packages.")
	licensesCmd.Flags().Bool("no-dev", false, "Disables search in require-dev packages.")

	rootCmd.AddCommand(licensesCmd)

	carapace.Gen(licensesCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "summary", "text"),
	})
}
