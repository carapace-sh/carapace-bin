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

	developerCmd.Flags().Bool("debug", false, "Display any debugging information.")
	developerCmd.Flags().Bool("help", false, "Show this message.")
	developerCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	developerCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(developerCmd)

	carapace.Gen(developerCmd).PositionalCompletion(
		carapace.ActionValues("on", "off").StyleF(style.ForKeyword),
	)
}
