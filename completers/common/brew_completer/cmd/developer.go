package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var developerCmd = &cobra.Command{
	Use:     "developer",
	Short:   "Control Homebrew's developer mode",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(developerCmd).Standalone()

	developerCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	developerCmd.Flags().BoolP("help", "h", false, "Show this message.")
	developerCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	developerCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(developerCmd)

	carapace.Gen(developerCmd).PositionalCompletion(
		carapace.ActionValues("off", "on").StyleF(style.ForKeyword),
	)
}
