package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tofu_completer/cmd/action"
	"github.com/spf13/cobra"
)

var refreshCmd = &cobra.Command{
	Use:   "refresh [options]",
	Short: "Update the state to match remote systems",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(refreshCmd).Standalone()

	refreshCmd.Flags().BoolS("compact-warnings", "compact-warnings", false, "show warnings in a more compact form")
	refreshCmd.Flags().BoolS("input", "input", false, "Ask for input for variables if not directly set")
	refreshCmd.Flags().BoolS("lock", "lock", false, "Don't hold a state lock during the operation")
	refreshCmd.Flags().StringS("lock-timeout", "lock-timeout", "", "Duration to retry a state lock")
	refreshCmd.Flags().BoolS("no-color", "no-color", false, "If specified, output won't contain any color")
	refreshCmd.Flags().StringS("parallelism", "parallelism", "", "Limit the number of parallel resource operations")
	refreshCmd.Flags().StringArrayS("target", "target", nil, "Resource to target")
	refreshCmd.Flags().StringArrayS("var", "var", nil, "Set a variable in the Terraform configuration")
	refreshCmd.Flags().StringS("var-file", "var-file", "", "Set variables in the Terraform configuration from a file")
	rootCmd.AddCommand(refreshCmd)

	refreshCmd.Flag("lock-timeout").NoOptDefVal = " "
	refreshCmd.Flag("parallelism").NoOptDefVal = " "
	refreshCmd.Flag("target").NoOptDefVal = " "
	refreshCmd.Flag("var-file").NoOptDefVal = " "

	// TODO complete var
	carapace.Gen(refreshCmd).FlagCompletion(carapace.ActionMap{
		"target":   action.ActionResources(refreshCmd).MultiParts("."),
		"var-file": carapace.ActionFiles(),
	})
}
