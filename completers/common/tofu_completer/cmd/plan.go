package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tofu_completer/cmd/action"
	"github.com/spf13/cobra"
)

var planCmd = &cobra.Command{
	Use:     "plan [options]",
	Short:   "Show changes required by the current configuration",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(planCmd).Standalone()

	planCmd.Flags().StringS("backup", "backup", "", "Path to backup the existing state file before modifying.")
	planCmd.Flags().BoolS("compact-warnings", "compact-warnings", false, "Show warnings in a more compact form that includes only the summary messages.")
	planCmd.Flags().BoolS("concise", "concise", false, "Disable progress-related messages.")
	planCmd.Flags().BoolS("consolidate-errors", "consolidate-errors", false, "If OpenTofu produces any errors, attempt to consolidate similar messages into a single item.")
	planCmd.Flags().BoolS("consolidate-warnings", "consolidate-warnings", false, "If OpenTofu produces any warnings, do not attempt to consolidate similar messages.")
	planCmd.Flags().StringS("deprecation", "deprecation", "", "Specify what type of warnings are shown.")
	planCmd.Flags().BoolS("destroy", "destroy", false, "Select the \"destroy\" planning mode.")
	planCmd.Flags().BoolS("detailed-exitcode", "detailed-exitcode", false, "Return detailed exit codes when the command exits.")
	planCmd.Flags().StringS("exclude", "exclude", "", "Limit the planning operation to not operate on the given module, resource, or resource instance.")
	planCmd.Flags().StringS("exclude-file", "exclude-file", "", "Similar to -exclude, but specifies zero or more resource addresses from a file.")
	planCmd.Flags().String("generate-config-out", "", "Write HCL configuration for resources to path.")
	planCmd.Flags().BoolS("input", "input", false, "Ask for input for variables if not directly set.")
	planCmd.Flags().BoolS("json", "json", false, "Produce output in a machine-readable JSON format.")
	planCmd.Flags().StringS("json-into", "json-into", "", "Produce the same output as -json, but sent directly to the given file.")
	planCmd.Flags().BoolS("lock", "lock", false, "Don't hold a state lock during the operation.")
	planCmd.Flags().StringS("lock-timeout", "lock-timeout", "", "Duration to retry a state lock.")
	planCmd.Flags().BoolS("no-color", "no-color", false, "If specified, output won't contain any color.")
	planCmd.Flags().StringS("out", "out", "", "Write a plan file to the given path.")
	planCmd.Flags().StringS("parallelism", "parallelism", "", "Limit the number of concurrent operations.")
	planCmd.Flags().BoolS("refresh", "refresh", false, "Skip checking for external changes to remote objects while creating the plan.")
	planCmd.Flags().BoolS("refresh-only", "refresh-only", false, "Select the \"refresh only\" planning mode.")
	planCmd.Flags().StringS("replace", "replace", "", "Force replacement of a particular resource instance using its resource address.")
	planCmd.Flags().BoolS("show-sensitive", "show-sensitive", false, "If specified, sensitive values will not be redacted in the UI output.")
	planCmd.Flags().StringS("state", "state", "", "A legacy option used for the local backend only.")
	planCmd.Flags().StringS("state-out", "state-out", "", "Path to write state to that is different than \"-state\".")
	planCmd.Flags().StringS("target", "target", "", "Limit the planning operation to only the given module, resource, or resource instance.")
	planCmd.Flags().StringS("target-file", "target-file", "", "Similar to -target, but specifies zero or more resource addresses from a file.")
	planCmd.Flags().StringArrayS("var", "var", nil, "Set a value for one of the input variables in the root module of the configuration.")
	planCmd.Flags().StringS("var-file", "var-file", "", "Load variable values from the given file.")
	rootCmd.AddCommand(planCmd)

	planCmd.Flag("backup").NoOptDefVal = " "
	planCmd.Flag("deprecation").NoOptDefVal = " "
	planCmd.Flag("exclude").NoOptDefVal = " "
	planCmd.Flag("exclude-file").NoOptDefVal = " "
	planCmd.Flag("generate-config-out").NoOptDefVal = " "
	planCmd.Flag("json-into").NoOptDefVal = " "
	planCmd.Flag("lock-timeout").NoOptDefVal = " "
	planCmd.Flag("out").NoOptDefVal = " "
	planCmd.Flag("parallelism").NoOptDefVal = " "
	planCmd.Flag("replace").NoOptDefVal = " "
	planCmd.Flag("state").NoOptDefVal = " "
	planCmd.Flag("state-out").NoOptDefVal = " "
	planCmd.Flag("target").NoOptDefVal = " "
	planCmd.Flag("target-file").NoOptDefVal = " "
	planCmd.Flag("var-file").NoOptDefVal = " "

	carapace.Gen(planCmd).FlagCompletion(carapace.ActionMap{
		"backup":              carapace.ActionFiles(),
		"deprecation":         carapace.ActionValues("module:all", "module:local", "module:none"),
		"exclude":             action.ActionResources(planCmd).MultiParts("."),
		"exclude-file":        carapace.ActionFiles(),
		"generate-config-out": carapace.ActionFiles(),
		"json-into":           carapace.ActionFiles(),
		"out":                 carapace.ActionFiles(),
		"replace":             action.ActionResources(planCmd).MultiParts("."),
		"state":               carapace.ActionFiles(),
		"state-out":           carapace.ActionFiles(),
		"target":              action.ActionResources(planCmd).MultiParts("."),
		"target-file":         carapace.ActionFiles(),
		"var-file":            carapace.ActionFiles(),
	})
}
