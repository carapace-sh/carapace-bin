package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var postinstallCmd = &cobra.Command{
	Use:     "postinstall",
	Short:   "Rerun the post-install steps for <formula>",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(postinstallCmd).Standalone()

	postinstallCmd.Flags().Bool("debug", false, "Display any debugging information.")
	postinstallCmd.Flags().Bool("help", false, "Show this message.")
	postinstallCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	postinstallCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(postinstallCmd)
}
