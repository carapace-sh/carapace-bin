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
	daemon_startCmd.Flags().String("agents-dir", "", "Directory containing agent config files and data directories for Connect My Computer.")
	daemon_startCmd.Flags().String("certs-dir", "", "Directory containing certs used to create secure gRPC connection with daemon service.")
	daemon_startCmd.Flags().Bool("hardware-key-agent", false, "Serve the hardware key agent as part of the daemon process.")
	daemon_startCmd.Flags().String("installation-id", "", "Unique ID identifying a specific Teleport Connect installation.")
	daemon_startCmd.Flags().String("kubeconfigs-dir", "", "Directory containing kubeconfig for Kubernetes Access.")
	daemon_startCmd.Flags().Bool("no-hardware-key-agent", false, "Serve the hardware key agent as part of the daemon process.")
	daemon_startCmd.Flags().String("prehog-addr", "", "URL where prehog events should be submitted.")
	daemon_startCmd.Flag("no-hardware-key-agent").Hidden = true
	daemonCmd.AddCommand(daemon_startCmd)
}
