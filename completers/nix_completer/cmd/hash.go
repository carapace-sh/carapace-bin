package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hashCmd = &cobra.Command{
	Use:     "hash",
	Short:   "compute and convert cryptographic hashes",
	GroupID: "utility",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hashCmd).Standalone()

	rootCmd.AddCommand(hashCmd)
}
