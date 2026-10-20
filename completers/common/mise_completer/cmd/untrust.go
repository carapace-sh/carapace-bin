package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var untrustCmd = &cobra.Command{
	Use:   "untrust",
	Short: "Remove explicit trust for a config",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(untrustCmd).Standalone()

	untrustCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(untrustCmd)

	carapace.Gen(untrustCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
