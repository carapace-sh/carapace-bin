package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var request_reviewCmd = &cobra.Command{
	Use:   "review",
	Short: "Review an access request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(request_reviewCmd).Standalone()

	request_reviewCmd.Flags().Bool("approve", false, "Review proposes approval.")
	request_reviewCmd.Flags().String("assume-start-time", "", "Sets time roles can be assumed by requestor (RFC3339 e.g 2023-12-12T23:20:50.52Z).")
	request_reviewCmd.Flags().Bool("deny", false, "Review proposes denial.")
	request_reviewCmd.Flags().Bool("no-approve", false, "Review proposes approval.")
	request_reviewCmd.Flags().Bool("no-deny", false, "Review proposes denial.")
	request_reviewCmd.Flags().String("reason", "", "Review reason message.")
	request_reviewCmd.Flag("no-approve").Hidden = true
	request_reviewCmd.Flag("no-deny").Hidden = true
	requestCmd.AddCommand(request_reviewCmd)
}
