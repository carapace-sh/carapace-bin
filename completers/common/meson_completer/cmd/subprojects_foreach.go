package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var subprojects_foreachCmd = &cobra.Command{
	Use:   "foreach",
	Short: "Execute a command in each subproject directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(subprojects_foreachCmd).Standalone()

	subprojects_foreachCmd.Flags().Bool("allow-insecure", false, "Allow insecure server connections.")
	subprojects_foreachCmd.Flags().String("filter", "", "Patterns of subprojects to operate on (default: all)")
	subprojects_foreachCmd.Flags().StringP("num-processes", "j", "", "How many parallel processes to use (Since 0.59.0).")
	subprojects_foreachCmd.Flags().String("sourcedir", "", "Path to source directory")
	subprojects_foreachCmd.Flags().String("types", "", "Comma-separated list of subproject types.")
	applyCommonFlagCompletion(subprojects_foreachCmd)
	subprojectsCmd.AddCommand(subprojects_foreachCmd)
}
