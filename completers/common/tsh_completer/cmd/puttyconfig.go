package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var puttyconfigCmd = &cobra.Command{
	Use:    "puttyconfig",
	Short:  "Add PuTTY saved session configuration for specified hostname to Windows registry",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(puttyconfigCmd).Standalone()

	puttyconfigCmd.Flags().String("leaf", "", "Add a configuration for connecting to a leaf cluster")
	puttyconfigCmd.Flags().StringP("port", "p", "", "SSH port on a remote host")
	rootCmd.AddCommand(puttyconfigCmd)

	carapace.Gen(puttyconfigCmd).FlagCompletion(carapace.ActionMap{
		"port": net.ActionPorts(),
	})
}
