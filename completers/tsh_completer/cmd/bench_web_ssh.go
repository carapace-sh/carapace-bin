package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var bench_web_sshCmd = &cobra.Command{
	Use:    "ssh",
	Short:  "Run SSH benchmark tests",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bench_web_sshCmd).Standalone()

	bench_web_sshCmd.Flags().Bool("no-random", false, "Connect to random hosts for each SSH session. The provided hostname must be all: tsh bench ssh --random <user>@all <command>")
	bench_web_sshCmd.Flags().StringP("port", "p", "", "SSH port on a remote host")
	bench_web_sshCmd.Flags().Bool("random", false, "Connect to random hosts for each SSH session. The provided hostname must be all: tsh bench ssh --random <user>@all <command>")
	bench_web_sshCmd.Flag("no-random").Hidden = true
	bench_webCmd.AddCommand(bench_web_sshCmd)

	carapace.Gen(bench_web_sshCmd).FlagCompletion(carapace.ActionMap{
		"port": net.ActionPorts(),
	})
}
