package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/docker-compose_completer/cmd/action"
	"github.com/spf13/cobra"
)

var imagesCmd = &cobra.Command{
	Use:   "images [OPTIONS] [SERVICE...]",
	Short: "List images used by the created containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagesCmd).Standalone()

	imagesCmd.Flags().String("format", "", "Format the output. Values: [table | json]")
	imagesCmd.Flags().BoolP("quiet", "q", false, "Only display IDs")
	rootCmd.AddCommand(imagesCmd)

	carapace.Gen(imagesCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("table", "json"),
	})

	carapace.Gen(imagesCmd).PositionalAnyCompletion(
		action.ActionServices(imagesCmd).FilterArgs(),
	)
}
