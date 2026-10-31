package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var subprojects_packagefilesCmd = &cobra.Command{
	Use:   "packagefiles",
	Short: "Manage the packagefiles overlay",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(subprojects_packagefilesCmd).Standalone()

	subprojects_packagefilesCmd.Flags().Bool("allow-insecure", false, "Allow insecure server connections.")
	subprojects_packagefilesCmd.Flags().Bool("apply", false, "Apply packagefiles to the subproject")
	subprojects_packagefilesCmd.Flags().StringP("num-processes", "j", "", "How many parallel processes to use (Since 0.59.0).")
	subprojects_packagefilesCmd.Flags().Bool("save", false, "Save packagefiles from the subproject")
	subprojects_packagefilesCmd.Flags().String("sourcedir", "", "Path to source directory")
	subprojects_packagefilesCmd.Flags().String("types", "", "Comma-separated list of subproject types.")
	applyCommonFlagCompletion(subprojects_packagefilesCmd)
	subprojectsCmd.AddCommand(subprojects_packagefilesCmd)
}
