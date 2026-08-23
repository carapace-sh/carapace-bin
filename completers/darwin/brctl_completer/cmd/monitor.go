package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(monitorCmd)
}

var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "monitor a container using NSMetadataQuery",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(monitorCmd).Standalone()

	monitorCmd.Flags().String("scope", "", "restrict the NSMDQ scope to DOCS, DATA, or BOTH")

	carapace.Gen(monitorCmd).FlagCompletion(carapace.ActionMap{
		"scope": carapace.ActionValues("DOCS", "DATA", "BOTH"),
	})
}