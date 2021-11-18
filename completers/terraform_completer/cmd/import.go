package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/terraform_completer/cmd/action"
	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Associate existing infrastructure with a Terraform resource",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(importCmd).Standalone()

	importCmd.Flags().Bool("allow-missing-config", false, "Allow import when no resource configuration block exists")
	importCmd.Flags().String("config", "", "Path to a directory of Terraform configuration files")
	importCmd.Flags().Bool("ignore-remote-version", false, "A rare option used for the remote backend only")
	importCmd.Flags().String("input", "", "Disable interactive input prompts")
	importCmd.Flags().String("lock", "", "Don't hold a state lock during the operation")
	importCmd.Flags().String("lock-timeout", "", "Duration to retry a state lock")
	importCmd.Flags().Bool("no-color", false, "If specified, output won't contain any color")
	importCmd.Flags().String("var", "", "Set a variable in the Terraform configuration")
	importCmd.Flags().String("var-file", "", "Set variables in the Terraform configuration from a file")
	rootCmd.AddCommand(importCmd)

	importCmd.Flag("config").NoOptDefVal = " "
	importCmd.Flag("input").NoOptDefVal = " "
	importCmd.Flag("lock").NoOptDefVal = " "
	importCmd.Flag("lock-timeout").NoOptDefVal = " "
	importCmd.Flag("var-file").NoOptDefVal = " "

	carapace.Gen(importCmd).FlagCompletion(carapace.ActionMap{
		"config":   carapace.ActionDirectories(),
		"input":    action.ActionBool(),
		"lock":     action.ActionBool(),
		"var-file": carapace.ActionFiles(),
	})

	// TODO positional completion
}
