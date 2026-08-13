package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var contractCmd = &cobra.Command{
	Use:   "contract",
	Short: "Show contract information for all billing periods",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(contractCmd).Standalone()

	contractCmd.Flags().String("format", "", "Output format")
	contractCmd.Flags().Bool("json", false, "Output as JSON")

	rootCmd.AddCommand(contractCmd)

	carapace.Gen(contractCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("plain", "json"),
	})
}
