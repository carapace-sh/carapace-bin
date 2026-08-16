package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var request_createCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create a new access request.",
	Aliases: []string{"new"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(request_createCmd).Standalone()

	request_createCmd.Flags().String("assume-start-time", "", "Sets time roles can be assumed by requestor (RFC3339 e.g 2023-12-12T23:20:50.52Z).")
	request_createCmd.Flags().String("max-duration", "", "How long the access should be granted for.")
	request_createCmd.Flags().Bool("no-nowait", false, "Finish without waiting for request resolution.")
	request_createCmd.Flags().Bool("nowait", false, "Finish without waiting for request resolution.")
	request_createCmd.Flags().String("reason", "", "Reason for requesting.")
	request_createCmd.Flags().String("request-ttl", "", "Expiration time for the access request.")
	request_createCmd.Flags().String("resource", "", "Resource ID to be requested.")
	request_createCmd.Flags().String("reviewers", "", "Suggested reviewers.")
	request_createCmd.Flags().String("roles", "", "Roles to be requested.")
	request_createCmd.Flags().String("session-ttl", "", "Expiration time for the elevated certificate.")
	request_createCmd.Flag("no-nowait").Hidden = true
	requestCmd.AddCommand(request_createCmd)
}
