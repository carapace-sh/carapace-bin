package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var blobl_serverCmd = &cobra.Command{
	Use:   "server",
	Short: "EXPERIMENTAL: Run a web server that hosts a Bloblang app",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(blobl_serverCmd).Standalone()

	blobl_serverCmd.Flags().String("host", "", "the host to bind to.")
	blobl_serverCmd.Flags().StringP("input-file", "i", "", "an optional path to an input file to load as the initial input to the mapping within the app.")
	blobl_serverCmd.Flags().StringP("mapping-file", "m", "", "an optional path to a mapping file to load as the initial mapping within the app.")
	blobl_serverCmd.Flags().BoolP("no-open", "n", false, "do not open the app in the browser automatically.")
	blobl_serverCmd.Flags().StringP("port", "p", "", "the port to bind to.")
	blobl_serverCmd.Flags().BoolP("write", "w", false, "when editing a mapping and/or input file write changes made back to the respective source file, if the file does not exist it will be created.")
	bloblCmd.AddCommand(blobl_serverCmd)

	carapace.Gen(blobl_serverCmd).FlagCompletion(carapace.ActionMap{
		"input-file":   carapace.ActionFiles(),
		"mapping-file": carapace.ActionFiles(),
		"port":         net.ActionPorts(),
	})
}
