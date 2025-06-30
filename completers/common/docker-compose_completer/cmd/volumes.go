package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/docker-compose_completer/cmd/action"
	"github.com/spf13/cobra"
)

var volumesCmd = &cobra.Command{
	Use:   "volumes [OPTIONS] [SERVICE...]",
	Short: "List volumes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volumesCmd).Standalone()

	volumesCmd.Flags().String("format", "", "Format output using a custom template:")
	volumesCmd.Flags().BoolP("quiet", "q", false, "Only display volume names")
	rootCmd.AddCommand(volumesCmd)

	carapace.Gen(volumesCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("table", "json"),
	})

	carapace.Gen(volumesCmd).PositionalAnyCompletion(
		action.ActionServices(volumesCmd),
	)
}
