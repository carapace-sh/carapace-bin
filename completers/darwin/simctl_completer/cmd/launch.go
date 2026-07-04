package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/simctl"
	"github.com/spf13/cobra"
)

var launchCmd = &cobra.Command{
	Use:   "launch",
	Short: "Launch an application by identifier on a device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(launchCmd).Standalone()
	launchCmd.Flags().Bool("console-pty", false, "Connect stdout/stderr of the app to the terminal")
	rootCmd.AddCommand(launchCmd)
	carapace.Gen(launchCmd).PositionalCompletion(
		simctl.ActionDevices(),
	)
}
