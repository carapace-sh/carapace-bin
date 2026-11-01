package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tofu_completer/cmd/action"
	"github.com/spf13/cobra"
)

var destroyCmd = &cobra.Command{
	Use:     "destroy [options] [PLAN]",
	Short:   "Destroy previously-created infrastructure",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(destroyCmd).Standalone()

	destroyCmd.Flags().BoolS("auto-approve", "auto-approve", false, "Skip interactive approval of plan before applying.")
	destroyCmd.Flags().StringS("backup", "backup", "", "Path to backup the existing state file before modifying.")
	destroyCmd.Flags().BoolS("compact-warnings", "compact-warnings", false, "Show warnings in a more compact form that includes only the summary messages.")
	destroyCmd.Flags().BoolS("concise", "concise", false, "Disable progress-related messages.")
	destroyCmd.Flags().BoolS("consolidate-errors", "consolidate-errors", false, "If OpenTofu produces any errors, attempt to consolidate similar messages into a single item.")
	destroyCmd.Flags().BoolS("consolidate-warnings", "consolidate-warnings", false, "If OpenTofu produces any warnings, do not attempt to consolidate similar messages.")
	destroyCmd.Flags().StringArrayS("deprecation", "deprecation", nil, "Specify what type of warnings are shown.")
	destroyCmd.Flags().BoolS("destroy", "destroy", false, "Select the \"destroy\" planning mode.")
	destroyCmd.Flags().StringArrayS("exclude", "exclude", nil, "Limit the operation to not operate on the given module, resource, or resource instance.")
	destroyCmd.Flags().StringArrayS("exclude-file", "exclude-file", nil, "Similar to -exclude, but specifies zero or more resource addresses from a file.")
	destroyCmd.Flags().BoolS("input", "input", false, "Ask for input for variables if not directly set.")
	destroyCmd.Flags().BoolS("json", "json", false, "Produce output in a machine-readable JSON format.")
	destroyCmd.Flags().StringS("json-into", "json-into", "", "Produce the same output as -json, but sent directly to the given file.")
	destroyCmd.Flags().BoolS("lock", "lock", false, "Don't hold a state lock during the operation.")
	destroyCmd.Flags().StringS("lock-timeout", "lock-timeout", "", "Duration to retry a state lock.")
	destroyCmd.Flags().BoolS("no-color", "no-color", false, "If specified, output won't contain any color.")
	destroyCmd.Flags().StringS("parallelism", "parallelism", "", "Limit the number of parallel resource operations.")
	destroyCmd.Flags().BoolS("refresh", "refresh", false, "Skip checking for external changes to remote objects while creating the plan.")
	destroyCmd.Flags().BoolS("refresh-only", "refresh-only", false, "Select the \"refresh only\" planning mode.")
	destroyCmd.Flags().StringArrayS("replace", "replace", nil, "Force replacement of a particular resource instance using its resource address.")
	destroyCmd.Flags().BoolS("show-sensitive", "show-sensitive", false, "If specified, sensitive values will be displayed.")
	destroyCmd.Flags().StringS("state", "state", "", "Path to read and save state.")
	destroyCmd.Flags().StringS("state-out", "state-out", "", "Path to write state to that is different than \"-state\".")
	destroyCmd.Flags().BoolS("suppress-forget-errors", "suppress-forget-errors", false, "Suppress the error that occurs when a destroy operation completes successfully but leaves forgotten instances behind.")
	destroyCmd.Flags().StringArrayS("target", "target", nil, "Limit the operation to only the given module, resource, or resource instance.")
	destroyCmd.Flags().StringArrayS("target-file", "target-file", nil, "Similar to -target, but specifies zero or more resource addresses from a file.")
	destroyCmd.Flags().StringArrayS("var", "var", nil, "Set a variable in the OpenTofu configuration.")
	destroyCmd.Flags().StringS("var-file", "var-file", "", "Set variables in the OpenTofu configuration from a file.")
	rootCmd.AddCommand(destroyCmd)

	carapace.Gen(destroyCmd).FlagCompletion(carapace.ActionMap{
		"backup":       carapace.ActionFiles(),
		"deprecation":  action.ActionDeprecationLevels(),
		"exclude":      action.ActionResources(destroyCmd).MultiParts("."),
		"exclude-file": carapace.ActionFiles(),
		"json-into":    carapace.ActionFiles(),
		"replace":      action.ActionResources(destroyCmd).MultiParts("."),
		"state":        carapace.ActionFiles(),
		"state-out":    carapace.ActionFiles(),
		"target":       action.ActionResources(destroyCmd).MultiParts("."),
		"target-file":  carapace.ActionFiles(),
		"var-file":     carapace.ActionFiles(),
	})

	carapace.Gen(destroyCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
