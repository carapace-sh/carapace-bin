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

	refreshCmd.Flags().StringS("backup", "backup", "", "Path to backup the existing state file before modifying.")
	refreshCmd.Flags().BoolS("compact-warnings", "compact-warnings", false, "Show warnings in a more compact form that includes only the summary messages.")
	refreshCmd.Flags().BoolS("concise", "concise", false, "Disable progress-related messages.")
	refreshCmd.Flags().BoolS("consolidate-errors", "consolidate-errors", false, "If OpenTofu produces any errors, attempt to consolidate similar messages into a single item.")
	refreshCmd.Flags().BoolS("consolidate-warnings", "consolidate-warnings", false, "If OpenTofu produces any warnings, do not attempt to consolidate similar messages.")
	refreshCmd.Flags().StringS("exclude", "exclude", "", "Resource to exclude. Operation will be limited to all resources not included in this and all of their dependencies.")
	refreshCmd.Flags().StringS("exclude-file", "exclude-file", "", "Similar to -exclude, but specifies zero or more resource addresses from a file.")
	refreshCmd.Flags().BoolS("input", "input", false, "Ask for input for variables if not directly set.")
	refreshCmd.Flags().BoolS("json", "json", false, "Produce output in a machine-readable JSON format.")
	refreshCmd.Flags().StringS("json-into", "json-into", "", "Produce the same output as -json, but sent directly to the given file.")
	refreshCmd.Flags().BoolS("lock", "lock", false, "Don't hold a state lock during the operation.")
	refreshCmd.Flags().StringS("lock-timeout", "lock-timeout", "", "Duration to retry a state lock.")
	refreshCmd.Flags().BoolS("no-color", "no-color", false, "If specified, output won't contain any color.")
	refreshCmd.Flags().StringS("parallelism", "parallelism", "", "Limit the number of parallel resource operations.")
	refreshCmd.Flags().BoolS("refresh", "refresh", false, "Skip checking for external changes to remote objects while creating the plan.")
	refreshCmd.Flags().StringS("state", "state", "", "Path to read and save state.")
	refreshCmd.Flags().StringS("state-out", "state-out", "", "Path to write state to that is different than \"-state\".")
	refreshCmd.Flags().StringS("target", "target", "", "Resource to target. Operation will be limited to this resource and all of its dependencies.")
	refreshCmd.Flags().StringS("target-file", "target-file", "", "Similar to -target, but specifies zero or more resource addresses from a file.")
	refreshCmd.Flags().StringArrayS("var", "var", nil, "Set a variable in the OpenTofu configuration.")
	refreshCmd.Flags().StringS("var-file", "var-file", "", "Set variables in the OpenTofu configuration from a file.")
	rootCmd.AddCommand(refreshCmd)

	refreshCmd.Flag("backup").NoOptDefVal = " "
	refreshCmd.Flag("exclude").NoOptDefVal = " "
	refreshCmd.Flag("exclude-file").NoOptDefVal = " "
	refreshCmd.Flag("json-into").NoOptDefVal = " "
	refreshCmd.Flag("lock-timeout").NoOptDefVal = " "
	refreshCmd.Flag("parallelism").NoOptDefVal = " "
	refreshCmd.Flag("state").NoOptDefVal = " "
	refreshCmd.Flag("state-out").NoOptDefVal = " "
	refreshCmd.Flag("target").NoOptDefVal = " "
	refreshCmd.Flag("target-file").NoOptDefVal = " "
	refreshCmd.Flag("var-file").NoOptDefVal = " "

	carapace.Gen(refreshCmd).FlagCompletion(carapace.ActionMap{
		"backup":       carapace.ActionFiles(),
		"exclude":      action.ActionResources(refreshCmd).MultiParts("."),
		"exclude-file": carapace.ActionFiles(),
		"json-into":    carapace.ActionFiles(),
		"state":        carapace.ActionFiles(),
		"state-out":    carapace.ActionFiles(),
		"target":       action.ActionResources(refreshCmd).MultiParts("."),
		"target-file":  carapace.ActionFiles(),
		"var-file":     carapace.ActionFiles(),
	})
}
