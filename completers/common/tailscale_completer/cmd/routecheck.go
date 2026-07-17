package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var routecheckCmd = &cobra.Command{
	Use:   "routecheck",
	Short: "Print a reachability report for routes with multiple paths",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(routecheckCmd).Standalone()

	routecheckCmd.Flags().String("format", "", "output format")
	routecheckCmd.Flags().Bool("json", false, "output in JSON format")
	routecheckCmd.Flags().Bool("probe", false, "probe now to generate a new reachability report")
	rootCmd.AddCommand(routecheckCmd)

	carapace.Gen(routecheckCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "json-line"),
	})
}
