package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var subprojects_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update all subprojects from wrap files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(subprojects_updateCmd).Standalone()

	subprojects_updateCmd.Flags().Bool("allow-insecure", false, "Allow insecure server connections.")
	subprojects_updateCmd.Flags().StringP("num-processes", "j", "", "How many parallel processes to use (Since 0.59.0).")
	subprojects_updateCmd.Flags().Bool("rebase", false, "Rebase your branch on top of wrap's revision.")
	subprojects_updateCmd.Flags().Bool("reset", false, "Checkout wrap's revision and hard reset to that commit.")
	subprojects_updateCmd.Flags().String("sourcedir", "", "Path to source directory")
	subprojects_updateCmd.Flags().String("types", "", "Comma-separated list of subproject types.")
	applyCommonFlagCompletion(subprojects_updateCmd)
	subprojectsCmd.AddCommand(subprojects_updateCmd)
}
