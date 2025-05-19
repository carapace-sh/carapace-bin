package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scoutIndexCmd = &cobra.Command{
	Use:   "scout:index [-k|--key [KEY]] [--] <name>",
	Short: "Create an index",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scoutIndexCmd).Standalone()

	scoutIndexCmd.Flags().String("key", "", "The name of the primary key")
	rootCmd.AddCommand(scoutIndexCmd)
}
