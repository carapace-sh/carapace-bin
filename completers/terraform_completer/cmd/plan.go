package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/terraform_completer/cmd/action"
	"github.com/spf13/cobra"
)

var planCmd = &cobra.Command{
	Use:   "plan",
	Short: "Show changes required by the current configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(planCmd).Standalone()

	planCmd.Flags().Bool("compact-warnings", false, "Show warnings in a more compact form that includes only the summary messages.")
	planCmd.Flags().Bool("destroy", false, "Select the \"destroy\" planning mode.")
	planCmd.Flags().Bool("detailed-exitcode", false, "Return detailed exit codes when the command exits.")
	planCmd.Flags().String("input", "", "Ask for input for variables if not directly set.")
	planCmd.Flags().String("lock", "", "Don't hold a state lock during the operation.")
	planCmd.Flags().String("lock-timeout", "", "Duration to retry a state lock.")
	planCmd.Flags().Bool("no-color", false, "If specified, output won't contain any color.")
	planCmd.Flags().String("out", "", "Write a plan file to the given path.")
	planCmd.Flags().String("parallelism", "", "Limit the number of concurrent operations.")
	planCmd.Flags().String("refresh", "", "Skip checking for external changes to remote objects while creating the plan.")
	planCmd.Flags().Bool("refresh-only", false, "Select the \"refresh only\" planning mode.")
	planCmd.Flags().String("replace", "", "Force replacement of a particular resource instance using its resource address.")
	planCmd.Flags().String("state", "", "A legacy option used for the local backend only.")
	planCmd.Flags().String("target", "", "Limit the planning operation to only the given module, resource.")
	planCmd.Flags().String("var", "", "Set a value for one of the input variables in the root module of the configuration.")
	planCmd.Flags().String("var-file", "", "Load variable values from the given file.")
	rootCmd.AddCommand(planCmd)

	planCmd.Flag("refresh").NoOptDefVal = " "
	planCmd.Flag("replace").NoOptDefVal = " "
	planCmd.Flag("target").NoOptDefVal = " "
	planCmd.Flag("var-file").NoOptDefVal = " "
	planCmd.Flag("input").NoOptDefVal = " "
	planCmd.Flag("lock").NoOptDefVal = " "
	planCmd.Flag("lock-timeout").NoOptDefVal = " "
	planCmd.Flag("out").NoOptDefVal = " "
	planCmd.Flag("parallelism").NoOptDefVal = " "
	planCmd.Flag("state").NoOptDefVal = " "

	// TODO resource completion
	carapace.Gen(planCmd).FlagCompletion(carapace.ActionMap{
		"lock":     action.ActionBool(),
		"out":      carapace.ActionFiles(),
		"refresh":  action.ActionBool(),
		"state":    carapace.ActionFiles(),
		"var-file": carapace.ActionFiles(),
	})
}
