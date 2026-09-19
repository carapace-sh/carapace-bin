package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugwireprotoCmd = &cobra.Command{
	Use:    "debugwireproto",
	Short:  "send wire protocol commands to a server",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugwireprotoCmd).Standalone()

	debugwireprotoCmd.Flags().Bool("insecure", false, "do not verify server certificate (ignoring web.cacerts config)")
	debugwireprotoCmd.Flags().Bool("localssh", false, "start an SSH server for this repo")
	debugwireprotoCmd.Flags().Bool("nologhandshake", false, "do not log I/O related to the peer handshake")
	debugwireprotoCmd.Flags().Bool("noreadstderr", false, "do not read from stderr of the remote")
	debugwireprotoCmd.Flags().String("peer", "", "construct a specific version of the peer")
	debugwireprotoCmd.Flags().String("remotecmd", "", "specify hg command to run on the remote side")
	debugwireprotoCmd.Flags().StringP("ssh", "e", "", "specify ssh command to use")
	rootCmd.AddCommand(debugwireprotoCmd)

	carapace.Gen(debugwireprotoCmd).FlagCompletion(carapace.ActionMap{
		"remotecmd": carapace.ActionExecutables(),
	})
}
