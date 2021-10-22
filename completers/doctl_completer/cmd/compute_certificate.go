package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_certificateCmd = &cobra.Command{
	Use:   "certificate",
	Short: "Display commands that manage SSL certificates and private keys",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_certificateCmd).Standalone()
	computeCmd.AddCommand(compute_certificateCmd)
}
