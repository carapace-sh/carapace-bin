package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tokens_rmCmd = &cobra.Command{
	Use:     "rm",
	Aliases: []string{"remove", "delete"},
	Short:   "Delete a personal authentication token by ID",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tokens_rmCmd).Standalone()

	tokens_rmCmd.Flags().String("format", "", "Output format")
	tokens_rmCmd.Flags().Bool("json", false, "Output as JSON")

	tokensCmd.AddCommand(tokens_rmCmd)

	carapace.Gen(tokens_rmCmd).PositionalCompletion(
		carapace.ActionValues(),
	)
}
