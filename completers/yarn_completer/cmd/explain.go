package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/yarn"
	"github.com/spf13/cobra"
)

var explainCmd = &cobra.Command{
	Use:     "explain",
	Short:   "Explain an error code",
	GroupID: "general",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(explainCmd).Standalone()

	explainCmd.Flags().Bool("json", false, "Format the output as an NDJSON stream")
	rootCmd.AddCommand(explainCmd)

	carapace.Gen(explainCmd).PositionalCompletion(
		yarn.ActionErrorCodes(),
	)
}
