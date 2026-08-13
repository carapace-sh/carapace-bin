package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tokens_lsCmd = &cobra.Command{
	Use:     "ls",
	Aliases: []string{"list"},
	Short:   "List your personal authentication tokens",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tokens_lsCmd).Standalone()

	tokens_lsCmd.Flags().String("format", "", "Output format")
	tokens_lsCmd.Flags().Bool("json", false, "Output as JSON")
	tokens_lsCmd.Flags().String("limit", "", "Number of results per page")

	tokensCmd.AddCommand(tokens_lsCmd)

	carapace.Gen(tokens_lsCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("plain", "json"),
	})
}
