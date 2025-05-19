package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var whereCmd = &cobra.Command{
	Use:   "where",
	Short: "Find where your cloud.charm.sh folder resides on your machine",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(whereCmd).Standalone()

	rootCmd.AddCommand(whereCmd)
}
