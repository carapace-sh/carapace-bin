package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/yarn"
	"github.com/spf13/cobra"
)

var npm_infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show information about a package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(npm_infoCmd).Standalone()

	npm_infoCmd.Flags().StringP("fields", "f", "", "A comma-separated list of manifest fields that should be displayed")
	npm_infoCmd.Flags().Bool("json", false, "Format the output as an NDJSON stream")
	npmCmd.AddCommand(npm_infoCmd)

	// TODO fields

	carapace.Gen(npm_infoCmd).PositionalAnyCompletion(
		yarn.ActionDependencies(),
	)
}
