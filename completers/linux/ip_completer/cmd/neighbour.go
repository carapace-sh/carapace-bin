package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neighbourCmd = &cobra.Command{
	Use:     "neighbour",
	Aliases: []string{"neigh", "n"},
	Short:   "manage ARP or NDISC cache entries",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neighbourCmd).Standalone()

	rootCmd.AddCommand(neighbourCmd)
}
