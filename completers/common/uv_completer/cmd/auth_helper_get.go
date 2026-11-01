package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auth_helper_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve credentials for a URI",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_helper_getCmd).Standalone()

	auth_helperCmd.AddCommand(auth_helper_getCmd)
}
