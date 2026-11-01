package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/docker-compose_completer/cmd/action"
	"github.com/spf13/cobra"
)

var portCmd = &cobra.Command{
	Use:   "port [OPTIONS] SERVICE PRIVATE_PORT",
	Short: "Print the public port for a port binding",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(portCmd).Standalone()

	portCmd.Flags().Int("index", 0, "Index of the container if service has multiple replicas")
	portCmd.Flags().String("protocol", "tcp", "tcp or udp")
	rootCmd.AddCommand(portCmd)

	// TODO index

	carapace.Gen(portCmd).FlagCompletion(carapace.ActionMap{
		"protocol": carapace.ActionValues("tcp", "udp"),
	})

	carapace.Gen(portCmd).PositionalCompletion(
		action.ActionServices(portCmd),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionPorts(portCmd, c.Args[0])
		}),
	)
}
