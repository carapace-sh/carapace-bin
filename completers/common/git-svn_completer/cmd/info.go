package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show info about the latest SVN revision on the current branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(infoCmd).Standalone()

	infoCmd.Flags().Bool("url", false, "Output only the value of the URL field")
	rootCmd.AddCommand(infoCmd)

	carapace.Gen(infoCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
