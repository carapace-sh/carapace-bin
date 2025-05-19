package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/minikube_completer/cmd/action"
	"github.com/spf13/cobra"
)

var image_rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove one or more images",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_rmCmd).Standalone()
	imageCmd.AddCommand(image_rmCmd)

	carapace.Gen(image_rmCmd).PositionalCompletion(
		action.ActionImages(),
	)
}
