package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var program_upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade an upgradeable program",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(program_upgradeCmd).Standalone()

	program_upgradeCmd.Flags().String("buffer", "", "Existing buffer account to upgrade from. If not provided, auto-discovers program from workspace")
	program_upgradeCmd.Flags().BoolP("help", "h", false, "Print help")
	program_upgradeCmd.Flags().String("max-retries", "0", "Max times to retry on failure")
	program_upgradeCmd.Flags().String("program-filepath", "", "Program filepath (e.g., target/deploy/my_program.so). If not provided, discovers from workspace")
	program_upgradeCmd.Flags().StringP("program-name", "p", "", "Program name to upgrade (from workspace). Used when program_filepath is not provided")
	program_upgradeCmd.Flags().String("upgrade-authority", "", "Upgrade authority (defaults to configured wallet)")
	program_upgradeCmd.Flags().Bool("use-rpc", false, "Send write transactions through RPC instead of TPU")
	programCmd.AddCommand(program_upgradeCmd)
}
