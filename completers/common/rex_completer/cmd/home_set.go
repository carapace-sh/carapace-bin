package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var homeSetCmd = &cobra.Command{
	Use:   "set ENDPOINT",
	Short: "Follow ENDPOINT for hosts and shortcuts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(homeSetCmd).Standalone()

	homeSetCmd.Flags().String("label", "", "call the home NAME in output")

	homeCmd.AddCommand(homeSetCmd)
}
