package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net/ssh"
	"github.com/spf13/cobra"
)

func ActionHosts(cmd *cobra.Command) carapace.Action {
	return carapace.Batch(
		net.ActionHosts(),
		ssh.ActionHosts(cmd.Flag("F").Value.String()).Style("yellow"),
	).ToA()
}
