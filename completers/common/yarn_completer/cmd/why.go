package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/yarn"
	"github.com/spf13/cobra"
)

var whyCmd = &cobra.Command{
	Use:     "why",
	Short:   "Display the reason why a package is needed",
	GroupID: "general",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(whyCmd).Standalone()

	whyCmd.Flags().Bool("json", false, "Format the output as an NDJSON stream")
	whyCmd.Flags().Bool("peers", false, "Also print the peer dependencies that match the specified name")
	whyCmd.Flags().BoolP("recursive", "R", false, "List, for each workspace, what are all the paths that lead to the dependency")
	rootCmd.AddCommand(whyCmd)

	carapace.Gen(whyCmd).PositionalCompletion(
		yarn.ActionDependencies(),
	)
}
