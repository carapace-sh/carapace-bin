package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/docker-compose_completer/cmd/action"
	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:   "export [OPTIONS] SERVICE",
	Short: "Export a service container's filesystem as a tar archive",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(exportCmd).Standalone()

	exportCmd.Flags().String("index", "", "index of the container if service has multiple replicas.")
	exportCmd.Flags().StringP("output", "o", "", "Write to a file, instead of STDOUT")
	rootCmd.AddCommand(exportCmd)

	carapace.Gen(exportCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionFiles(),
	})

	carapace.Gen(exportCmd).PositionalCompletion(
		action.ActionServices(exportCmd),
	)
}
