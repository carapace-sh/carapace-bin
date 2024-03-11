package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var startCmd = &cobra.Command{
	Use:   "start [OPTIONS] [PROG]...",
	Short: "Start the GUI, optionally running an alternative program",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(startCmd).Standalone()
	startCmd.Flags().SetInterspersed(false)

	startCmd.Flags().Bool("always-new-process", false, "Do not try to ask an existing wezterm GUI instance to start the command")
	startCmd.Flags().Bool("attach", false, "Attach to already running pane of domain")
	startCmd.Flags().String("class", "", "Override the default windowing system class")
	startCmd.Flags().String("cwd", "", "Specify the current working directory for the initially spawned program")
	startCmd.Flags().String("domain", "", "Name of the multiplexer domain section from the configuration to which you'd like to connect")
	startCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	startCmd.Flags().Bool("no-auto-connect", false, "Do not connect to domains marked as connect_automatically")
	startCmd.Flags().String("position", "", "Override the position for the initial window launched by this process")
	startCmd.Flags().String("workspace", "", "Override the default workspace with the provided name")
	rootCmd.AddCommand(startCmd)

	carapace.Gen(startCmd).FlagCompletion(carapace.ActionMap{
		"cwd": carapace.ActionDirectories(),
		// TODO domain
	})

	carapace.Gen(startCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin(),
	)

	carapace.Gen(startCmd).PreInvoke(func(cmd *cobra.Command, flag *pflag.Flag, action carapace.Action) carapace.Action {
		return action.Chdir(cmd.Flag("cwd").Value.String())
	})
}
