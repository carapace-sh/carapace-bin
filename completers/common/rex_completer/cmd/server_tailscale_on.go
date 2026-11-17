package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serverTailscaleOnCmd = &cobra.Command{
	Use:   "on",
	Short: "Serve the API over the tailnet",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serverTailscaleOnCmd).Standalone()

	serverTailscaleOnCmd.Flags().StringArray("advertise-tag", nil, "advertise ACL TAG when this call enrolls a new node (repeatable)")
	serverTailscaleOnCmd.Flags().Bool("allow-tailnet-members", false, "allow any caller admitted by the tailnet instead of only this node's owner")
	serverTailscaleOnCmd.Flags().String("auth-key-file", "", "enroll a new node with the one-time tailscale auth key read from FILE")
	serverTailscaleOnCmd.Flags().String("control-url", "", "register a node this call starts with the control server at URL")
	serverTailscaleOnCmd.Flags().String("hostname", "", "tailnet device NAME when this call starts the node")
	serverTailscaleOnCmd.Flags().Bool("json", false, "output the resulting state as JSON")

	serverTailscaleCmd.AddCommand(serverTailscaleOnCmd)

	carapace.Gen(serverTailscaleOnCmd).FlagCompletion(carapace.ActionMap{
		"auth-key-file": carapace.Batch(carapace.ActionValues("-"), carapace.ActionFiles()).ToA(),
	})
}
