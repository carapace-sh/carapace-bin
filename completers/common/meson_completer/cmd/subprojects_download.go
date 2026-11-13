package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var subprojects_downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Ensure subprojects are fetched, even if not in use.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(subprojects_downloadCmd).Standalone()

	subprojects_downloadCmd.Flags().Bool("allow-insecure", false, "Allow insecure server connections.")
	subprojects_downloadCmd.Flags().StringP("num-processes", "j", "", "How many parallel processes to use (Since 0.59.0).")
	subprojects_downloadCmd.Flags().String("sourcedir", "", "Path to source directory")
	subprojects_downloadCmd.Flags().String("types", "", "Comma-separated list of subproject types.")
	applyCommonFlagCompletion(subprojects_downloadCmd)
	subprojectsCmd.AddCommand(subprojects_downloadCmd)
}
