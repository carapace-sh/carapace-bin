package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var consoleCmd = &cobra.Command{
	Use:   "console",
	Short: "Try Terraform expressions at an interactive command prompt",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(consoleCmd).Standalone()

	consoleCmd.Flags().String("state", "", "Legacy option for the local backend only.")
	consoleCmd.Flags().StringArray("var", []string{}, "Set a variable in the Terraform configuration.")
	consoleCmd.Flags().String("var-file", "", "Set variables in the Terraform configuration from a file.")

	consoleCmd.Flag("state").NoOptDefVal = " "
	consoleCmd.Flag("var-file").NoOptDefVal = " "

	rootCmd.AddCommand(consoleCmd)

	carapace.Gen(consoleCmd).FlagCompletion(carapace.ActionMap{
		"state":    carapace.ActionFiles(),
		"var-file": carapace.ActionFiles(),
	})
}
