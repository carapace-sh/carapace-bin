package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var subprojectsCmd = &cobra.Command{
	Use:   "subprojects",
	Short: "Manage subprojects",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(subprojectsCmd).Standalone()

	applyCommonFlagCompletion(subprojectsCmd)
	rootCmd.AddCommand(subprojectsCmd)
}
