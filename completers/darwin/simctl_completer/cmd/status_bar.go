package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/simctl"
	"github.com/spf13/cobra"
)

var statusBarCmd = &cobra.Command{
	Use:   "status_bar",
	Short: "Set or clear status bar overrides",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(statusBarCmd).Standalone()
	rootCmd.AddCommand(statusBarCmd)
	carapace.Gen(statusBarCmd).PositionalCompletion(
		simctl.ActionDevices(),
		carapace.ActionValues("list", "clear", "override"),
	)
}
