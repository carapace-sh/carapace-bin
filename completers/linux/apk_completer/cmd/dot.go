package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/apk_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/completers/linux/apk_completer/cmd/common"
	"github.com/spf13/cobra"
)

var dotCmd = &cobra.Command{
	Use:     "dot",
	Short:   "Render dependencies as graphviz graphs",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: "querying package information",
}

func init() {
	carapace.Gen(dotCmd).Standalone()

	dotCmd.Flags().Bool("errors", false, "Consider only packages with errors")
	dotCmd.Flags().Bool("installed", false, "Consider only installed packages")
	common.AddGlobalFlags(dotCmd)
	common.AddSourceFlags(dotCmd)
	rootCmd.AddCommand(dotCmd)

	carapace.Gen(dotCmd).PositionalAnyCompletion(
		action.ActionPackageSearch(dotCmd).FilterArgs(),
	)
}
