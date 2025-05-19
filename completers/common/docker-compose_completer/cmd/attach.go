package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/docker-compose_completer/cmd/action"
	"github.com/spf13/cobra"
)

var attachCmd = &cobra.Command{
	Use:   "attach [OPTIONS] SERVICE",
	Short: "Attach local standard input, output, and error streams to a service's running container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(attachCmd).Standalone()

	attachCmd.Flags().String("detach-keys", "", "Override the key sequence for detaching from a container.")
	attachCmd.Flags().String("index", "", "index of the container if service has multiple replicas.")
	attachCmd.Flags().Bool("no-stdin", false, "Do not attach STDIN")
	attachCmd.Flags().Bool("sig-proxy", false, "Proxy all received signals to the process")
	rootCmd.AddCommand(attachCmd)

	carapace.Gen(attachCmd).PositionalCompletion(
		action.ActionServices(attachCmd).FilterArgs(),
	)
}
