package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var reverbStartCmd = &cobra.Command{
	Use:   "reverb:start [--host [HOST]] [--port [PORT]] [--hostname [HOSTNAME]] [--debug]",
	Short: "Start the Reverb server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reverbStartCmd).Standalone()

	reverbStartCmd.Flags().Bool("debug", false, "Indicates whether debug messages should be displayed in the terminal")
	reverbStartCmd.Flags().String("host", "", "The IP address the server should bind to")
	reverbStartCmd.Flags().String("hostname", "", "The hostname the server is accessible from")
	reverbStartCmd.Flags().String("port", "", "The port the server should listen on")
	rootCmd.AddCommand(reverbStartCmd)
}
