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

	untrustCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	untrustCmd.Flags().BoolP("help", "h", false, "Show this message.")
	untrustCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	untrustCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(untrustCmd)
}
