package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bugCmd = &cobra.Command{
	Use:   "bug",
	Short: "start a bug report",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bugCmd).Standalone()
	bugCmd.Flags().SetInterspersed(false)

	bugCmd.Flags().StringS("C", "C", "", "Change to dir before running the command")
	bugCmd.Flags().BoolS("v", "v", false, "print verbose output")
	rootCmd.AddCommand(bugCmd)

	carapace.Gen(bugCmd).FlagCompletion(carapace.ActionMap{
		"C": carapace.ActionDirectories(),
	})
}
