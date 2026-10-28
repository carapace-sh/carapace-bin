package cmd

import (
	"time"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/faas-cli_completer/cmd/action"
	"github.com/spf13/cobra"
)

var readyCmd = &cobra.Command{
	Use:   "ready",
	Short: "Block until the gateway or a function is ready for use",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(readyCmd).Standalone()

	readyCmd.Flags().Int("attempts", 60, "Number of attempts to check the gateway")
	readyCmd.Flags().StringP("gateway", "g", "http://127.0.0.1:8080", "Gateway URL starting with http(s)://")
	readyCmd.Flags().Duration("interval", time.Second, "Interval between attempts in seconds")
	readyCmd.Flags().StringP("namespace", "n", "", "Namespace of the function")
	readyCmd.Flags().Bool("tls-no-verify", false, "Disable TLS validation")
	rootCmd.AddCommand(readyCmd)

	carapace.Gen(readyCmd).FlagCompletion(carapace.ActionMap{
		"namespace": action.ActionNamespaces(),
	})

	carapace.Gen(readyCmd).PositionalCompletion(
		action.ActionFunctions(),
	)
}
