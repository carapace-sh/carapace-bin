package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/vercel_completer/cmd/action"
	"github.com/spf13/cobra"
)

var redeployCmd = &cobra.Command{
	Use:   "redeploy",
	Short: "Rebuild and deploy a previous deployment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redeployCmd).Standalone()

	redeployCmd.Flags().Bool("no-wait", false, "Don't wait for the redeploy to finish")
	redeployCmd.Flags().String("target", "", "Redeploy to a specific target environment")

	rootCmd.AddCommand(redeployCmd)

	carapace.Gen(redeployCmd).PositionalCompletion(
		action.ActionDeployments(redeployCmd),
	)
}
