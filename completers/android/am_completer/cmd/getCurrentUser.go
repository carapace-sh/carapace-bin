package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var getCurrentUserCmd = &cobra.Command{
	Use:   "get-current-user",
	Short: "Return the id of the current foreground user",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getCurrentUserCmd).Standalone()

	rootCmd.AddCommand(getCurrentUserCmd)

}
