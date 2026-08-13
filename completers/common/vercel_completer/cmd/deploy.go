package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy your project to Vercel",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deployCmd).Standalone()

	rootCmd.AddCommand(deployCmd)

	carapace.Gen(deployCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
