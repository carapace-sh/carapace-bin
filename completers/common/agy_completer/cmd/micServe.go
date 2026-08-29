package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var micServeCmd = &cobra.Command{
	Use:     "mic-serve",
	GroupID: "integration",
	Short:   "Serve this machine's microphone to a CLI on another host",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(micServeCmd).Standalone()
	rootCmd.AddCommand(micServeCmd)
}
