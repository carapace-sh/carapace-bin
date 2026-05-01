package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/nh_completer/cmd/common"
	"github.com/spf13/cobra"
)

var homeCmd = &cobra.Command{
	Use:   "home",
	Short: "Home-manager functionality",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(homeCmd).Standalone()

	rootCmd.AddCommand(homeCmd)
}

func addHomeRebuildFlags(cmd *cobra.Command) {
	common.AddCommonRebuildFlags(cmd)
	common.AddUpdateFlags(cmd)
	cmd.Flags().StringP("backup-extension", "b", "", "Move existing files by backing up with this file extension")
	cmd.Flags().String("build-host", "", "Build the configuration on a different host over SSH")
	cmd.Flags().StringP("configuration", "c", "", "Name of the flake homeConfigurations attribute")
	cmd.Flags().BoolP("no-specialisation", "S", false, "Ignore specialisations")
	cmd.Flags().Bool("show-activation-logs", false, "Show activation logs")
	cmd.Flags().StringP("specialisation", "s", "", "Explicitly select some specialisation")
}
