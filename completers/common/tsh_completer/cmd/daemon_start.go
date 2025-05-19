package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var daemon_startCmd = &cobra.Command{
	Use:    "start",
	Short:  "Starts tsh daemon service.",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(daemon_startCmd).Standalone()

	daemon_startCmd.Flags().String("addr", "", "Addr is the daemon listening address.")
	daemon_startCmd.Flags().String("certs-dir", "", "Directory containing certs used to create secure gRPC connection with daemon service")
	daemon_startCmd.Flags().String("prehog-addr", "", "URL where prehog events should be submitted")
	daemonCmd.AddCommand(daemon_startCmd)

	carapace.Gen(daemon_startCmd).FlagCompletion(carapace.ActionMap{
		"certs-dir": carapace.ActionDirectories(),
	})
}
