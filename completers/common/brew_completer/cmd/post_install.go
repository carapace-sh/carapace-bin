package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var post_installCmd = &cobra.Command{
	Use:   "post_install",
	Short: "Rerun the post-install steps for <formula>",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(post_installCmd).Standalone()

	post_installCmd.Flags().Bool("debug", false, "Display any debugging information.")
	post_installCmd.Flags().Bool("help", false, "Show this message.")
	post_installCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	post_installCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(post_installCmd)
}
