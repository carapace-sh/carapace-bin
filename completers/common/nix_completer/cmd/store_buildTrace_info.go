package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/nix_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var store_buildTrace_infoCmd = &cobra.Command{
	Use:   "info",
	Short: "query information about one or several build traces",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_buildTrace_infoCmd).Standalone()

	common.AddBuiltPathsFlags(store_buildTrace_infoCmd)
	common.AddJSONFlags(store_buildTrace_infoCmd)
	common.AddEvaluationFlags(store_buildTrace_infoCmd)
	common.AddFlakeFlags(store_buildTrace_infoCmd)
	common.AddLoggingFlags(store_buildTrace_infoCmd)

	store_buildTraceCmd.AddCommand(store_buildTrace_infoCmd)

	carapace.Gen(store_buildTrace_infoCmd).FlagCompletion(carapace.ActionMap{})
	carapace.Gen(store_buildTrace_infoCmd).PositionalAnyCompletion(nix.ActionInstallables())
}
