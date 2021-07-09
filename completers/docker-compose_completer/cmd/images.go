package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/docker-compose_completer/cmd/action"
	"github.com/spf13/cobra"
)

var imagesCmd = &cobra.Command{
	Use:   "images",
	Short: "List images",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagesCmd).Standalone()

	imagesCmd.Flags().BoolP("quiet", "q", false, "Only display IDs")
	rootCmd.AddCommand(imagesCmd)

	carapace.Gen(imagesCmd).PositionalAnyCompletion(
		action.ActionServices(imagesCmd),
	)
}
