package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var store_pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "test whether a store can be accessed",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_pingCmd).Standalone()

	store_pingCmd.Flags().Bool("json", false, "Produce output in JSON format")
	storeCmd.AddCommand(store_pingCmd)

	addLoggingFlags(store_pingCmd)
}
