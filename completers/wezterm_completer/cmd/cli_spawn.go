package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/wezterm"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/carapace-sh/carapace/pkg/traverse"
	"github.com/spf13/cobra"
)

var cli_spawnCmd = &cobra.Command{
	Use:   "spawn [OPTIONS] [PROG]...",
	Short: "Spawn a command into a new window or tab",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cli_spawnCmd).Standalone()
	cli_spawnCmd.Flags().SetInterspersed(false)

	cli_spawnCmd.Flags().String("cwd", "", "Specify the current working directory for the initially spawned program")
	cli_spawnCmd.Flags().String("domain-name", "", "")
	cli_spawnCmd.Flags().BoolP("help", "h", false, "Print help")
	cli_spawnCmd.Flags().Bool("new-window", false, "Spawn into a new window, rather than a new tab")
	cli_spawnCmd.Flags().String("pane-id", "", "Specify the current pane")
	cli_spawnCmd.Flags().String("window-id", "", "Specify the window into which to spawn a tab")
	cli_spawnCmd.Flags().String("workspace", "", "Override the default workspace name with the provided name")
	cliCmd.AddCommand(cli_spawnCmd)

	// TODO flag completion
	carapace.Gen(cli_spawnCmd).FlagCompletion(carapace.ActionMap{
		"cwd":       carapace.ActionDirectories(),
		"pane-id":   wezterm.ActionPanes(),
		"window-id": wezterm.ActionWindows(),
	})

	carapace.Gen(cli_spawnCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin().ChdirF(traverse.Flag(cli_spawnCmd.Flag("cwd"))),
	)
}
