package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/hg_completer/cmd/action"
	"github.com/spf13/cobra"
)

var debuggetbundleCmd = &cobra.Command{
	Use:    "debuggetbundle",
	Short:  "REPO FILE [-H|-C ID]...",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debuggetbundleCmd).Standalone()

	debuggetbundleCmd.Flags().StringArrayP("common", "C", nil, "id of common node")
	debuggetbundleCmd.Flags().StringArrayP("head", "H", nil, "id of head node")
	debuggetbundleCmd.Flags().StringP("type", "t", "", "bundle compression type to use")
	rootCmd.AddCommand(debuggetbundleCmd)

	carapace.Gen(debuggetbundleCmd).FlagCompletion(carapace.ActionMap{
		"common": action.ActionRevisions(),
		"head":   action.ActionRevisions(),
	})
}
