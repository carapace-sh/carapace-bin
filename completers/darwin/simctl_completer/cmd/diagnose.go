package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var diagnoseCmd = &cobra.Command{
	Use:   "diagnose",
	Short: "Collect diagnostic information and logs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(diagnoseCmd).Standalone()
	diagnoseCmd.Flags().StringP("log", "l", "", "Write output to a specific path")
	rootCmd.AddCommand(diagnoseCmd)
	carapace.Gen(diagnoseCmd).FlagCompletion(carapace.ActionMap{
		"log": carapace.ActionFiles(),
	})
}
