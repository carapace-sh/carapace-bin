package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/terramate"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run command in the stacks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()
	runCmd.Flags().SetInterspersed(false)

	runCmd.Flags().BoolS("X", "X", false, "Disable all safeguards")
	runCmd.Flags().Bool("cloud-sync-deployment", false, "Enable synchronization of stack execution with the Terramate Cloud")
	runCmd.Flags().Bool("cloud-sync-drift-status", false, "Enable drift detection and synchronization with the Terramate Cloud")
	runCmd.Flags().String("cloud-sync-terraform-plan-file", "", "Enable sync of Terraform plan file")
	runCmd.Flags().Bool("continue-on-error", false, "Continue executing in other stacks in case of error")
	runCmd.Flags().Bool("disable-check-gen-code", false, "Disable outdated generated code check")
	runCmd.Flags().Bool("disable-check-git-remote", false, "Disable checking if local default branch is updated with remote")
	runCmd.Flags().String("disable-safeguards", "", "Disable safeguards: {all,none,git,git-untracked,git-uncommitted,outdated-code,git-out-of-sync}")
	runCmd.Flags().Bool("dry-run", false, "Plan the execution but do not execute it")
	runCmd.Flags().Bool("eval", false, "Evaluate command line arguments as HCL strings")
	runCmd.Flags().Bool("no-recursive", false, "Do not recurse into child stacks")
	runCmd.Flags().Bool("reverse", false, "Reverse the order of execution")
	rootCmd.AddCommand(runCmd)

	carapace.Gen(runCmd).FlagCompletion(carapace.ActionMap{
		"cloud-sync-terraform-plan-file": carapace.ActionValues(), // TODO enable dynamic completion
		"disable-safeguards":             terramate.ActionSafeguards().UniqueList(","),
	})

	carapace.Gen(runCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin(),
	)

	carapace.Gen(runCmd).DashAnyCompletion(
		carapace.ActionPositional(runCmd),
	)
}
