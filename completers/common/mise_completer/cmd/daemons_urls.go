package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var daemons_urlsCmd = &cobra.Command{
	Use:   "urls",
	Short: "[experimental] Show each daemon's URL and port",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(daemons_urlsCmd).Standalone()

	daemons_urlsCmd.Flags().BoolP("help", "h", false, "Print help")
	daemons_urlsCmd.Flags().Bool("json", false, "Output in JSON format, including ports, URLs, and data directories")
	daemonsCmd.AddCommand(daemons_urlsCmd)
}
