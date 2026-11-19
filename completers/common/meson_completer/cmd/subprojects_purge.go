package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var subprojects_purgeCmd = &cobra.Command{
	Use:   "purge",
	Short: "Remove all wrap-based subproject artifacts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(subprojects_purgeCmd).Standalone()

	subprojects_purgeCmd.Flags().Bool("allow-insecure", false, "Allow insecure server connections.")
	subprojects_purgeCmd.Flags().Bool("confirm", false, "Confirm the removal of subproject artifacts")
	subprojects_purgeCmd.Flags().Bool("include-cache", false, "Remove the package cache as well")
	subprojects_purgeCmd.Flags().StringP("num-processes", "j", "", "How many parallel processes to use (Since 0.59.0).")
	subprojects_purgeCmd.Flags().String("sourcedir", "", "Path to source directory")
	subprojects_purgeCmd.Flags().String("types", "", "Comma-separated list of subproject types.")
	applyCommonFlagCompletion(subprojects_purgeCmd)
	subprojectsCmd.AddCommand(subprojects_purgeCmd)
}
