package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/yarn"
	"github.com/spf13/cobra"
)

var patchCmd = &cobra.Command{
	Use:     "patch",
	Short:   "Prepare a package for patching",
	GroupID: "general",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(patchCmd).Standalone()

	patchCmd.Flags().Bool("json", false, "Format the output as an NDJSON stream")
	rootCmd.AddCommand(patchCmd)

	carapace.Gen(patchCmd).PositionalCompletion(
		yarn.ActionDependencies(),
	)
}
