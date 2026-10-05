package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tofu_completer/cmd/action"
	"github.com/spf13/cobra"
)

var applyCmd = &cobra.Command{
	Use:     "apply [options] [PLAN]",
	Short:   "Create or update infrastructure",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applyCmd).Standalone()

	applyCmd.Flags().BoolS("auto-approve", "auto-approve", false, "Skip interactive approval of plan before applying.")
	applyCmd.Flags().StringS("backup", "backup", "", "Path to backup the existing state file before modifying.")
	applyCmd.Flags().BoolS("compact-warnings", "compact-warnings", false, "Show warnings in a more compact form that includes only the summary messages.")
	applyCmd.Flags().BoolS("concise", "concise", false, "Disable progress-related messages.")
	applyCmd.Flags().BoolS("consolidate-errors", "consolidate-errors", false, "If OpenTofu produces any errors, attempt to consolidate similar messages into a single item.")
	applyCmd.Flags().BoolS("consolidate-warnings", "consolidate-warnings", false, "If OpenTofu produces any warnings, do not attempt to consolidate similar messages.")
	applyCmd.Flags().BoolS("destroy", "destroy", false, "Destroy OpenTofu-managed infrastructure.")
	applyCmd.Flags().StringS("exclude", "exclude", "", "Limit the operation to not operate on the given module, resource, or resource instance.")
	applyCmd.Flags().StringS("exclude-file", "exclude-file", "", "Similar to -exclude, but specifies zero or more resource addresses from a file.")
	applyCmd.Flags().BoolS("input", "input", false, "Ask for input for variables if not directly set.")
	applyCmd.Flags().BoolS("json", "json", false, "Produce output in a machine-readable JSON format.")
	applyCmd.Flags().StringS("json-into", "json-into", "", "Produce the same output as -json, but sent directly to the given file.")
	applyCmd.Flags().BoolS("lock", "lock", false, "Don't hold a state lock during the operation.")
	applyCmd.Flags().StringS("lock-timeout", "lock-timeout", "", "Duration to retry a state lock.")
	applyCmd.Flags().BoolS("no-color", "no-color", false, "If specified, output won't contain any color.")
	applyCmd.Flags().StringS("parallelism", "parallelism", "", "Limit the number of parallel resource operations.")
	applyCmd.Flags().BoolS("refresh", "refresh", false, "Skip checking for external changes to remote objects while creating the plan.")
	applyCmd.Flags().BoolS("refresh-only", "refresh-only", false, "Select the \"refresh only\" planning mode.")
	applyCmd.Flags().StringS("replace", "replace", "", "Force replacement of a particular resource instance using its resource address.")
	applyCmd.Flags().BoolS("show-sensitive", "show-sensitive", false, "If specified, sensitive values will be displayed.")
	applyCmd.Flags().StringS("state", "state", "", "Path to read and save state.")
	applyCmd.Flags().StringS("state-out", "state-out", "", "Path to write state to that is different than \"-state\".")
	applyCmd.Flags().BoolS("suppress-forget-errors", "suppress-forget-errors", false, "Suppress the error that occurs when a destroy operation completes successfully but leaves forgotten instances behind.")
	applyCmd.Flags().StringS("target", "target", "", "Limit the operation to only the given module, resource, or resource instance.")
	applyCmd.Flags().StringS("target-file", "target-file", "", "Similar to -target, but specifies zero or more resource addresses from a file.")
	applyCmd.Flags().StringArrayS("var", "var", nil, "Set a variable in the OpenTofu configuration.")
	applyCmd.Flags().StringS("var-file", "var-file", "", "Set variables in the OpenTofu configuration from a file.")
	rootCmd.AddCommand(applyCmd)

	applyCmd.Flag("backup").NoOptDefVal = " "
	applyCmd.Flag("exclude").NoOptDefVal = " "
	applyCmd.Flag("exclude-file").NoOptDefVal = " "
	applyCmd.Flag("json-into").NoOptDefVal = " "
	applyCmd.Flag("lock-timeout").NoOptDefVal = " "
	applyCmd.Flag("parallelism").NoOptDefVal = " "
	applyCmd.Flag("replace").NoOptDefVal = " "
	applyCmd.Flag("state").NoOptDefVal = " "
	applyCmd.Flag("state-out").NoOptDefVal = " "
	applyCmd.Flag("target").NoOptDefVal = " "
	applyCmd.Flag("target-file").NoOptDefVal = " "
	applyCmd.Flag("var-file").NoOptDefVal = " "

	carapace.Gen(applyCmd).FlagCompletion(carapace.ActionMap{
		"backup":       carapace.ActionFiles(),
		"exclude":      action.ActionResources(applyCmd).MultiParts("."),
		"exclude-file": carapace.ActionFiles(),
		"json-into":    carapace.ActionFiles(),
		"replace":      action.ActionResources(applyCmd).MultiParts("."),
		"state":        carapace.ActionFiles(),
		"state-out":    carapace.ActionFiles(),
		"target":       action.ActionResources(applyCmd).MultiParts("."),
		"target-file":  carapace.ActionFiles(),
		"var-file":     carapace.ActionFiles(),
	})

	carapace.Gen(applyCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
