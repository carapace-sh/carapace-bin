package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tofu_completer/cmd/action"
	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import [options] ADDR ID",
	Short: "Associate existing infrastructure with a OpenTofu resource",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(importCmd).Standalone()

	importCmd.Flags().StringS("backup", "backup", "", "Path to backup the existing state file before modifying.")
	importCmd.Flags().BoolS("compact-warnings", "compact-warnings", false, "Show warnings in a more compact form that includes only the summary messages.")
	importCmd.Flags().StringS("config", "config", "", "Path to a directory of OpenTofu configuration files to use for the import plan.")
	importCmd.Flags().BoolS("consolidate-errors", "consolidate-errors", false, "If OpenTofu produces any errors, attempt to consolidate similar messages into a single item.")
	importCmd.Flags().BoolS("consolidate-warnings", "consolidate-warnings", false, "If OpenTofu produces any warnings, do not attempt to consolidate similar messages.")
	importCmd.Flags().BoolS("ignore-remote-version", "ignore-remote-version", false, "A rare option used for the remote backend only.")
	importCmd.Flags().BoolS("input", "input", false, "Disable interactive input prompts.")
	importCmd.Flags().BoolS("json", "json", false, "The output of the command is printed in JSON format.")
	importCmd.Flags().StringS("json-into", "json-into", "", "Produce the same output as -json, but sent directly to the given file.")
	importCmd.Flags().BoolS("lock", "lock", false, "Don't hold a state lock during the operation.")
	importCmd.Flags().StringS("lock-timeout", "lock-timeout", "", "Duration to retry a state lock.")
	importCmd.Flags().BoolS("no-color", "no-color", false, "If specified, output won't contain any color.")
	importCmd.Flags().StringS("parallelism", "parallelism", "", "Limit the number of concurrent operations.")
	importCmd.Flags().StringS("state", "state", "", "Path to read and save state.")
	importCmd.Flags().StringS("state-out", "state-out", "", "Path to write state to that is different than \"-state\".")
	importCmd.Flags().StringArrayS("var", "var", nil, "Set a variable in the OpenTofu configuration.")
	importCmd.Flags().StringS("var-file", "var-file", "", "Set variables in the OpenTofu configuration from a file.")
	rootCmd.AddCommand(importCmd)

	importCmd.Flag("backup").NoOptDefVal = " "
	importCmd.Flag("config").NoOptDefVal = " "
	importCmd.Flag("json-into").NoOptDefVal = " "
	importCmd.Flag("lock-timeout").NoOptDefVal = " "
	importCmd.Flag("parallelism").NoOptDefVal = " "
	importCmd.Flag("state").NoOptDefVal = " "
	importCmd.Flag("state-out").NoOptDefVal = " "
	importCmd.Flag("var-file").NoOptDefVal = " "

	carapace.Gen(importCmd).FlagCompletion(carapace.ActionMap{
		"backup":    carapace.ActionFiles(),
		"config":    carapace.ActionDirectories(),
		"json-into": carapace.ActionFiles(),
		"state":     carapace.ActionFiles(),
		"state-out": carapace.ActionFiles(),
		"var-file":  carapace.ActionFiles(),
	})

	carapace.Gen(importCmd).PositionalCompletion(
		action.ActionResources(importCmd).MultiParts("."),
		carapace.ActionValues(),
	)
}
