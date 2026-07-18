package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var unaliasCmd = &cobra.Command{
	Use:     "unalias",
	Short:   "Remove aliases",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unaliasCmd).Standalone()

	unaliasCmd.Flags().Bool("debug", false, "Display any debugging information.")
	unaliasCmd.Flags().Bool("help", false, "Show this message.")
	unaliasCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	unaliasCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(unaliasCmd)
}
