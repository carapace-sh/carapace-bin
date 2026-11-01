package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var allocateASNCmd = &cobra.Command{
	Use:   "allocateASN",
	Short: "Ask launchservicesd to allocate an ASN",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(allocateASNCmd).Standalone()
	rootCmd.AddCommand(allocateASNCmd)
}
