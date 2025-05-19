package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var signArtifactsCmd = &cobra.Command{
	Use:   "sign-artifacts",
	Short: "Produce artifact signatures and update repodata.json",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(signArtifactsCmd).Standalone()

	signArtifactsCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	rootCmd.AddCommand(signArtifactsCmd)

	carapace.Gen(signArtifactsCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
