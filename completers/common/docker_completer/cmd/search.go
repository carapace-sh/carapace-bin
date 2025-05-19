package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:     "search",
	Short:   "Search the Docker Hub for images",
	GroupID: "common",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(searchCmd).Standalone()

	searchCmd.Flags().StringP("filter", "f", "", "Filter output based on conditions provided")
	searchCmd.Flags().String("format", "", "Pretty-print search using a Go template")
	searchCmd.Flags().Int("limit", 0, "Max number of search results")
	searchCmd.Flags().Bool("no-trunc", false, "Don't truncate output")
	rootCmd.AddCommand(searchCmd)
}
