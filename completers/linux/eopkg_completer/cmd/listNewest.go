package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listNewestCmd = &cobra.Command{
	Use:     "list-newest",
	Short:   "List newest packages in the repositories",
	Aliases: []string{"ln"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listNewestCmd).Standalone()

	listNewestCmd.Flags().String("since", "", "List packages added after given date")
	listNewestCmd.Flags().StringP("last", "l", "", "List last n packages")

	rootCmd.AddCommand(listNewestCmd)
}
