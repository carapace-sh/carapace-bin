package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var darwinCmd = &cobra.Command{
	Use:   "darwin",
	Short: "Nix-darwin functionality",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(darwinCmd).Standalone()

	rootCmd.AddCommand(darwinCmd)
}

func addDarwinRebuildFlags(cmd *cobra.Command) {
	addCommonRebuildFlags(cmd)
	addUpdateFlags(cmd)
	cmd.Flags().StringP("hostname", "H", "", "Select this hostname from darwinConfigurations")
	cmd.Flags().BoolP("bypass-root-check", "R", false, "Don't panic if calling nh as root")
	cmd.Flags().Bool("show-activation-logs", false, "Show activation logs")
	cmd.Flags().String("build-host", "", "Build the configuration on a different host over SSH")
}
