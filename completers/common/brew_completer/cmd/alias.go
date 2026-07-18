package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var aliasCmd = &cobra.Command{
	Use:     "alias",
	Short:   "Show an alias's command, or print the whole list if no alias is given",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(aliasCmd).Standalone()

	aliasCmd.Flags().Bool("debug", false, "Display any debugging information.")
	aliasCmd.Flags().Bool("help", false, "Show this message.")
	aliasCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	aliasCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(aliasCmd)
}
