package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls [OPTIONS]",
	Short: "List running compose projects",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lsCmd).Standalone()

	lsCmd.Flags().BoolP("all", "a", false, "Show all stopped Compose projects")
	lsCmd.Flags().String("filter", "", "Filter output based on conditions provided")
	lsCmd.Flags().String("format", "", "Format the output. Values: [table | json]")
	lsCmd.Flags().BoolP("quiet", "q", false, "Only display project names")
	rootCmd.AddCommand(lsCmd)

	carapace.Gen(lsCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("table", "json"),
	})
}
