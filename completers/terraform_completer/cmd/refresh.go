package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/terraform_completer/cmd/action"
	"github.com/spf13/cobra"
)

var refreshCmd = &cobra.Command{
	Use:   "refresh",
	Short: "Update the state to match remote systems",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(refreshCmd).Standalone()

	refreshCmd.Flags().Bool("compact-warnings", false, "show warnings in a more compact form")
	refreshCmd.Flags().String("input", "", "Ask for input for variables if not directly set")
	refreshCmd.Flags().String("lock", "", "Don't hold a state lock during the operation")
	refreshCmd.Flags().String("lock-timeout", "", "Duration to retry a state lock")
	refreshCmd.Flags().Bool("no-color", false, "If specified, output won't contain any color")
	refreshCmd.Flags().StringArray("target", []string{}, "Resource to target")
	refreshCmd.Flags().StringArray("var", []string{}, "Set a variable in the Terraform configuration")
	refreshCmd.Flags().String("var-file", "", "Set variables in the Terraform configuration from a file")
	rootCmd.AddCommand(refreshCmd)

	refreshCmd.Flag("input").NoOptDefVal = " "
	refreshCmd.Flag("lock").NoOptDefVal = " "
	refreshCmd.Flag("lock-timeout").NoOptDefVal = " "
	refreshCmd.Flag("target").NoOptDefVal = " "
	refreshCmd.Flag("var-file").NoOptDefVal = " "

	// TODO complete target/var
	carapace.Gen(refreshCmd).FlagCompletion(carapace.ActionMap{
		"input":    action.ActionBool(),
		"lock":     action.ActionBool(),
		"var-file": carapace.ActionFiles(),
	})
}
