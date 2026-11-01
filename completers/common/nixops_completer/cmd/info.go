package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "show the state of the deployment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(infoCmd).Standalone()

	infoCmd.Flags().Bool("all", false, "Show all deployments")
	infoCmd.Flags().Bool("no-eval", false, "Do not evaluate the deployment specification")
	infoCmd.Flags().Bool("plain", false, "Do not pretty-print the output")

	rootCmd.AddCommand(infoCmd)
}
