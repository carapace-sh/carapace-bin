package cmd

import (
	"github.com/carapace-sh/carapace"
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

	applyCmd.Flags().BoolS("auto-approve", "auto-approve", false, "Skip interactive approval of plan before applying")
	applyCmd.Flags().StringS("backup", "backup", "", "Path to backup the existing state file before modifying")
	applyCmd.Flags().BoolS("compact-warnings", "compact-warnings", false, "Show wanings in a more compact form that includes only the summary messages")
	applyCmd.Flags().BoolS("destroy", "destroy", false, "Destroy Terraform-managed infrastructure")
	applyCmd.Flags().BoolS("input", "input", false, "Ask for input for variables if not directly set")
	applyCmd.Flags().BoolS("lock", "lock", false, "Don't hold a state lock during the operation dangerous if others might concurrently run commands against the same workspace.")
	applyCmd.Flags().StringS("lock-timeout", "lock-timeout", "", "Duration to retry a state lock")
	applyCmd.Flags().BoolS("no-color", "no-color", false, "If specified, output won't contain any color")
	applyCmd.Flags().StringS("parallelism", "parallelism", "", "Limit the number of parallel resource operations")
	applyCmd.Flags().StringS("state", "state", "", "Path to read and save state")
	applyCmd.Flags().StringS("state-out", "state-out", "", "Path to write state to that is different than \"-state\"")
	rootCmd.AddCommand(applyCmd)

	applyCmd.Flag("backup").NoOptDefVal = " "
	applyCmd.Flag("lock-timeout").NoOptDefVal = " "
	applyCmd.Flag("parallelism").NoOptDefVal = " "
	applyCmd.Flag("state").NoOptDefVal = " "
	applyCmd.Flag("state-out").NoOptDefVal = " "

	carapace.Gen(applyCmd).FlagCompletion(carapace.ActionMap{
		"backup":    carapace.ActionFiles(),
		"state":     carapace.ActionFiles(),
		"state-out": carapace.ActionFiles(),
	})

	carapace.Gen(applyCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
