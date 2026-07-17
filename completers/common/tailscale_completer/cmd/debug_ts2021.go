package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_ts2021Cmd = &cobra.Command{
	Use:   "ts2021",
	Short: "Debug ts2021 protocol connectivity",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_ts2021Cmd).Standalone()

	debug_ts2021Cmd.Flags().String("ace", "", "ACE server to use")
	debug_ts2021Cmd.Flags().String("dial-plan", "", "JSON file path for dial plan")
	debug_ts2021Cmd.Flags().String("host", "", "host to connect to")
	debug_ts2021Cmd.Flags().Bool("verbose", false, "verbose output")
	debug_ts2021Cmd.Flags().Int("version", 0, "protocol version")
	debugCmd.AddCommand(debug_ts2021Cmd)
}
