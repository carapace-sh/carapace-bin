package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var findConfigsCmd = &cobra.Command{
	Use:    "find-configs",
	Short:  "Find in a given directory all skaffold yamls files that are parseable or upgradeable with their versions.",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(findConfigsCmd).Standalone()

	findConfigsCmd.Flags().StringP("directory", "d", ".", "Root directory to lookup the config files.")
	findConfigsCmd.Flags().StringP("output", "o", "table", "Result format, default to table. [(-o|--output=)json|table]")
	applyCommonFlagCompletion(findConfigsCmd)
	rootCmd.AddCommand(findConfigsCmd)
}
