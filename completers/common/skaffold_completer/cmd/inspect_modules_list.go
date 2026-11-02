package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspect_modules_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Print the list of module names that can be invoked with the --module flag in other skaffold commands.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspect_modules_listCmd).Standalone()

	inspect_modules_listCmd.Flags().BoolP("all", "a", false, "Include unnamed modules in the result.")
	applyCommonFlagCompletion(inspect_modules_listCmd)
	inspect_modulesCmd.AddCommand(inspect_modules_listCmd)
}
