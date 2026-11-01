package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var untrustCmd = &cobra.Command{
	Use:     "untrust",
	Short:   "Stop trusting non-official tap formulae, casks or commands",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(untrustCmd).Standalone()

	untrustCmd.Flags().Bool("debug", false, "Display any debugging information.")
	untrustCmd.Flags().Bool("help", false, "Show this message.")
	untrustCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	untrustCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(untrustCmd)
}
