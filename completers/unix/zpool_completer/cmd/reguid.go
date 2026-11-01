package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var reguidCmd = &cobra.Command{
	Use:     "reguid pool",
	Short:   "generate a new unique pool identifier",
	GroupID: "io",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reguidCmd).Standalone()

	reguidCmd.Flags().StringS("g", "g", "", "set the pool GUID to the provided value")

	rootCmd.AddCommand(reguidCmd)

	carapace.Gen(reguidCmd).PositionalCompletion(
		zfs.ActionPools(),
	)
}
