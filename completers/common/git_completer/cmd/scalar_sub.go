package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scalar_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an existing Scalar enlistment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var scalar_diagnoseCmd = &cobra.Command{
	Use:   "diagnose",
	Short: "Collect diagnostic information for a Scalar enlistment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var scalar_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List Scalar enlistments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var scalar_reconfigureCmd = &cobra.Command{
	Use:   "reconfigure",
	Short: "Reconfigure Scalar enlistment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var scalar_registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a repository with Scalar",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var scalar_runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a Scalar maintenance task",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var scalar_unregisterCmd = &cobra.Command{
	Use:   "unregister",
	Short: "Unregister a repository from Scalar",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var scalar_versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show Scalar version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scalar_deleteCmd).Standalone()
	carapace.Gen(scalar_diagnoseCmd).Standalone()
	carapace.Gen(scalar_listCmd).Standalone()
	carapace.Gen(scalar_reconfigureCmd).Standalone()
	carapace.Gen(scalar_registerCmd).Standalone()
	carapace.Gen(scalar_runCmd).Standalone()
	carapace.Gen(scalar_unregisterCmd).Standalone()
	carapace.Gen(scalar_versionCmd).Standalone()

	scalar_reconfigureCmd.Flags().Bool("all", false, "reconfigure all enlistments")
	scalar_versionCmd.Flags().Bool("build-options", false, "also print build options")
	scalar_registerCmd.Flags().Bool("maintenance", false, "enable automatic maintenance")
	scalar_reconfigureCmd.Flags().String("maintenance", "", "maintenance mode")
	scalar_runCmd.Flags().String("task", "", "run a specific task")
	scalar_versionCmd.Flags().BoolP("verbose", "v", false, "include Git version")

	scalarCmd.AddCommand(scalar_deleteCmd)
	scalarCmd.AddCommand(scalar_diagnoseCmd)
	scalarCmd.AddCommand(scalar_listCmd)
	scalarCmd.AddCommand(scalar_reconfigureCmd)
	scalarCmd.AddCommand(scalar_registerCmd)
	scalarCmd.AddCommand(scalar_runCmd)
	scalarCmd.AddCommand(scalar_unregisterCmd)
	scalarCmd.AddCommand(scalar_versionCmd)

	carapace.Gen(scalar_reconfigureCmd).FlagCompletion(carapace.ActionMap{
		"maintenance": carapace.ActionValues("enable", "disable", "keep"),
	})

	carapace.Gen(scalar_runCmd).FlagCompletion(carapace.ActionMap{
		"task": carapace.ActionValues("all", "config", "commit-graph", "fetch", "loose-objects", "pack-files"),
	})
}
