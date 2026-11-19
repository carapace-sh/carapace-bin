package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var creditsCmd = &cobra.Command{
	Use:    "credits",
	Short:  "Export third party notices to given path (./skaffold-credits by default)",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(creditsCmd).Standalone()

	creditsCmd.Flags().StringP("dir", "d", "./skaffold-credits", "destination directory to place third party licenses")
	applyCommonFlagCompletion(creditsCmd)
	rootCmd.AddCommand(creditsCmd)
}
