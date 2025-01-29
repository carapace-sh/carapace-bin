package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var machine_client9pCmd = &cobra.Command{
	Use:    "client9p PORT DIR",
	Short:  "Mount a remote directory using 9p over hvsock",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(machine_client9pCmd).Standalone()

	machineCmd.AddCommand(machine_client9pCmd)
}
