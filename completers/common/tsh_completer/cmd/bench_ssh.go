package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bench_sshCmd = &cobra.Command{
	Use:    "ssh",
	Short:  "Run SSH benchmark tests.",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bench_sshCmd).Standalone()

	bench_sshCmd.Flags().Bool("no-no-resume", false, "Disable SSH connection resumption.")
	bench_sshCmd.Flags().Bool("no-random", false, "Connect to random hosts for each SSH session. The provided hostname must be all: tsh bench ssh --random <user>@all <command>.")
	bench_sshCmd.Flags().Bool("no-resume", false, "Disable SSH connection resumption.")
	bench_sshCmd.Flags().StringP("port", "p", "", "SSH port on a remote host.")
	bench_sshCmd.Flags().Bool("random", false, "Connect to random hosts for each SSH session. The provided hostname must be all: tsh bench ssh --random <user>@all <command>.")
	bench_sshCmd.Flag("no-no-resume").Hidden = true
	bench_sshCmd.Flag("no-random").Hidden = true
	benchCmd.AddCommand(bench_sshCmd)
}
