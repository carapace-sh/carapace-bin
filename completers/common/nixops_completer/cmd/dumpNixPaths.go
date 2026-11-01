package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dumpNixPathsCmd = &cobra.Command{
	Use:   "dump-nix-paths",
	Short: "dump Nix paths referenced in deployments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dumpNixPathsCmd).Standalone()

	dumpNixPathsCmd.Flags().Bool("all", false, "Dump Nix paths for all deployments")
	dumpNixPathsCmd.Flags().Bool("allow-reboot", false, "Allow NixOps to reboot the instance if necessary")
	dumpNixPathsCmd.Flags().Bool("check", false, "Verify current deployment state")
	dumpNixPathsCmd.Flags().StringSlice("exclude", nil, "Operate on all machines except the specified ones")
	dumpNixPathsCmd.Flags().Bool("force-reboot", false, "Reboot to activate the new configuration")
	dumpNixPathsCmd.Flags().StringSlice("include", nil, "Only operate on the specified machines")
	dumpNixPathsCmd.Flags().String("max-concurrent-activate", "", "Max concurrent machine activations")
	dumpNixPathsCmd.Flags().String("max-concurrent-copy", "", "Max concurrent nix-copy-closure processes")
	dumpNixPathsCmd.Flags().Bool("no-sync", false, "Do not flush buffers to disk")

	rootCmd.AddCommand(dumpNixPathsCmd)
}
