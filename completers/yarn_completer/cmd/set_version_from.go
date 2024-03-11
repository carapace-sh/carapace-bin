package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var set_version_fromCmd = &cobra.Command{
	Use:   "from",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(set_version_fromCmd).Standalone()

	set_versionCmd.AddCommand(set_version_fromCmd)
}
