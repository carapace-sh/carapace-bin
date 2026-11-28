package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/kind_completer/cmd/action"
	"github.com/spf13/cobra"
)

var export_logsCmd = &cobra.Command{
	Use:   "logs [output-dir]",
	Short: "Exports logs to a tempdir or [output-dir] if specified",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(export_logsCmd).Standalone()

	export_logsCmd.Flags().StringP("name", "n", "kind", "the cluster context name")

	carapace.Gen(export_logsCmd).FlagCompletion(carapace.ActionMap{
		"name": action.ActionClusterNames(),
	})

	carapace.Gen(export_logsCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)

	exportCmd.AddCommand(export_logsCmd)
}
