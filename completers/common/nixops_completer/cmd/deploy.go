package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy machines",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deployCmd).Standalone()

	deployCmd.Flags().Bool("allow-reboot", false, "Allow NixOps to reboot the instance if necessary")
	deployCmd.Flags().Bool("allow-recreate", false, "Recreate resources that have disappeared")
	deployCmd.Flags().Bool("build-only", false, "Just build the configuration locally")
	deployCmd.Flags().Bool("check", false, "Verify current deployment state")
	deployCmd.Flags().Bool("copy-only", false, "Exit after building and copying closures")
	deployCmd.Flags().Bool("dry-run", false, "Show what would be done without actually doing it")
	deployCmd.Flags().StringSlice("exclude", nil, "Operate on all machines except the specified ones")
	deployCmd.Flags().Bool("force-reboot", false, "Reboot to activate the new configuration")
	deployCmd.Flags().StringSlice("include", nil, "Only operate on the specified machines")
	deployCmd.Flags().BoolP("kill-obsolete", "k", false, "Destroy virtual machines that are no longer in the spec")
	deployCmd.Flags().String("max-concurrent-copy", "", "Max concurrent nix-copy-closure processes")
	deployCmd.Flags().Bool("repair", false, "Use --repair when calling nix-build")

	rootCmd.AddCommand(deployCmd)
}
