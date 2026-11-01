package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appcRoutesCmd = &cobra.Command{
	Use:   "appc-routes",
	Short: "Print the current app connector routes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appcRoutesCmd).Standalone()

	appcRoutesCmd.Flags().Bool("all", false, "print learned domains, routes and extra policy configured routes")
	appcRoutesCmd.Flags().Bool("map", false, "print the map of learned domains to routes")
	appcRoutesCmd.Flags().Bool("n", false, "print the total number of routes this node advertises")
	rootCmd.AddCommand(appcRoutesCmd)
}
