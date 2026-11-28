package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fragments_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List current fragments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fragments_listCmd).Standalone()

	fragments_listCmd.Flags().Bool("aliases", false, "Show only fragments with aliases")
	fragments_listCmd.Flags().Bool("json", false, "Output as JSON")
	fragments_listCmd.Flags().StringP("query", "q", "", "Search for fragments matching these strings")
	fragmentsCmd.AddCommand(fragments_listCmd)
}
