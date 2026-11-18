package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var analyticsCmd = &cobra.Command{
	Use:     "analytics",
	Short:   "Control Homebrew's anonymous aggregate user behaviour analytics",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(analyticsCmd).Standalone()

	analyticsCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	analyticsCmd.Flags().BoolP("help", "h", false, "Show this message.")
	analyticsCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	analyticsCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(analyticsCmd)

	carapace.Gen(analyticsCmd).PositionalCompletion(
		carapace.ActionValues("off", "on", "state").StyleF(style.ForKeyword),
	)
}
