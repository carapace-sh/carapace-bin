package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pip_debugCmd = &cobra.Command{
	Use:    "debug",
	Short:  "Display debug information (unsupported)",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pip_debugCmd).Standalone()

	pip_debugCmd.Flags().String("abi", "", "")
	pip_debugCmd.Flags().String("implementation", "", "")
	pip_debugCmd.Flags().String("platform", "", "")
	pip_debugCmd.Flags().String("python-version", "", "")
	pip_debugCmd.Flag("abi").Hidden = true
	pip_debugCmd.Flag("implementation").Hidden = true
	pip_debugCmd.Flag("platform").Hidden = true
	pip_debugCmd.Flag("python-version").Hidden = true
	pipCmd.AddCommand(pip_debugCmd)
}
