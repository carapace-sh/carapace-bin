package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import [options] ADDR ID",
	Short: "Associate existing infrastructure with a Terraform resource",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(importCmd).Standalone()

	importCmd.Flags().StringS("config", "config", "", "Path to a directory of Terraform configuration files")
	importCmd.Flags().BoolS("ignore-remote-version", "ignore-remote-version", false, "A rare option used for the remote backend only")
	importCmd.Flags().BoolS("input", "input", false, "Disable interactive input prompts")
	importCmd.Flags().BoolS("lock", "lock", false, "Don't hold a state lock during the operation")
	importCmd.Flags().StringS("lock-timeout", "lock-timeout", "", "Duration to retry a state lock")
	importCmd.Flags().BoolS("no-color", "no-color", false, "If specified, output won't contain any color")
	importCmd.Flags().StringSliceS("var", "var", nil, "Set a variable in the Terraform configuration")
	importCmd.Flags().StringSliceS("var-file", "var-file", nil, "Set variables in the Terraform configuration from a file")
	rootCmd.AddCommand(importCmd)

	importCmd.Flag("config").NoOptDefVal = " "
	importCmd.Flag("lock-timeout").NoOptDefVal = " "
	importCmd.Flag("var-file").NoOptDefVal = " "

	carapace.Gen(importCmd).FlagCompletion(carapace.ActionMap{
		"config":   carapace.ActionDirectories(),
		"var-file": carapace.ActionFiles(),
	})

	// TODO positional completion
}
