package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var studio_pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Run deployments configured within a Benthos Studio session",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(studio_pullCmd).Standalone()

	studio_pullCmd.Flags().String("name", "", "An explicit name to adopt in this instance, used to identify its connection to the session. Each running node must have a unique name, if left unset a name is generated each time the command is run.")
	studio_pullCmd.Flags().Bool("send-traces", false, "Whether to send trace data back to Studio during execution. This is opt-in and is used as a way to add trace events to the graph editor for testing and debugging configs. This is a very useful feature but should be used with caution as it exports information about messages passing through the stream.")
	studio_pullCmd.Flags().StringP("session", "s", "", "The session ID to synchronise with.")
	studio_pullCmd.Flags().String("token", "", "A token for the session, used to authenticate requests. If left blank the environment variable BSTDIO_NODE_TOKEN will be used instead.")
	studio_pullCmd.Flags().String("token-secret", "", "A token secret the session, used to authenticate requests. If left blank the environment variable BSTDIO_NODE_SECRET will be used instead.")
	studioCmd.AddCommand(studio_pullCmd)
}
